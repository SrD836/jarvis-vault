// Package crypto is the niche estimator for asset price-level markets
// (BTC/ETH/SOL/WTI/SPX/...). It prices "will <asset> be above/below/hit $K by
// T" from live spot + realized volatility instead of asking an LLM to guess —
// the bot's only category with a real, structural edge.
package crypto

import "math"

// Prob returns P(YES resolves true) for an asset price-level market under a
// driftless geometric-Brownian-motion model.
//
//	spot   live price S
//	target strike K
//	sigma  ANNUALIZED volatility
//	tau    time to resolution in YEARS
//	dir    "above" | "below"
//	mode   "terminal" (resolves on S_T at expiry) | "touch" (resolves if the
//	       barrier is hit any time before T; uses the reflection principle)
//
// Driftless (mu=0) is the honest prior for short-dated crypto: we have no
// edge on direction, only on the distribution. Result clamped to [0.01, 0.99].
func Prob(spot, target, sigma, tau float64, dir, mode string) float64 {
	sd := sigma * math.Sqrt(tau) // stdev of the log-return over the horizon
	if sd <= 0 || spot <= 0 || target <= 0 {
		return 0.5 // no usable inputs → no opinion
	}

	if mode == "touch" {
		switch dir {
		case "above":
			if spot >= target {
				return 0.99 // already at/through the barrier
			}
			// P(max >= K) = 2*(1 - N(ln(K/S)/sd)) for driftless BM (reflection).
			return clamp(2 * (1 - normCDF(math.Log(target/spot)/sd)))
		default: // below
			if spot <= target {
				return 0.99
			}
			return clamp(2 * (1 - normCDF(math.Log(spot/target)/sd)))
		}
	}

	// terminal: P(S_T > K) = N((ln(S/K) - 0.5*sigma^2*tau)/sd).
	above := normCDF((math.Log(spot/target) - 0.5*sigma*sigma*tau) / sd)
	if dir == "below" {
		return clamp(1 - above)
	}
	return clamp(above)
}

// normCDF is the standard normal CDF via the complementary error function.
func normCDF(x float64) float64 {
	return 0.5 * math.Erfc(-x/math.Sqrt2)
}

func clamp(p float64) float64 {
	if p < 0.01 {
		return 0.01
	}
	if p > 0.99 {
		return 0.99
	}
	return p
}
