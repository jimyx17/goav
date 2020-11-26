package libavcodec

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
import "C"

// AV_ROUND_xxx
const (
	AvRoundZero    = C.AV_ROUND_ZERO     ///< Round toward zero.
	AvRoundInf     = C.AV_ROUND_INF      ///< Round away from zero.
	AvRoundDown    = C.AV_ROUND_DOWN     ///< Round toward -infinity.
	AvRoundUp      = C.AV_ROUND_UP       ///< Round toward +infinity.
	AvRoundNearInf = C.AV_ROUND_NEAR_INF ///< Round to nearest and halfway cases away from zero.
	/**
	 * Flag telling rescaling functions to pass `INT64_MIN`/`MAX` through
	 * unchanged, avoiding special cases for #AV_NOPTS_VALUE.
	 *
	 * Unlike other values of the enumeration AvRounding, this value is a
	 * bitmask that must be used in conjunction with another value of the
	 * enumeration through a bitwise OR, in order to set behavior for normal
	 * cases.
	 *
	 * @code{.c}
	 * av_rescale_rnd(3, 1, 2, AV_ROUND_UP | AV_ROUND_PASS_MINMAX);
	 * // Rescaling 3:
	 * //     Calculating 3 * 1 / 2
	 * //     3 / 2 is rounded up to 2
	 * //     => 2
	 *
	 * av_rescale_rnd(AV_NOPTS_VALUE, 1, 2, AV_ROUND_UP | AV_ROUND_PASS_MINMAX);
	 * // Rescaling AV_NOPTS_VALUE:
	 * //     AV_NOPTS_VALUE == INT64_MIN
	 * //     AV_NOPTS_VALUE is passed through
	 * //     => AV_NOPTS_VALUE
	 * @endcode
	 */
	AvRoundPassMinmax = C.AV_ROUND_PASS_MINMAX
)
