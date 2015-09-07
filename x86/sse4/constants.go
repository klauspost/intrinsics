package sse4

/* These specify the type of data that we're comparing.  */
const UBYTE_OPS = 0x00 // a and b contain strings of unsigned 8-bit characters.
const UWORD_OPS = 0x01 // a and b contain strings of unsigned 16-bit characters.
const SBYTE_OPS = 0x02 // a and b contain strings of signed 8-bit characters.
const SWORD_OPS = 0x03 // a and b contain strings of signed 16-bit characters.

/* These specify the type of comparison operation.  */
const CMP_EQUAL_ANY = 0x00     // For each character c in a, determine whether any character in b is equal to c. The first operand is a character set, the second is a string (think of strspn or strcspn).
const CMP_RANGES = 0x04        // For each character c in a, determine whether b0 <= c <= b1or b2 <= c <= b3… The first operand consists of ranges, for example, "azAZ" means "all characters from a to z and all characters from A to Z
const CMP_EQUAL_EACH = 0x08    // This implements the string equality algorithm. This operation compares two strings (think of strcmp or memcmp). The result of comparison is a bit mask (1 if the corresponding bytes are equal, 0 if not equal).
const CMP_EQUAL_ORDERED = 0x0c // This implements the substring search algorithm.  The first operand contains a string to search for, the second is a string to search in. The bit mask includes 1 if the substring is found at the corresponding position:

/* These macros specify the polarity of the operation.  */
const POSITIVE_POLARITY = 0x00        // No effect
const NEGATIVE_POLARITY = 0x10        // Negation of resulting bitmask.
const MASKED_POSITIVE_POLARITY = 0x20 //
const MASKED_NEGATIVE_POLARITY = 0x30 // Negation of resulting bitmask except for bits that have an index larger than the size of a or b (see details of pcmpestri instruction).

/* These macros are used in CmpXstri() to specify the return.  */
const LEAST_SIGNIFICANT = 0x00 //  The index of the rightmost bit set to 1 is returned.
const MOST_SIGNIFICANT = 0x40  // The index of the leftmost bit set to 1 is returned.

/* These macros are used in CmpXstri() to specify the return.  */
const BIT_MASK = 0x00
const UNIT_MASK = 0x40
