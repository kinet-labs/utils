package units

// Denominations of value
// KINET uses 6 decimals (like USDC), allowing max supply of ~18.4 trillion KINET in uint64
const (
	MicroKinet  uint64 = 1                    // Base unit (6 decimals) - 0.000001 KNT
	MilliKinet  uint64 = 1000 * MicroKinet      // 0.001 KNT
	Kinet       uint64 = 1000 * MilliKinet      // 1 KNT = 10^6 microKNT
	KiloKinet   uint64 = 1000 * Kinet           // 1,000 KNT
	MegaKinet   uint64 = 1000 * KiloKinet       // 1,000,000 KNT
	GigaKinet   uint64 = 1000 * MegaKinet       // 1,000,000,000 KNT (1 billion)
	TeraKinet   uint64 = 1000 * GigaKinet       // 1,000,000,000,000 KNT (1 trillion)

	// Schmeckle preserved for compatibility (≈49.463 milliKNT)
	Schmeckle uint64 = 49*MilliKinet + 463*MicroKinet

	// NanoKinet deprecated - use MicroKinet as base unit
	// Kept for backward compatibility but represents same as MicroKinet
	NanoKinet uint64 = MicroKinet
)
