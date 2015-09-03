# 1 "test.c"
# 1 "<command-line>"
# 1 "test.c"
# 1 "arm_neon.h" 1
# 30 "arm_neon.h"
#pragma GCC push_options
#pragma GCC target ("+nothing+simd")

# 1 "c:\\tdm-gcc-64\\lib\\gcc\\x86_64-w64-mingw32\\4.8.1\\include\\stdint.h" 1 3 4
# 9 "c:\\tdm-gcc-64\\lib\\gcc\\x86_64-w64-mingw32\\4.8.1\\include\\stdint.h" 3 4
# 1 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\stdint.h" 1 3 4
# 28 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\stdint.h" 3 4
# 1 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\crtdefs.h" 1 3 4
# 10 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\crtdefs.h" 3 4
# 1 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\_mingw.h" 1 3 4
# 12 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\_mingw.h" 3 4
# 1 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\_mingw_mac.h" 1 3 4
# 46 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\_mingw_mac.h" 3 4
             
# 55 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\_mingw_mac.h" 3 4
             
# 13 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\_mingw.h" 2 3 4
# 1 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\_mingw_secapi.h" 1 3 4
# 14 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\_mingw.h" 2 3 4
# 282 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\_mingw.h" 3 4
# 1 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\vadefs.h" 1 3 4
# 9 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\vadefs.h" 3 4
# 1 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\_mingw.h" 1 3 4
# 686 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\_mingw.h" 3 4
# 1 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\sdks/_mingw_directx.h" 1 3 4
# 687 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\_mingw.h" 2 3 4
# 1 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\sdks/_mingw_ddk.h" 1 3 4
# 688 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\_mingw.h" 2 3 4
# 10 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\vadefs.h" 2 3 4
#pragma pack(push,_CRT_PACKING)
# 22 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\vadefs.h" 3 4
  typedef builtin_va_list gnuc_va_list;


  typedef gnuc_va_list va_list;
# 101 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\vadefs.h" 3 4
#pragma pack(pop)
# 283 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\_mingw.h" 2 3 4
#pragma pack(push,_CRT_PACKING)
# 377 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\_mingw.h" 3 4
extension typedef unsigned long long size_t;
# 387 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\_mingw.h" 3 4
extension typedef long long ssize_t;
# 399 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\_mingw.h" 3 4
extension typedef long long intptr_t;
# 412 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\_mingw.h" 3 4
extension typedef unsigned long long uintptr_t;
# 425 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\_mingw.h" 3 4
extension typedef long long ptrdiff_t;
# 435 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\_mingw.h" 3 4
typedef unsigned short wchar_t;



typedef unsigned short wint_t;
typedef unsigned short wctype_t;
# 463 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\_mingw.h" 3 4
typedef int errno_t;


typedef long time32_t;


extension typedef long long time64_t;



typedef time64_t time_t;
# 656 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\_mingw.h" 3 4
void attribute((cdecl)) debugbreak(void);
extern inline attribute((always_inline,gnu_inline)) void attribute((cdecl)) debugbreak(void)
{
  asm volatile("int {$}3":);
}


const char *mingw_get_crt_info (void);


#pragma pack(pop)
# 11 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\crtdefs.h" 2 3 4
# 26 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\crtdefs.h" 3 4
typedef size_t rsize_t;
# 153 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\crtdefs.h" 3 4
struct threadlocaleinfostruct;
struct threadmbcinfostruct;
typedef struct threadlocaleinfostruct *pthreadlocinfo;
typedef struct threadmbcinfostruct *pthreadmbcinfo;
struct lc_time_data;

typedef struct localeinfo_struct {
  pthreadlocinfo locinfo;
  pthreadmbcinfo mbcinfo;
} _locale_tstruct,*_locale_t;

typedef struct tagLC_ID {
  unsigned short wLanguage;
  unsigned short wCountry;
  unsigned short wCodePage;
} LC_ID,*LPLC_ID;


typedef struct threadlocaleinfostruct {
  int refcount;
  unsigned int lc_codepage;
  unsigned int lc_collate_cp;
  unsigned long lc_handle[6];
  LC_ID lc_id[6];
  struct {
    char *locale;
    wchar_t *wlocale;
    int *refcount;
    int *wrefcount;
  } lc_category[6];
  int lc_clike;
  int mb_cur_max;
  int *lconv_intl_refcount;
  int *lconv_num_refcount;
  int *lconv_mon_refcount;
  struct lconv *lconv;
  int *ctype1_refcount;
  unsigned short *ctype1;
  const unsigned short *pctype;
  const unsigned char *pclmap;
  const unsigned char *pcumap;
  struct lc_time_data *lc_time_curr;
} threadlocinfo;
# 29 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\stdint.h" 2 3 4

# 1 "c:\\tdm-gcc-64\\lib\\gcc\\x86_64-w64-mingw32\\4.8.1\\include\\stddef.h" 1 3 4
# 1 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\stddef.h" 1 3 4
# 18 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\stddef.h" 3 4
  attribute ((dllimport)) extern int *attribute((cdecl)) _errno(void);

  errno_t attribute((cdecl)) _set_errno(int _Value);
  errno_t attribute((cdecl)) _get_errno(int *_Value);
  attribute ((dllimport)) extern unsigned long attribute((cdecl)) threadid(void);

  attribute ((dllimport)) extern uintptr_t attribute((cdecl)) threadhandle(void);
# 2 "c:\\tdm-gcc-64\\lib\\gcc\\x86_64-w64-mingw32\\4.8.1\\include\\stddef.h" 2 3 4
# 33 "c:\\tdm-gcc-64\\x86_64-w64-mingw32\\include\\stdint.h" 2 3 4
typedef signed char int8_t;
typedef unsigned char uint8_t;
typedef short int16_t;
typedef unsigned short uint16_t;
typedef int int32_t;
typedef unsigned uint32_t;
extension typedef long long int64_t;
extension typedef unsigned long long uint64_t;
typedef signed char int_least8_t;
typedef unsigned char uint_least8_t;
typedef short int_least16_t;
typedef unsigned short uint_least16_t;
typedef int int_least32_t;
typedef unsigned uint_least32_t;
extension typedef long long int_least64_t;
extension typedef unsigned long long uint_least64_t;

typedef signed char int_fast8_t;
typedef unsigned char uint_fast8_t;
typedef short int_fast16_t;
typedef unsigned short uint_fast16_t;
typedef int int_fast32_t;
typedef unsigned int uint_fast32_t;
extension typedef long long int_fast64_t;
extension typedef unsigned long long uint_fast64_t;
extension typedef long long intmax_t;
extension typedef unsigned long long uintmax_t;
# 10 "c:\\tdm-gcc-64\\lib\\gcc\\x86_64-w64-mingw32\\4.8.1\\include\\stdint.h" 2 3 4
# 34 "arm_neon.h" 2


typedef Int8x8_t int8x8_t;
typedef Int16x4_t int16x4_t;
typedef Int32x2_t int32x2_t;
typedef Int64x1_t int64x1_t;
typedef Float32x2_t float32x2_t;
typedef Poly8x8_t poly8x8_t;
typedef Poly16x4_t poly16x4_t;
typedef Uint8x8_t uint8x8_t;
typedef Uint16x4_t uint16x4_t;
typedef Uint32x2_t uint32x2_t;
typedef Float64x1_t float64x1_t;
typedef Uint64x1_t uint64x1_t;
typedef Int8x16_t int8x16_t;
typedef Int16x8_t int16x8_t;
typedef Int32x4_t int32x4_t;
typedef Int64x2_t int64x2_t;
typedef Float32x4_t float32x4_t;
typedef Float64x2_t float64x2_t;
typedef Poly8x16_t poly8x16_t;
typedef Poly16x8_t poly16x8_t;
typedef Poly64x2_t poly64x2_t;
typedef Uint8x16_t uint8x16_t;
typedef Uint16x8_t uint16x8_t;
typedef Uint32x4_t uint32x4_t;
typedef Uint64x2_t uint64x2_t;

typedef Poly8_t poly8_t;
typedef Poly16_t poly16_t;
typedef Poly64_t poly64_t;
typedef Poly128_t poly128_t;

typedef float float32_t;
typedef double float64_t;

typedef struct int8x8x2_t
{
  int8x8_t val[2];
} int8x8x2_t;

typedef struct int8x16x2_t
{
  int8x16_t val[2];
} int8x16x2_t;

typedef struct int16x4x2_t
{
  int16x4_t val[2];
} int16x4x2_t;

typedef struct int16x8x2_t
{
  int16x8_t val[2];
} int16x8x2_t;

typedef struct int32x2x2_t
{
  int32x2_t val[2];
} int32x2x2_t;

typedef struct int32x4x2_t
{
  int32x4_t val[2];
} int32x4x2_t;

typedef struct int64x1x2_t
{
  int64x1_t val[2];
} int64x1x2_t;

typedef struct int64x2x2_t
{
  int64x2_t val[2];
} int64x2x2_t;

typedef struct uint8x8x2_t
{
  uint8x8_t val[2];
} uint8x8x2_t;

typedef struct uint8x16x2_t
{
  uint8x16_t val[2];
} uint8x16x2_t;

typedef struct uint16x4x2_t
{
  uint16x4_t val[2];
} uint16x4x2_t;

typedef struct uint16x8x2_t
{
  uint16x8_t val[2];
} uint16x8x2_t;

typedef struct uint32x2x2_t
{
  uint32x2_t val[2];
} uint32x2x2_t;

typedef struct uint32x4x2_t
{
  uint32x4_t val[2];
} uint32x4x2_t;

typedef struct uint64x1x2_t
{
  uint64x1_t val[2];
} uint64x1x2_t;

typedef struct uint64x2x2_t
{
  uint64x2_t val[2];
} uint64x2x2_t;

typedef struct float32x2x2_t
{
  float32x2_t val[2];
} float32x2x2_t;

typedef struct float32x4x2_t
{
  float32x4_t val[2];
} float32x4x2_t;

typedef struct float64x2x2_t
{
  float64x2_t val[2];
} float64x2x2_t;

typedef struct float64x1x2_t
{
  float64x1_t val[2];
} float64x1x2_t;

typedef struct poly8x8x2_t
{
  poly8x8_t val[2];
} poly8x8x2_t;

typedef struct poly8x16x2_t
{
  poly8x16_t val[2];
} poly8x16x2_t;

typedef struct poly16x4x2_t
{
  poly16x4_t val[2];
} poly16x4x2_t;

typedef struct poly16x8x2_t
{
  poly16x8_t val[2];
} poly16x8x2_t;

typedef struct int8x8x3_t
{
  int8x8_t val[3];
} int8x8x3_t;

typedef struct int8x16x3_t
{
  int8x16_t val[3];
} int8x16x3_t;

typedef struct int16x4x3_t
{
  int16x4_t val[3];
} int16x4x3_t;

typedef struct int16x8x3_t
{
  int16x8_t val[3];
} int16x8x3_t;

typedef struct int32x2x3_t
{
  int32x2_t val[3];
} int32x2x3_t;

typedef struct int32x4x3_t
{
  int32x4_t val[3];
} int32x4x3_t;

typedef struct int64x1x3_t
{
  int64x1_t val[3];
} int64x1x3_t;

typedef struct int64x2x3_t
{
  int64x2_t val[3];
} int64x2x3_t;

typedef struct uint8x8x3_t
{
  uint8x8_t val[3];
} uint8x8x3_t;

typedef struct uint8x16x3_t
{
  uint8x16_t val[3];
} uint8x16x3_t;

typedef struct uint16x4x3_t
{
  uint16x4_t val[3];
} uint16x4x3_t;

typedef struct uint16x8x3_t
{
  uint16x8_t val[3];
} uint16x8x3_t;

typedef struct uint32x2x3_t
{
  uint32x2_t val[3];
} uint32x2x3_t;

typedef struct uint32x4x3_t
{
  uint32x4_t val[3];
} uint32x4x3_t;

typedef struct uint64x1x3_t
{
  uint64x1_t val[3];
} uint64x1x3_t;

typedef struct uint64x2x3_t
{
  uint64x2_t val[3];
} uint64x2x3_t;

typedef struct float32x2x3_t
{
  float32x2_t val[3];
} float32x2x3_t;

typedef struct float32x4x3_t
{
  float32x4_t val[3];
} float32x4x3_t;

typedef struct float64x2x3_t
{
  float64x2_t val[3];
} float64x2x3_t;

typedef struct float64x1x3_t
{
  float64x1_t val[3];
} float64x1x3_t;

typedef struct poly8x8x3_t
{
  poly8x8_t val[3];
} poly8x8x3_t;

typedef struct poly8x16x3_t
{
  poly8x16_t val[3];
} poly8x16x3_t;

typedef struct poly16x4x3_t
{
  poly16x4_t val[3];
} poly16x4x3_t;

typedef struct poly16x8x3_t
{
  poly16x8_t val[3];
} poly16x8x3_t;

typedef struct int8x8x4_t
{
  int8x8_t val[4];
} int8x8x4_t;

typedef struct int8x16x4_t
{
  int8x16_t val[4];
} int8x16x4_t;

typedef struct int16x4x4_t
{
  int16x4_t val[4];
} int16x4x4_t;

typedef struct int16x8x4_t
{
  int16x8_t val[4];
} int16x8x4_t;

typedef struct int32x2x4_t
{
  int32x2_t val[4];
} int32x2x4_t;

typedef struct int32x4x4_t
{
  int32x4_t val[4];
} int32x4x4_t;

typedef struct int64x1x4_t
{
  int64x1_t val[4];
} int64x1x4_t;

typedef struct int64x2x4_t
{
  int64x2_t val[4];
} int64x2x4_t;

typedef struct uint8x8x4_t
{
  uint8x8_t val[4];
} uint8x8x4_t;

typedef struct uint8x16x4_t
{
  uint8x16_t val[4];
} uint8x16x4_t;

typedef struct uint16x4x4_t
{
  uint16x4_t val[4];
} uint16x4x4_t;

typedef struct uint16x8x4_t
{
  uint16x8_t val[4];
} uint16x8x4_t;

typedef struct uint32x2x4_t
{
  uint32x2_t val[4];
} uint32x2x4_t;

typedef struct uint32x4x4_t
{
  uint32x4_t val[4];
} uint32x4x4_t;

typedef struct uint64x1x4_t
{
  uint64x1_t val[4];
} uint64x1x4_t;

typedef struct uint64x2x4_t
{
  uint64x2_t val[4];
} uint64x2x4_t;

typedef struct float32x2x4_t
{
  float32x2_t val[4];
} float32x2x4_t;

typedef struct float32x4x4_t
{
  float32x4_t val[4];
} float32x4x4_t;

typedef struct float64x2x4_t
{
  float64x2_t val[4];
} float64x2x4_t;

typedef struct float64x1x4_t
{
  float64x1_t val[4];
} float64x1x4_t;

typedef struct poly8x8x4_t
{
  poly8x8_t val[4];
} poly8x8x4_t;

typedef struct poly8x16x4_t
{
  poly8x16_t val[4];
} poly8x16x4_t;

typedef struct poly16x4x4_t
{
  poly16x4_t val[4];
} poly16x4x4_t;

typedef struct poly16x8x4_t
{
  poly16x8_t val[4];
} poly16x8x4_t;
# 571 "arm_neon.h"
DEF: int8x8
vadd_s8 (int8x8_t a, int8x8_t b)
{
  return a + b;
}

DEF: int16x4
vadd_s16 (int16x4_t a, int16x4_t b)
{
  return a + b;
}

DEF: int32x2
vadd_s32 (int32x2_t a, int32x2_t b)
{
  return a + b;
}

DEF: float32x2
vadd_f32 (float32x2_t a, float32x2_t b)
{
  return a + b;
}

DEF: float64x1
vadd_f64 (float64x1_t a, float64x1_t b)
{
  return a + b;
}

DEF: uint8x8
vadd_u8 (uint8x8_t a, uint8x8_t b)
{
  return a + b;
}

DEF: uint16x4
vadd_u16 (uint16x4_t a, uint16x4_t b)
{
  return a + b;
}

DEF: uint32x2
vadd_u32 (uint32x2_t a, uint32x2_t b)
{
  return a + b;
}

DEF: int64x1
vadd_s64 (int64x1_t a, int64x1_t b)
{
  return a + b;
}

DEF: uint64x1
vadd_u64 (uint64x1_t a, uint64x1_t b)
{
  return a + b;
}

DEF: int8x16
vaddq_s8 (int8x16_t a, int8x16_t b)
{
  return a + b;
}

DEF: int16x8
vaddq_s16 (int16x8_t a, int16x8_t b)
{
  return a + b;
}

DEF: int32x4
vaddq_s32 (int32x4_t a, int32x4_t b)
{
  return a + b;
}

DEF: int64x2
vaddq_s64 (int64x2_t a, int64x2_t b)
{
  return a + b;
}

DEF: float32x4
vaddq_f32 (float32x4_t a, float32x4_t b)
{
  return a + b;
}

DEF: float64x2
vaddq_f64 (float64x2_t a, float64x2_t b)
{
  return a + b;
}

DEF: uint8x16
vaddq_u8 (uint8x16_t a, uint8x16_t b)
{
  return a + b;
}

DEF: uint16x8
vaddq_u16 (uint16x8_t a, uint16x8_t b)
{
  return a + b;
}

DEF: uint32x4
vaddq_u32 (uint32x4_t a, uint32x4_t b)
{
  return a + b;
}

DEF: uint64x2
vaddq_u64 (uint64x2_t a, uint64x2_t b)
{
  return a + b;
}

DEF: int16x8
vaddl_s8 (int8x8_t a, int8x8_t b)
{
  return (int16x8_t) builtin_aarch64_saddlv8qi (a, b);
}

DEF: int32x4
vaddl_s16 (int16x4_t a, int16x4_t b)
{
  return (int32x4_t) builtin_aarch64_saddlv4hi (a, b);
}

DEF: int64x2
vaddl_s32 (int32x2_t a, int32x2_t b)
{
  return (int64x2_t) builtin_aarch64_saddlv2si (a, b);
}

DEF: uint16x8
vaddl_u8 (uint8x8_t a, uint8x8_t b)
{
  return (uint16x8_t) builtin_aarch64_uaddlv8qi ((int8x8_t) a,
         (int8x8_t) b);
}

DEF: uint32x4
vaddl_u16 (uint16x4_t a, uint16x4_t b)
{
  return (uint32x4_t) builtin_aarch64_uaddlv4hi ((int16x4_t) a,
         (int16x4_t) b);
}

DEF: uint64x2
vaddl_u32 (uint32x2_t a, uint32x2_t b)
{
  return (uint64x2_t) builtin_aarch64_uaddlv2si ((int32x2_t) a,
         (int32x2_t) b);
}

DEF: int16x8
vaddl_high_s8 (int8x16_t a, int8x16_t b)
{
  return (int16x8_t) builtin_aarch64_saddl2v16qi (a, b);
}

DEF: int32x4
vaddl_high_s16 (int16x8_t a, int16x8_t b)
{
  return (int32x4_t) builtin_aarch64_saddl2v8hi (a, b);
}

DEF: int64x2
vaddl_high_s32 (int32x4_t a, int32x4_t b)
{
  return (int64x2_t) builtin_aarch64_saddl2v4si (a, b);
}

DEF: uint16x8
vaddl_high_u8 (uint8x16_t a, uint8x16_t b)
{
  return (uint16x8_t) builtin_aarch64_uaddl2v16qi ((int8x16_t) a,
           (int8x16_t) b);
}

DEF: uint32x4
vaddl_high_u16 (uint16x8_t a, uint16x8_t b)
{
  return (uint32x4_t) builtin_aarch64_uaddl2v8hi ((int16x8_t) a,
          (int16x8_t) b);
}

DEF: uint64x2
vaddl_high_u32 (uint32x4_t a, uint32x4_t b)
{
  return (uint64x2_t) builtin_aarch64_uaddl2v4si ((int32x4_t) a,
          (int32x4_t) b);
}

DEF: int16x8
vaddw_s8 (int16x8_t a, int8x8_t b)
{
  return (int16x8_t) builtin_aarch64_saddwv8qi (a, b);
}

DEF: int32x4
vaddw_s16 (int32x4_t a, int16x4_t b)
{
  return (int32x4_t) builtin_aarch64_saddwv4hi (a, b);
}

DEF: int64x2
vaddw_s32 (int64x2_t a, int32x2_t b)
{
  return (int64x2_t) builtin_aarch64_saddwv2si (a, b);
}

DEF: uint16x8
vaddw_u8 (uint16x8_t a, uint8x8_t b)
{
  return (uint16x8_t) builtin_aarch64_uaddwv8qi ((int16x8_t) a,
         (int8x8_t) b);
}

DEF: uint32x4
vaddw_u16 (uint32x4_t a, uint16x4_t b)
{
  return (uint32x4_t) builtin_aarch64_uaddwv4hi ((int32x4_t) a,
         (int16x4_t) b);
}

DEF: uint64x2
vaddw_u32 (uint64x2_t a, uint32x2_t b)
{
  return (uint64x2_t) builtin_aarch64_uaddwv2si ((int64x2_t) a,
         (int32x2_t) b);
}

DEF: int16x8
vaddw_high_s8 (int16x8_t a, int8x16_t b)
{
  return (int16x8_t) builtin_aarch64_saddw2v16qi (a, b);
}

DEF: int32x4
vaddw_high_s16 (int32x4_t a, int16x8_t b)
{
  return (int32x4_t) builtin_aarch64_saddw2v8hi (a, b);
}

DEF: int64x2
vaddw_high_s32 (int64x2_t a, int32x4_t b)
{
  return (int64x2_t) builtin_aarch64_saddw2v4si (a, b);
}

DEF: uint16x8
vaddw_high_u8 (uint16x8_t a, uint8x16_t b)
{
  return (uint16x8_t) builtin_aarch64_uaddw2v16qi ((int16x8_t) a,
           (int8x16_t) b);
}

DEF: uint32x4
vaddw_high_u16 (uint32x4_t a, uint16x8_t b)
{
  return (uint32x4_t) builtin_aarch64_uaddw2v8hi ((int32x4_t) a,
          (int16x8_t) b);
}

DEF: uint64x2
vaddw_high_u32 (uint64x2_t a, uint32x4_t b)
{
  return (uint64x2_t) builtin_aarch64_uaddw2v4si ((int64x2_t) a,
          (int32x4_t) b);
}

DEF: int8x8
vhadd_s8 (int8x8_t a, int8x8_t b)
{
  return (int8x8_t) builtin_aarch64_shaddv8qi (a, b);
}

DEF: int16x4
vhadd_s16 (int16x4_t a, int16x4_t b)
{
  return (int16x4_t) builtin_aarch64_shaddv4hi (a, b);
}

DEF: int32x2
vhadd_s32 (int32x2_t a, int32x2_t b)
{
  return (int32x2_t) builtin_aarch64_shaddv2si (a, b);
}

DEF: uint8x8
vhadd_u8 (uint8x8_t a, uint8x8_t b)
{
  return (uint8x8_t) builtin_aarch64_uhaddv8qi ((int8x8_t) a,
        (int8x8_t) b);
}

DEF: uint16x4
vhadd_u16 (uint16x4_t a, uint16x4_t b)
{
  return (uint16x4_t) builtin_aarch64_uhaddv4hi ((int16x4_t) a,
         (int16x4_t) b);
}

DEF: uint32x2
vhadd_u32 (uint32x2_t a, uint32x2_t b)
{
  return (uint32x2_t) builtin_aarch64_uhaddv2si ((int32x2_t) a,
         (int32x2_t) b);
}

DEF: int8x16
vhaddq_s8 (int8x16_t a, int8x16_t b)
{
  return (int8x16_t) builtin_aarch64_shaddv16qi (a, b);
}

DEF: int16x8
vhaddq_s16 (int16x8_t a, int16x8_t b)
{
  return (int16x8_t) builtin_aarch64_shaddv8hi (a, b);
}

DEF: int32x4
vhaddq_s32 (int32x4_t a, int32x4_t b)
{
  return (int32x4_t) builtin_aarch64_shaddv4si (a, b);
}

DEF: uint8x16
vhaddq_u8 (uint8x16_t a, uint8x16_t b)
{
  return (uint8x16_t) builtin_aarch64_uhaddv16qi ((int8x16_t) a,
          (int8x16_t) b);
}

DEF: uint16x8
vhaddq_u16 (uint16x8_t a, uint16x8_t b)
{
  return (uint16x8_t) builtin_aarch64_uhaddv8hi ((int16x8_t) a,
         (int16x8_t) b);
}

DEF: uint32x4
vhaddq_u32 (uint32x4_t a, uint32x4_t b)
{
  return (uint32x4_t) builtin_aarch64_uhaddv4si ((int32x4_t) a,
         (int32x4_t) b);
}

DEF: int8x8
vrhadd_s8 (int8x8_t a, int8x8_t b)
{
  return (int8x8_t) builtin_aarch64_srhaddv8qi (a, b);
}

DEF: int16x4
vrhadd_s16 (int16x4_t a, int16x4_t b)
{
  return (int16x4_t) builtin_aarch64_srhaddv4hi (a, b);
}

DEF: int32x2
vrhadd_s32 (int32x2_t a, int32x2_t b)
{
  return (int32x2_t) builtin_aarch64_srhaddv2si (a, b);
}

DEF: uint8x8
vrhadd_u8 (uint8x8_t a, uint8x8_t b)
{
  return (uint8x8_t) builtin_aarch64_urhaddv8qi ((int8x8_t) a,
         (int8x8_t) b);
}

DEF: uint16x4
vrhadd_u16 (uint16x4_t a, uint16x4_t b)
{
  return (uint16x4_t) builtin_aarch64_urhaddv4hi ((int16x4_t) a,
          (int16x4_t) b);
}

DEF: uint32x2
vrhadd_u32 (uint32x2_t a, uint32x2_t b)
{
  return (uint32x2_t) builtin_aarch64_urhaddv2si ((int32x2_t) a,
          (int32x2_t) b);
}

DEF: int8x16
vrhaddq_s8 (int8x16_t a, int8x16_t b)
{
  return (int8x16_t) builtin_aarch64_srhaddv16qi (a, b);
}

DEF: int16x8
vrhaddq_s16 (int16x8_t a, int16x8_t b)
{
  return (int16x8_t) builtin_aarch64_srhaddv8hi (a, b);
}

DEF: int32x4
vrhaddq_s32 (int32x4_t a, int32x4_t b)
{
  return (int32x4_t) builtin_aarch64_srhaddv4si (a, b);
}

DEF: uint8x16
vrhaddq_u8 (uint8x16_t a, uint8x16_t b)
{
  return (uint8x16_t) builtin_aarch64_urhaddv16qi ((int8x16_t) a,
           (int8x16_t) b);
}

DEF: uint16x8
vrhaddq_u16 (uint16x8_t a, uint16x8_t b)
{
  return (uint16x8_t) builtin_aarch64_urhaddv8hi ((int16x8_t) a,
          (int16x8_t) b);
}

DEF: uint32x4
vrhaddq_u32 (uint32x4_t a, uint32x4_t b)
{
  return (uint32x4_t) builtin_aarch64_urhaddv4si ((int32x4_t) a,
          (int32x4_t) b);
}

DEF: int8x8
vaddhn_s16 (int16x8_t a, int16x8_t b)
{
  return (int8x8_t) builtin_aarch64_addhnv8hi (a, b);
}

DEF: int16x4
vaddhn_s32 (int32x4_t a, int32x4_t b)
{
  return (int16x4_t) builtin_aarch64_addhnv4si (a, b);
}

DEF: int32x2
vaddhn_s64 (int64x2_t a, int64x2_t b)
{
  return (int32x2_t) builtin_aarch64_addhnv2di (a, b);
}

DEF: uint8x8
vaddhn_u16 (uint16x8_t a, uint16x8_t b)
{
  return (uint8x8_t) builtin_aarch64_addhnv8hi ((int16x8_t) a,
        (int16x8_t) b);
}

DEF: uint16x4
vaddhn_u32 (uint32x4_t a, uint32x4_t b)
{
  return (uint16x4_t) builtin_aarch64_addhnv4si ((int32x4_t) a,
         (int32x4_t) b);
}

DEF: uint32x2
vaddhn_u64 (uint64x2_t a, uint64x2_t b)
{
  return (uint32x2_t) builtin_aarch64_addhnv2di ((int64x2_t) a,
         (int64x2_t) b);
}

DEF: int8x8
vraddhn_s16 (int16x8_t a, int16x8_t b)
{
  return (int8x8_t) builtin_aarch64_raddhnv8hi (a, b);
}

DEF: int16x4
vraddhn_s32 (int32x4_t a, int32x4_t b)
{
  return (int16x4_t) builtin_aarch64_raddhnv4si (a, b);
}

DEF: int32x2
vraddhn_s64 (int64x2_t a, int64x2_t b)
{
  return (int32x2_t) builtin_aarch64_raddhnv2di (a, b);
}

DEF: uint8x8
vraddhn_u16 (uint16x8_t a, uint16x8_t b)
{
  return (uint8x8_t) builtin_aarch64_raddhnv8hi ((int16x8_t) a,
         (int16x8_t) b);
}

DEF: uint16x4
vraddhn_u32 (uint32x4_t a, uint32x4_t b)
{
  return (uint16x4_t) builtin_aarch64_raddhnv4si ((int32x4_t) a,
          (int32x4_t) b);
}

DEF: uint32x2
vraddhn_u64 (uint64x2_t a, uint64x2_t b)
{
  return (uint32x2_t) builtin_aarch64_raddhnv2di ((int64x2_t) a,
          (int64x2_t) b);
}

DEF: int8x16
vaddhn_high_s16 (int8x8_t a, int16x8_t b, int16x8_t c)
{
  return (int8x16_t) builtin_aarch64_addhn2v8hi (a, b, c);
}

DEF: int16x8
vaddhn_high_s32 (int16x4_t a, int32x4_t b, int32x4_t c)
{
  return (int16x8_t) builtin_aarch64_addhn2v4si (a, b, c);
}

DEF: int32x4
vaddhn_high_s64 (int32x2_t a, int64x2_t b, int64x2_t c)
{
  return (int32x4_t) builtin_aarch64_addhn2v2di (a, b, c);
}

DEF: uint8x16
vaddhn_high_u16 (uint8x8_t a, uint16x8_t b, uint16x8_t c)
{
  return (uint8x16_t) builtin_aarch64_addhn2v8hi ((int8x8_t) a,
          (int16x8_t) b,
          (int16x8_t) c);
}

DEF: uint16x8
vaddhn_high_u32 (uint16x4_t a, uint32x4_t b, uint32x4_t c)
{
  return (uint16x8_t) builtin_aarch64_addhn2v4si ((int16x4_t) a,
          (int32x4_t) b,
          (int32x4_t) c);
}

DEF: uint32x4
vaddhn_high_u64 (uint32x2_t a, uint64x2_t b, uint64x2_t c)
{
  return (uint32x4_t) builtin_aarch64_addhn2v2di ((int32x2_t) a,
          (int64x2_t) b,
          (int64x2_t) c);
}

DEF: int8x16
vraddhn_high_s16 (int8x8_t a, int16x8_t b, int16x8_t c)
{
  return (int8x16_t) builtin_aarch64_raddhn2v8hi (a, b, c);
}

DEF: int16x8
vraddhn_high_s32 (int16x4_t a, int32x4_t b, int32x4_t c)
{
  return (int16x8_t) builtin_aarch64_raddhn2v4si (a, b, c);
}

DEF: int32x4
vraddhn_high_s64 (int32x2_t a, int64x2_t b, int64x2_t c)
{
  return (int32x4_t) builtin_aarch64_raddhn2v2di (a, b, c);
}

DEF: uint8x16
vraddhn_high_u16 (uint8x8_t a, uint16x8_t b, uint16x8_t c)
{
  return (uint8x16_t) builtin_aarch64_raddhn2v8hi ((int8x8_t) a,
           (int16x8_t) b,
           (int16x8_t) c);
}

DEF: uint16x8
vraddhn_high_u32 (uint16x4_t a, uint32x4_t b, uint32x4_t c)
{
  return (uint16x8_t) builtin_aarch64_raddhn2v4si ((int16x4_t) a,
           (int32x4_t) b,
           (int32x4_t) c);
}

DEF: uint32x4
vraddhn_high_u64 (uint32x2_t a, uint64x2_t b, uint64x2_t c)
{
  return (uint32x4_t) builtin_aarch64_raddhn2v2di ((int32x2_t) a,
           (int64x2_t) b,
           (int64x2_t) c);
}

DEF: float32x2
vdiv_f32 (float32x2_t a, float32x2_t b)
{
  return a / b;
}

DEF: float64x1
vdiv_f64 (float64x1_t a, float64x1_t b)
{
  return a / b;
}

DEF: float32x4
vdivq_f32 (float32x4_t a, float32x4_t b)
{
  return a / b;
}

DEF: float64x2
vdivq_f64 (float64x2_t a, float64x2_t b)
{
  return a / b;
}

DEF: int8x8
vmul_s8 (int8x8_t a, int8x8_t b)
{
  return a * b;
}

DEF: int16x4
vmul_s16 (int16x4_t a, int16x4_t b)
{
  return a * b;
}

DEF: int32x2
vmul_s32 (int32x2_t a, int32x2_t b)
{
  return a * b;
}

DEF: float32x2
vmul_f32 (float32x2_t a, float32x2_t b)
{
  return a * b;
}

DEF: float64x1
vmul_f64 (float64x1_t a, float64x1_t b)
{
  return a * b;
}

DEF: uint8x8
vmul_u8 (uint8x8_t a, uint8x8_t b)
{
  return a * b;
}

DEF: uint16x4
vmul_u16 (uint16x4_t a, uint16x4_t b)
{
  return a * b;
}

DEF: uint32x2
vmul_u32 (uint32x2_t a, uint32x2_t b)
{
  return a * b;
}

DEF: poly8x8
vmul_p8 (poly8x8_t a, poly8x8_t b)
{
  return (poly8x8_t) builtin_aarch64_pmulv8qi ((int8x8_t) a,
       (int8x8_t) b);
}

DEF: int8x16
vmulq_s8 (int8x16_t a, int8x16_t b)
{
  return a * b;
}

DEF: int16x8
vmulq_s16 (int16x8_t a, int16x8_t b)
{
  return a * b;
}

DEF: int32x4
vmulq_s32 (int32x4_t a, int32x4_t b)
{
  return a * b;
}

DEF: float32x4
vmulq_f32 (float32x4_t a, float32x4_t b)
{
  return a * b;
}

DEF: float64x2
vmulq_f64 (float64x2_t a, float64x2_t b)
{
  return a * b;
}

DEF: uint8x16
vmulq_u8 (uint8x16_t a, uint8x16_t b)
{
  return a * b;
}

DEF: uint16x8
vmulq_u16 (uint16x8_t a, uint16x8_t b)
{
  return a * b;
}

DEF: uint32x4
vmulq_u32 (uint32x4_t a, uint32x4_t b)
{
  return a * b;
}

DEF: poly8x16
vmulq_p8 (poly8x16_t a, poly8x16_t b)
{
  return (poly8x16_t) builtin_aarch64_pmulv16qi ((int8x16_t) a,
         (int8x16_t) b);
}

DEF: int8x8
vand_s8 (int8x8_t a, int8x8_t b)
{
  return a & b;
}

DEF: int16x4
vand_s16 (int16x4_t a, int16x4_t b)
{
  return a & b;
}

DEF: int32x2
vand_s32 (int32x2_t a, int32x2_t b)
{
  return a & b;
}

DEF: uint8x8
vand_u8 (uint8x8_t a, uint8x8_t b)
{
  return a & b;
}

DEF: uint16x4
vand_u16 (uint16x4_t a, uint16x4_t b)
{
  return a & b;
}

DEF: uint32x2
vand_u32 (uint32x2_t a, uint32x2_t b)
{
  return a & b;
}

DEF: int64x1
vand_s64 (int64x1_t a, int64x1_t b)
{
  return a & b;
}

DEF: uint64x1
vand_u64 (uint64x1_t a, uint64x1_t b)
{
  return a & b;
}

DEF: int8x16
vandq_s8 (int8x16_t a, int8x16_t b)
{
  return a & b;
}

DEF: int16x8
vandq_s16 (int16x8_t a, int16x8_t b)
{
  return a & b;
}

DEF: int32x4
vandq_s32 (int32x4_t a, int32x4_t b)
{
  return a & b;
}

DEF: int64x2
vandq_s64 (int64x2_t a, int64x2_t b)
{
  return a & b;
}

DEF: uint8x16
vandq_u8 (uint8x16_t a, uint8x16_t b)
{
  return a & b;
}

DEF: uint16x8
vandq_u16 (uint16x8_t a, uint16x8_t b)
{
  return a & b;
}

DEF: uint32x4
vandq_u32 (uint32x4_t a, uint32x4_t b)
{
  return a & b;
}

DEF: uint64x2
vandq_u64 (uint64x2_t a, uint64x2_t b)
{
  return a & b;
}

DEF: int8x8
vorr_s8 (int8x8_t a, int8x8_t b)
{
  return a | b;
}

DEF: int16x4
vorr_s16 (int16x4_t a, int16x4_t b)
{
  return a | b;
}

DEF: int32x2
vorr_s32 (int32x2_t a, int32x2_t b)
{
  return a | b;
}

DEF: uint8x8
vorr_u8 (uint8x8_t a, uint8x8_t b)
{
  return a | b;
}

DEF: uint16x4
vorr_u16 (uint16x4_t a, uint16x4_t b)
{
  return a | b;
}

DEF: uint32x2
vorr_u32 (uint32x2_t a, uint32x2_t b)
{
  return a | b;
}

DEF: int64x1
vorr_s64 (int64x1_t a, int64x1_t b)
{
  return a | b;
}

DEF: uint64x1
vorr_u64 (uint64x1_t a, uint64x1_t b)
{
  return a | b;
}

DEF: int8x16
vorrq_s8 (int8x16_t a, int8x16_t b)
{
  return a | b;
}

DEF: int16x8
vorrq_s16 (int16x8_t a, int16x8_t b)
{
  return a | b;
}

DEF: int32x4
vorrq_s32 (int32x4_t a, int32x4_t b)
{
  return a | b;
}

DEF: int64x2
vorrq_s64 (int64x2_t a, int64x2_t b)
{
  return a | b;
}

DEF: uint8x16
vorrq_u8 (uint8x16_t a, uint8x16_t b)
{
  return a | b;
}

DEF: uint16x8
vorrq_u16 (uint16x8_t a, uint16x8_t b)
{
  return a | b;
}

DEF: uint32x4
vorrq_u32 (uint32x4_t a, uint32x4_t b)
{
  return a | b;
}

DEF: uint64x2
vorrq_u64 (uint64x2_t a, uint64x2_t b)
{
  return a | b;
}

DEF: int8x8
veor_s8 (int8x8_t a, int8x8_t b)
{
  return a ^ b;
}

DEF: int16x4
veor_s16 (int16x4_t a, int16x4_t b)
{
  return a ^ b;
}

DEF: int32x2
veor_s32 (int32x2_t a, int32x2_t b)
{
  return a ^ b;
}

DEF: uint8x8
veor_u8 (uint8x8_t a, uint8x8_t b)
{
  return a ^ b;
}

DEF: uint16x4
veor_u16 (uint16x4_t a, uint16x4_t b)
{
  return a ^ b;
}

DEF: uint32x2
veor_u32 (uint32x2_t a, uint32x2_t b)
{
  return a ^ b;
}

DEF: int64x1
veor_s64 (int64x1_t a, int64x1_t b)
{
  return a ^ b;
}

DEF: uint64x1
veor_u64 (uint64x1_t a, uint64x1_t b)
{
  return a ^ b;
}

DEF: int8x16
veorq_s8 (int8x16_t a, int8x16_t b)
{
  return a ^ b;
}

DEF: int16x8
veorq_s16 (int16x8_t a, int16x8_t b)
{
  return a ^ b;
}

DEF: int32x4
veorq_s32 (int32x4_t a, int32x4_t b)
{
  return a ^ b;
}

DEF: int64x2
veorq_s64 (int64x2_t a, int64x2_t b)
{
  return a ^ b;
}

DEF: uint8x16
veorq_u8 (uint8x16_t a, uint8x16_t b)
{
  return a ^ b;
}

DEF: uint16x8
veorq_u16 (uint16x8_t a, uint16x8_t b)
{
  return a ^ b;
}

DEF: uint32x4
veorq_u32 (uint32x4_t a, uint32x4_t b)
{
  return a ^ b;
}

DEF: uint64x2
veorq_u64 (uint64x2_t a, uint64x2_t b)
{
  return a ^ b;
}

DEF: int8x8
vbic_s8 (int8x8_t a, int8x8_t b)
{
  return a & ~b;
}

DEF: int16x4
vbic_s16 (int16x4_t a, int16x4_t b)
{
  return a & ~b;
}

DEF: int32x2
vbic_s32 (int32x2_t a, int32x2_t b)
{
  return a & ~b;
}

DEF: uint8x8
vbic_u8 (uint8x8_t a, uint8x8_t b)
{
  return a & ~b;
}

DEF: uint16x4
vbic_u16 (uint16x4_t a, uint16x4_t b)
{
  return a & ~b;
}

DEF: uint32x2
vbic_u32 (uint32x2_t a, uint32x2_t b)
{
  return a & ~b;
}

DEF: int64x1
vbic_s64 (int64x1_t a, int64x1_t b)
{
  return a & ~b;
}

DEF: uint64x1
vbic_u64 (uint64x1_t a, uint64x1_t b)
{
  return a & ~b;
}

DEF: int8x16
vbicq_s8 (int8x16_t a, int8x16_t b)
{
  return a & ~b;
}

DEF: int16x8
vbicq_s16 (int16x8_t a, int16x8_t b)
{
  return a & ~b;
}

DEF: int32x4
vbicq_s32 (int32x4_t a, int32x4_t b)
{
  return a & ~b;
}

DEF: int64x2
vbicq_s64 (int64x2_t a, int64x2_t b)
{
  return a & ~b;
}

DEF: uint8x16
vbicq_u8 (uint8x16_t a, uint8x16_t b)
{
  return a & ~b;
}

DEF: uint16x8
vbicq_u16 (uint16x8_t a, uint16x8_t b)
{
  return a & ~b;
}

DEF: uint32x4
vbicq_u32 (uint32x4_t a, uint32x4_t b)
{
  return a & ~b;
}

DEF: uint64x2
vbicq_u64 (uint64x2_t a, uint64x2_t b)
{
  return a & ~b;
}

DEF: int8x8
vorn_s8 (int8x8_t a, int8x8_t b)
{
  return a | ~b;
}

DEF: int16x4
vorn_s16 (int16x4_t a, int16x4_t b)
{
  return a | ~b;
}

DEF: int32x2
vorn_s32 (int32x2_t a, int32x2_t b)
{
  return a | ~b;
}

DEF: uint8x8
vorn_u8 (uint8x8_t a, uint8x8_t b)
{
  return a | ~b;
}

DEF: uint16x4
vorn_u16 (uint16x4_t a, uint16x4_t b)
{
  return a | ~b;
}

DEF: uint32x2
vorn_u32 (uint32x2_t a, uint32x2_t b)
{
  return a | ~b;
}

DEF: int64x1
vorn_s64 (int64x1_t a, int64x1_t b)
{
  return a | ~b;
}

DEF: uint64x1
vorn_u64 (uint64x1_t a, uint64x1_t b)
{
  return a | ~b;
}

DEF: int8x16
vornq_s8 (int8x16_t a, int8x16_t b)
{
  return a | ~b;
}

DEF: int16x8
vornq_s16 (int16x8_t a, int16x8_t b)
{
  return a | ~b;
}

DEF: int32x4
vornq_s32 (int32x4_t a, int32x4_t b)
{
  return a | ~b;
}

DEF: int64x2
vornq_s64 (int64x2_t a, int64x2_t b)
{
  return a | ~b;
}

DEF: uint8x16
vornq_u8 (uint8x16_t a, uint8x16_t b)
{
  return a | ~b;
}

DEF: uint16x8
vornq_u16 (uint16x8_t a, uint16x8_t b)
{
  return a | ~b;
}

DEF: uint32x4
vornq_u32 (uint32x4_t a, uint32x4_t b)
{
  return a | ~b;
}

DEF: uint64x2
vornq_u64 (uint64x2_t a, uint64x2_t b)
{
  return a | ~b;
}

DEF: int8x8
vsub_s8 (int8x8_t a, int8x8_t b)
{
  return a - b;
}

DEF: int16x4
vsub_s16 (int16x4_t a, int16x4_t b)
{
  return a - b;
}

DEF: int32x2
vsub_s32 (int32x2_t a, int32x2_t b)
{
  return a - b;
}

DEF: float32x2
vsub_f32 (float32x2_t a, float32x2_t b)
{
  return a - b;
}

DEF: float64x1
vsub_f64 (float64x1_t a, float64x1_t b)
{
  return a - b;
}

DEF: uint8x8
vsub_u8 (uint8x8_t a, uint8x8_t b)
{
  return a - b;
}

DEF: uint16x4
vsub_u16 (uint16x4_t a, uint16x4_t b)
{
  return a - b;
}

DEF: uint32x2
vsub_u32 (uint32x2_t a, uint32x2_t b)
{
  return a - b;
}

DEF: int64x1
vsub_s64 (int64x1_t a, int64x1_t b)
{
  return a - b;
}

DEF: uint64x1
vsub_u64 (uint64x1_t a, uint64x1_t b)
{
  return a - b;
}

DEF: int8x16
vsubq_s8 (int8x16_t a, int8x16_t b)
{
  return a - b;
}

DEF: int16x8
vsubq_s16 (int16x8_t a, int16x8_t b)
{
  return a - b;
}

DEF: int32x4
vsubq_s32 (int32x4_t a, int32x4_t b)
{
  return a - b;
}

DEF: int64x2
vsubq_s64 (int64x2_t a, int64x2_t b)
{
  return a - b;
}

DEF: float32x4
vsubq_f32 (float32x4_t a, float32x4_t b)
{
  return a - b;
}

DEF: float64x2
vsubq_f64 (float64x2_t a, float64x2_t b)
{
  return a - b;
}

DEF: uint8x16
vsubq_u8 (uint8x16_t a, uint8x16_t b)
{
  return a - b;
}

DEF: uint16x8
vsubq_u16 (uint16x8_t a, uint16x8_t b)
{
  return a - b;
}

DEF: uint32x4
vsubq_u32 (uint32x4_t a, uint32x4_t b)
{
  return a - b;
}

DEF: uint64x2
vsubq_u64 (uint64x2_t a, uint64x2_t b)
{
  return a - b;
}

DEF: int16x8
vsubl_s8 (int8x8_t a, int8x8_t b)
{
  return (int16x8_t) builtin_aarch64_ssublv8qi (a, b);
}

DEF: int32x4
vsubl_s16 (int16x4_t a, int16x4_t b)
{
  return (int32x4_t) builtin_aarch64_ssublv4hi (a, b);
}

DEF: int64x2
vsubl_s32 (int32x2_t a, int32x2_t b)
{
  return (int64x2_t) builtin_aarch64_ssublv2si (a, b);
}

DEF: uint16x8
vsubl_u8 (uint8x8_t a, uint8x8_t b)
{
  return (uint16x8_t) builtin_aarch64_usublv8qi ((int8x8_t) a,
         (int8x8_t) b);
}

DEF: uint32x4
vsubl_u16 (uint16x4_t a, uint16x4_t b)
{
  return (uint32x4_t) builtin_aarch64_usublv4hi ((int16x4_t) a,
         (int16x4_t) b);
}

DEF: uint64x2
vsubl_u32 (uint32x2_t a, uint32x2_t b)
{
  return (uint64x2_t) builtin_aarch64_usublv2si ((int32x2_t) a,
         (int32x2_t) b);
}

DEF: int16x8
vsubl_high_s8 (int8x16_t a, int8x16_t b)
{
  return (int16x8_t) builtin_aarch64_ssubl2v16qi (a, b);
}

DEF: int32x4
vsubl_high_s16 (int16x8_t a, int16x8_t b)
{
  return (int32x4_t) builtin_aarch64_ssubl2v8hi (a, b);
}

DEF: int64x2
vsubl_high_s32 (int32x4_t a, int32x4_t b)
{
  return (int64x2_t) builtin_aarch64_ssubl2v4si (a, b);
}

DEF: uint16x8
vsubl_high_u8 (uint8x16_t a, uint8x16_t b)
{
  return (uint16x8_t) builtin_aarch64_usubl2v16qi ((int8x16_t) a,
           (int8x16_t) b);
}

DEF: uint32x4
vsubl_high_u16 (uint16x8_t a, uint16x8_t b)
{
  return (uint32x4_t) builtin_aarch64_usubl2v8hi ((int16x8_t) a,
          (int16x8_t) b);
}

DEF: uint64x2
vsubl_high_u32 (uint32x4_t a, uint32x4_t b)
{
  return (uint64x2_t) builtin_aarch64_usubl2v4si ((int32x4_t) a,
          (int32x4_t) b);
}

DEF: int16x8
vsubw_s8 (int16x8_t a, int8x8_t b)
{
  return (int16x8_t) builtin_aarch64_ssubwv8qi (a, b);
}

DEF: int32x4
vsubw_s16 (int32x4_t a, int16x4_t b)
{
  return (int32x4_t) builtin_aarch64_ssubwv4hi (a, b);
}

DEF: int64x2
vsubw_s32 (int64x2_t a, int32x2_t b)
{
  return (int64x2_t) builtin_aarch64_ssubwv2si (a, b);
}

DEF: uint16x8
vsubw_u8 (uint16x8_t a, uint8x8_t b)
{
  return (uint16x8_t) builtin_aarch64_usubwv8qi ((int16x8_t) a,
         (int8x8_t) b);
}

DEF: uint32x4
vsubw_u16 (uint32x4_t a, uint16x4_t b)
{
  return (uint32x4_t) builtin_aarch64_usubwv4hi ((int32x4_t) a,
         (int16x4_t) b);
}

DEF: uint64x2
vsubw_u32 (uint64x2_t a, uint32x2_t b)
{
  return (uint64x2_t) builtin_aarch64_usubwv2si ((int64x2_t) a,
         (int32x2_t) b);
}

DEF: int16x8
vsubw_high_s8 (int16x8_t a, int8x16_t b)
{
  return (int16x8_t) builtin_aarch64_ssubw2v16qi (a, b);
}

DEF: int32x4
vsubw_high_s16 (int32x4_t a, int16x8_t b)
{
  return (int32x4_t) builtin_aarch64_ssubw2v8hi (a, b);
}

DEF: int64x2
vsubw_high_s32 (int64x2_t a, int32x4_t b)
{
  return (int64x2_t) builtin_aarch64_ssubw2v4si (a, b);
}

DEF: uint16x8
vsubw_high_u8 (uint16x8_t a, uint8x16_t b)
{
  return (uint16x8_t) builtin_aarch64_usubw2v16qi ((int16x8_t) a,
           (int8x16_t) b);
}

DEF: uint32x4
vsubw_high_u16 (uint32x4_t a, uint16x8_t b)
{
  return (uint32x4_t) builtin_aarch64_usubw2v8hi ((int32x4_t) a,
          (int16x8_t) b);
}

DEF: uint64x2
vsubw_high_u32 (uint64x2_t a, uint32x4_t b)
{
  return (uint64x2_t) builtin_aarch64_usubw2v4si ((int64x2_t) a,
          (int32x4_t) b);
}

DEF: int8x8
vqadd_s8 (int8x8_t a, int8x8_t b)
{
  return (int8x8_t) builtin_aarch64_sqaddv8qi (a, b);
}

DEF: int16x4
vqadd_s16 (int16x4_t a, int16x4_t b)
{
  return (int16x4_t) builtin_aarch64_sqaddv4hi (a, b);
}

DEF: int32x2
vqadd_s32 (int32x2_t a, int32x2_t b)
{
  return (int32x2_t) builtin_aarch64_sqaddv2si (a, b);
}

DEF: int64x1
vqadd_s64 (int64x1_t a, int64x1_t b)
{
  return (int64x1_t) {builtin_aarch64_sqadddi (a[0], b[0])};
}

DEF: uint8x8
vqadd_u8 (uint8x8_t a, uint8x8_t b)
{
  return builtin_aarch64_uqaddv8qi_uuu (a, b);
}

DEF: int8x8
vhsub_s8 (int8x8_t a, int8x8_t b)
{
  return (int8x8_t)builtin_aarch64_shsubv8qi (a, b);
}

DEF: int16x4
vhsub_s16 (int16x4_t a, int16x4_t b)
{
  return (int16x4_t) builtin_aarch64_shsubv4hi (a, b);
}

DEF: int32x2
vhsub_s32 (int32x2_t a, int32x2_t b)
{
  return (int32x2_t) builtin_aarch64_shsubv2si (a, b);
}

DEF: uint8x8
vhsub_u8 (uint8x8_t a, uint8x8_t b)
{
  return (uint8x8_t) builtin_aarch64_uhsubv8qi ((int8x8_t) a,
        (int8x8_t) b);
}

DEF: uint16x4
vhsub_u16 (uint16x4_t a, uint16x4_t b)
{
  return (uint16x4_t) builtin_aarch64_uhsubv4hi ((int16x4_t) a,
         (int16x4_t) b);
}

DEF: uint32x2
vhsub_u32 (uint32x2_t a, uint32x2_t b)
{
  return (uint32x2_t) builtin_aarch64_uhsubv2si ((int32x2_t) a,
         (int32x2_t) b);
}

DEF: int8x16
vhsubq_s8 (int8x16_t a, int8x16_t b)
{
  return (int8x16_t) builtin_aarch64_shsubv16qi (a, b);
}

DEF: int16x8
vhsubq_s16 (int16x8_t a, int16x8_t b)
{
  return (int16x8_t) builtin_aarch64_shsubv8hi (a, b);
}

DEF: int32x4
vhsubq_s32 (int32x4_t a, int32x4_t b)
{
  return (int32x4_t) builtin_aarch64_shsubv4si (a, b);
}

DEF: uint8x16
vhsubq_u8 (uint8x16_t a, uint8x16_t b)
{
  return (uint8x16_t) builtin_aarch64_uhsubv16qi ((int8x16_t) a,
          (int8x16_t) b);
}

DEF: uint16x8
vhsubq_u16 (uint16x8_t a, uint16x8_t b)
{
  return (uint16x8_t) builtin_aarch64_uhsubv8hi ((int16x8_t) a,
         (int16x8_t) b);
}

DEF: uint32x4
vhsubq_u32 (uint32x4_t a, uint32x4_t b)
{
  return (uint32x4_t) builtin_aarch64_uhsubv4si ((int32x4_t) a,
         (int32x4_t) b);
}

DEF: int8x8
vsubhn_s16 (int16x8_t a, int16x8_t b)
{
  return (int8x8_t) builtin_aarch64_subhnv8hi (a, b);
}

DEF: int16x4
vsubhn_s32 (int32x4_t a, int32x4_t b)
{
  return (int16x4_t) builtin_aarch64_subhnv4si (a, b);
}

DEF: int32x2
vsubhn_s64 (int64x2_t a, int64x2_t b)
{
  return (int32x2_t) builtin_aarch64_subhnv2di (a, b);
}

DEF: uint8x8
vsubhn_u16 (uint16x8_t a, uint16x8_t b)
{
  return (uint8x8_t) builtin_aarch64_subhnv8hi ((int16x8_t) a,
        (int16x8_t) b);
}

DEF: uint16x4
vsubhn_u32 (uint32x4_t a, uint32x4_t b)
{
  return (uint16x4_t) builtin_aarch64_subhnv4si ((int32x4_t) a,
         (int32x4_t) b);
}

DEF: uint32x2
vsubhn_u64 (uint64x2_t a, uint64x2_t b)
{
  return (uint32x2_t) builtin_aarch64_subhnv2di ((int64x2_t) a,
         (int64x2_t) b);
}

DEF: int8x8
vrsubhn_s16 (int16x8_t a, int16x8_t b)
{
  return (int8x8_t) builtin_aarch64_rsubhnv8hi (a, b);
}

DEF: int16x4
vrsubhn_s32 (int32x4_t a, int32x4_t b)
{
  return (int16x4_t) builtin_aarch64_rsubhnv4si (a, b);
}

DEF: int32x2
vrsubhn_s64 (int64x2_t a, int64x2_t b)
{
  return (int32x2_t) builtin_aarch64_rsubhnv2di (a, b);
}

DEF: uint8x8
vrsubhn_u16 (uint16x8_t a, uint16x8_t b)
{
  return (uint8x8_t) builtin_aarch64_rsubhnv8hi ((int16x8_t) a,
         (int16x8_t) b);
}

DEF: uint16x4
vrsubhn_u32 (uint32x4_t a, uint32x4_t b)
{
  return (uint16x4_t) builtin_aarch64_rsubhnv4si ((int32x4_t) a,
          (int32x4_t) b);
}

DEF: uint32x2
vrsubhn_u64 (uint64x2_t a, uint64x2_t b)
{
  return (uint32x2_t) builtin_aarch64_rsubhnv2di ((int64x2_t) a,
          (int64x2_t) b);
}

DEF: int8x16
vrsubhn_high_s16 (int8x8_t a, int16x8_t b, int16x8_t c)
{
  return (int8x16_t) builtin_aarch64_rsubhn2v8hi (a, b, c);
}

DEF: int16x8
vrsubhn_high_s32 (int16x4_t a, int32x4_t b, int32x4_t c)
{
  return (int16x8_t) builtin_aarch64_rsubhn2v4si (a, b, c);
}

DEF: int32x4
vrsubhn_high_s64 (int32x2_t a, int64x2_t b, int64x2_t c)
{
  return (int32x4_t) builtin_aarch64_rsubhn2v2di (a, b, c);
}

DEF: uint8x16
vrsubhn_high_u16 (uint8x8_t a, uint16x8_t b, uint16x8_t c)
{
  return (uint8x16_t) builtin_aarch64_rsubhn2v8hi ((int8x8_t) a,
           (int16x8_t) b,
           (int16x8_t) c);
}

DEF: uint16x8
vrsubhn_high_u32 (uint16x4_t a, uint32x4_t b, uint32x4_t c)
{
  return (uint16x8_t) builtin_aarch64_rsubhn2v4si ((int16x4_t) a,
           (int32x4_t) b,
           (int32x4_t) c);
}

DEF: uint32x4
vrsubhn_high_u64 (uint32x2_t a, uint64x2_t b, uint64x2_t c)
{
  return (uint32x4_t) builtin_aarch64_rsubhn2v2di ((int32x2_t) a,
           (int64x2_t) b,
           (int64x2_t) c);
}

DEF: int8x16
vsubhn_high_s16 (int8x8_t a, int16x8_t b, int16x8_t c)
{
  return (int8x16_t) builtin_aarch64_subhn2v8hi (a, b, c);
}

DEF: int16x8
vsubhn_high_s32 (int16x4_t a, int32x4_t b, int32x4_t c)
{
  return (int16x8_t) builtin_aarch64_subhn2v4si (a, b, c);;
}

DEF: int32x4
vsubhn_high_s64 (int32x2_t a, int64x2_t b, int64x2_t c)
{
  return (int32x4_t) builtin_aarch64_subhn2v2di (a, b, c);
}

DEF: uint8x16
vsubhn_high_u16 (uint8x8_t a, uint16x8_t b, uint16x8_t c)
{
  return (uint8x16_t) builtin_aarch64_subhn2v8hi ((int8x8_t) a,
          (int16x8_t) b,
          (int16x8_t) c);
}

DEF: uint16x8
vsubhn_high_u32 (uint16x4_t a, uint32x4_t b, uint32x4_t c)
{
  return (uint16x8_t) builtin_aarch64_subhn2v4si ((int16x4_t) a,
          (int32x4_t) b,
          (int32x4_t) c);
}

DEF: uint32x4
vsubhn_high_u64 (uint32x2_t a, uint64x2_t b, uint64x2_t c)
{
  return (uint32x4_t) builtin_aarch64_subhn2v2di ((int32x2_t) a,
          (int64x2_t) b,
          (int64x2_t) c);
}

DEF: uint16x4
vqadd_u16 (uint16x4_t a, uint16x4_t b)
{
  return builtin_aarch64_uqaddv4hi_uuu (a, b);
}

DEF: uint32x2
vqadd_u32 (uint32x2_t a, uint32x2_t b)
{
  return builtin_aarch64_uqaddv2si_uuu (a, b);
}

DEF: uint64x1
vqadd_u64 (uint64x1_t a, uint64x1_t b)
{
  return (uint64x1_t) {builtin_aarch64_uqadddi_uuu (a[0], b[0])};
}

DEF: int8x16
vqaddq_s8 (int8x16_t a, int8x16_t b)
{
  return (int8x16_t) builtin_aarch64_sqaddv16qi (a, b);
}

DEF: int16x8
vqaddq_s16 (int16x8_t a, int16x8_t b)
{
  return (int16x8_t) builtin_aarch64_sqaddv8hi (a, b);
}

DEF: int32x4
vqaddq_s32 (int32x4_t a, int32x4_t b)
{
  return (int32x4_t) builtin_aarch64_sqaddv4si (a, b);
}

DEF: int64x2
vqaddq_s64 (int64x2_t a, int64x2_t b)
{
  return (int64x2_t) builtin_aarch64_sqaddv2di (a, b);
}

DEF: uint8x16
vqaddq_u8 (uint8x16_t a, uint8x16_t b)
{
  return builtin_aarch64_uqaddv16qi_uuu (a, b);
}

DEF: uint16x8
vqaddq_u16 (uint16x8_t a, uint16x8_t b)
{
  return builtin_aarch64_uqaddv8hi_uuu (a, b);
}

DEF: uint32x4
vqaddq_u32 (uint32x4_t a, uint32x4_t b)
{
  return builtin_aarch64_uqaddv4si_uuu (a, b);
}

DEF: uint64x2
vqaddq_u64 (uint64x2_t a, uint64x2_t b)
{
  return builtin_aarch64_uqaddv2di_uuu (a, b);
}

DEF: int8x8
vqsub_s8 (int8x8_t a, int8x8_t b)
{
  return (int8x8_t) builtin_aarch64_sqsubv8qi (a, b);
}

DEF: int16x4
vqsub_s16 (int16x4_t a, int16x4_t b)
{
  return (int16x4_t) builtin_aarch64_sqsubv4hi (a, b);
}

DEF: int32x2
vqsub_s32 (int32x2_t a, int32x2_t b)
{
  return (int32x2_t) builtin_aarch64_sqsubv2si (a, b);
}

DEF: int64x1
vqsub_s64 (int64x1_t a, int64x1_t b)
{
  return (int64x1_t) {builtin_aarch64_sqsubdi (a[0], b[0])};
}

DEF: uint8x8
vqsub_u8 (uint8x8_t a, uint8x8_t b)
{
  return builtin_aarch64_uqsubv8qi_uuu (a, b);
}

DEF: uint16x4
vqsub_u16 (uint16x4_t a, uint16x4_t b)
{
  return builtin_aarch64_uqsubv4hi_uuu (a, b);
}

DEF: uint32x2
vqsub_u32 (uint32x2_t a, uint32x2_t b)
{
  return builtin_aarch64_uqsubv2si_uuu (a, b);
}

DEF: uint64x1
vqsub_u64 (uint64x1_t a, uint64x1_t b)
{
  return (uint64x1_t) {builtin_aarch64_uqsubdi_uuu (a[0], b[0])};
}

DEF: int8x16
vqsubq_s8 (int8x16_t a, int8x16_t b)
{
  return (int8x16_t) builtin_aarch64_sqsubv16qi (a, b);
}

DEF: int16x8
vqsubq_s16 (int16x8_t a, int16x8_t b)
{
  return (int16x8_t) builtin_aarch64_sqsubv8hi (a, b);
}

DEF: int32x4
vqsubq_s32 (int32x4_t a, int32x4_t b)
{
  return (int32x4_t) builtin_aarch64_sqsubv4si (a, b);
}

DEF: int64x2
vqsubq_s64 (int64x2_t a, int64x2_t b)
{
  return (int64x2_t) builtin_aarch64_sqsubv2di (a, b);
}

DEF: uint8x16
vqsubq_u8 (uint8x16_t a, uint8x16_t b)
{
  return builtin_aarch64_uqsubv16qi_uuu (a, b);
}

DEF: uint16x8
vqsubq_u16 (uint16x8_t a, uint16x8_t b)
{
  return builtin_aarch64_uqsubv8hi_uuu (a, b);
}

DEF: uint32x4
vqsubq_u32 (uint32x4_t a, uint32x4_t b)
{
  return builtin_aarch64_uqsubv4si_uuu (a, b);
}

DEF: uint64x2
vqsubq_u64 (uint64x2_t a, uint64x2_t b)
{
  return builtin_aarch64_uqsubv2di_uuu (a, b);
}

DEF: int8x8
vqneg_s8 (int8x8_t a)
{
  return (int8x8_t) builtin_aarch64_sqnegv8qi (a);
}

DEF: int16x4
vqneg_s16 (int16x4_t a)
{
  return (int16x4_t) builtin_aarch64_sqnegv4hi (a);
}

DEF: int32x2
vqneg_s32 (int32x2_t a)
{
  return (int32x2_t) builtin_aarch64_sqnegv2si (a);
}

DEF: int64x1
vqneg_s64 (int64x1_t a)
{
  return (int64x1_t) {builtin_aarch64_sqnegdi (a[0])};
}

DEF: int8x16
vqnegq_s8 (int8x16_t a)
{
  return (int8x16_t) builtin_aarch64_sqnegv16qi (a);
}

DEF: int16x8
vqnegq_s16 (int16x8_t a)
{
  return (int16x8_t) builtin_aarch64_sqnegv8hi (a);
}

DEF: int32x4
vqnegq_s32 (int32x4_t a)
{
  return (int32x4_t) builtin_aarch64_sqnegv4si (a);
}

DEF: int8x8
vqabs_s8 (int8x8_t a)
{
  return (int8x8_t) builtin_aarch64_sqabsv8qi (a);
}

DEF: int16x4
vqabs_s16 (int16x4_t a)
{
  return (int16x4_t) builtin_aarch64_sqabsv4hi (a);
}

DEF: int32x2
vqabs_s32 (int32x2_t a)
{
  return (int32x2_t) builtin_aarch64_sqabsv2si (a);
}

DEF: int64x1
vqabs_s64 (int64x1_t a)
{
  return (int64x1_t) {builtin_aarch64_sqabsdi (a[0])};
}

DEF: int8x16
vqabsq_s8 (int8x16_t a)
{
  return (int8x16_t) builtin_aarch64_sqabsv16qi (a);
}

DEF: int16x8
vqabsq_s16 (int16x8_t a)
{
  return (int16x8_t) builtin_aarch64_sqabsv8hi (a);
}

DEF: int32x4
vqabsq_s32 (int32x4_t a)
{
  return (int32x4_t) builtin_aarch64_sqabsv4si (a);
}

DEF: int16x4
vqdmulh_s16 (int16x4_t a, int16x4_t b)
{
  return (int16x4_t) builtin_aarch64_sqdmulhv4hi (a, b);
}

DEF: int32x2
vqdmulh_s32 (int32x2_t a, int32x2_t b)
{
  return (int32x2_t) builtin_aarch64_sqdmulhv2si (a, b);
}

DEF: int16x8
vqdmulhq_s16 (int16x8_t a, int16x8_t b)
{
  return (int16x8_t) builtin_aarch64_sqdmulhv8hi (a, b);
}

DEF: int32x4
vqdmulhq_s32 (int32x4_t a, int32x4_t b)
{
  return (int32x4_t) builtin_aarch64_sqdmulhv4si (a, b);
}

DEF: int16x4
vqrdmulh_s16 (int16x4_t a, int16x4_t b)
{
  return (int16x4_t) builtin_aarch64_sqrdmulhv4hi (a, b);
}

DEF: int32x2
vqrdmulh_s32 (int32x2_t a, int32x2_t b)
{
  return (int32x2_t) builtin_aarch64_sqrdmulhv2si (a, b);
}

DEF: int16x8
vqrdmulhq_s16 (int16x8_t a, int16x8_t b)
{
  return (int16x8_t) builtin_aarch64_sqrdmulhv8hi (a, b);
}

DEF: int32x4
vqrdmulhq_s32 (int32x4_t a, int32x4_t b)
{
  return (int32x4_t) builtin_aarch64_sqrdmulhv4si (a, b);
}

DEF: int8x8
vcreate_s8 (uint64_t a)
{
  return (int8x8_t) a;
}

DEF: int16x4
vcreate_s16 (uint64_t a)
{
  return (int16x4_t) a;
}

DEF: int32x2
vcreate_s32 (uint64_t a)
{
  return (int32x2_t) a;
}

DEF: int64x1
vcreate_s64 (uint64_t a)
{
  return (int64x1_t) {a};
}

DEF: float32x2
vcreate_f32 (uint64_t a)
{
  return (float32x2_t) a;
}

DEF: uint8x8
vcreate_u8 (uint64_t a)
{
  return (uint8x8_t) a;
}

DEF: uint16x4
vcreate_u16 (uint64_t a)
{
  return (uint16x4_t) a;
}

DEF: uint32x2
vcreate_u32 (uint64_t a)
{
  return (uint32x2_t) a;
}

DEF: uint64x1
vcreate_u64 (uint64_t a)
{
  return (uint64x1_t) {a};
}

DEF: float64x1
vcreate_f64 (uint64_t a)
{
  return (float64x1_t) a;
}

DEF: poly8x8
vcreate_p8 (uint64_t a)
{
  return (poly8x8_t) a;
}

DEF: poly16x4
vcreate_p16 (uint64_t a)
{
  return (poly16x4_t) a;
}

DEF: float32
vget_lane_f32 (float32x2_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: float64
vget_lane_f64 (float64x1_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: poly8
vget_lane_p8 (poly8x8_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: poly16
vget_lane_p16 (poly16x4_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: int8
vget_lane_s8 (int8x8_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: int16
vget_lane_s16 (int16x4_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: int32
vget_lane_s32 (int32x2_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: int64
vget_lane_s64 (int64x1_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: uint8
vget_lane_u8 (uint8x8_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: uint16
vget_lane_u16 (uint16x4_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: uint32
vget_lane_u32 (uint32x2_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: uint64
vget_lane_u64 (uint64x1_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: float32
vgetq_lane_f32 (float32x4_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: float64
vgetq_lane_f64 (float64x2_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: poly8
vgetq_lane_p8 (poly8x16_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: poly16
vgetq_lane_p16 (poly16x8_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: int8
vgetq_lane_s8 (int8x16_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: int16
vgetq_lane_s16 (int16x8_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: int32
vgetq_lane_s32 (int32x4_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: int64
vgetq_lane_s64 (int64x2_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: uint8
vgetq_lane_u8 (uint8x16_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: uint16
vgetq_lane_u16 (uint16x8_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: uint32
vgetq_lane_u32 (uint32x4_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: uint64
vgetq_lane_u64 (uint64x2_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: poly8x8
vreinterpret_p8_f64 (float64x1_t a)
{
  return (poly8x8_t) a;
}

DEF: poly8x8
vreinterpret_p8_s8 (int8x8_t a)
{
  return (poly8x8_t) a;
}

DEF: poly8x8
vreinterpret_p8_s16 (int16x4_t a)
{
  return (poly8x8_t) a;
}

DEF: poly8x8
vreinterpret_p8_s32 (int32x2_t a)
{
  return (poly8x8_t) a;
}

DEF: poly8x8
vreinterpret_p8_s64 (int64x1_t a)
{
  return (poly8x8_t) a;
}

DEF: poly8x8
vreinterpret_p8_f32 (float32x2_t a)
{
  return (poly8x8_t) a;
}

DEF: poly8x8
vreinterpret_p8_u8 (uint8x8_t a)
{
  return (poly8x8_t) a;
}

DEF: poly8x8
vreinterpret_p8_u16 (uint16x4_t a)
{
  return (poly8x8_t) a;
}

DEF: poly8x8
vreinterpret_p8_u32 (uint32x2_t a)
{
  return (poly8x8_t) a;
}

DEF: poly8x8
vreinterpret_p8_u64 (uint64x1_t a)
{
  return (poly8x8_t) a;
}

DEF: poly8x8
vreinterpret_p8_p16 (poly16x4_t a)
{
  return (poly8x8_t) a;
}

DEF: poly8x16
vreinterpretq_p8_f64 (float64x2_t a)
{
  return (poly8x16_t) a;
}

DEF: poly8x16
vreinterpretq_p8_s8 (int8x16_t a)
{
  return (poly8x16_t) a;
}

DEF: poly8x16
vreinterpretq_p8_s16 (int16x8_t a)
{
  return (poly8x16_t) a;
}

DEF: poly8x16
vreinterpretq_p8_s32 (int32x4_t a)
{
  return (poly8x16_t) a;
}

DEF: poly8x16
vreinterpretq_p8_s64 (int64x2_t a)
{
  return (poly8x16_t) a;
}

DEF: poly8x16
vreinterpretq_p8_f32 (float32x4_t a)
{
  return (poly8x16_t) a;
}

DEF: poly8x16
vreinterpretq_p8_u8 (uint8x16_t a)
{
  return (poly8x16_t) a;
}

DEF: poly8x16
vreinterpretq_p8_u16 (uint16x8_t a)
{
  return (poly8x16_t) a;
}

DEF: poly8x16
vreinterpretq_p8_u32 (uint32x4_t a)
{
  return (poly8x16_t) a;
}

DEF: poly8x16
vreinterpretq_p8_u64 (uint64x2_t a)
{
  return (poly8x16_t) a;
}

DEF: poly8x16
vreinterpretq_p8_p16 (poly16x8_t a)
{
  return (poly8x16_t) a;
}

DEF: poly16x4
vreinterpret_p16_f64 (float64x1_t a)
{
  return (poly16x4_t) a;
}

DEF: poly16x4
vreinterpret_p16_s8 (int8x8_t a)
{
  return (poly16x4_t) a;
}

DEF: poly16x4
vreinterpret_p16_s16 (int16x4_t a)
{
  return (poly16x4_t) a;
}

DEF: poly16x4
vreinterpret_p16_s32 (int32x2_t a)
{
  return (poly16x4_t) a;
}

DEF: poly16x4
vreinterpret_p16_s64 (int64x1_t a)
{
  return (poly16x4_t) a;
}

DEF: poly16x4
vreinterpret_p16_f32 (float32x2_t a)
{
  return (poly16x4_t) a;
}

DEF: poly16x4
vreinterpret_p16_u8 (uint8x8_t a)
{
  return (poly16x4_t) a;
}

DEF: poly16x4
vreinterpret_p16_u16 (uint16x4_t a)
{
  return (poly16x4_t) a;
}

DEF: poly16x4
vreinterpret_p16_u32 (uint32x2_t a)
{
  return (poly16x4_t) a;
}

DEF: poly16x4
vreinterpret_p16_u64 (uint64x1_t a)
{
  return (poly16x4_t) a;
}

DEF: poly16x4
vreinterpret_p16_p8 (poly8x8_t a)
{
  return (poly16x4_t) a;
}

DEF: poly16x8
vreinterpretq_p16_f64 (float64x2_t a)
{
  return (poly16x8_t) a;
}

DEF: poly16x8
vreinterpretq_p16_s8 (int8x16_t a)
{
  return (poly16x8_t) a;
}

DEF: poly16x8
vreinterpretq_p16_s16 (int16x8_t a)
{
  return (poly16x8_t) a;
}

DEF: poly16x8
vreinterpretq_p16_s32 (int32x4_t a)
{
  return (poly16x8_t) a;
}

DEF: poly16x8
vreinterpretq_p16_s64 (int64x2_t a)
{
  return (poly16x8_t) a;
}

DEF: poly16x8
vreinterpretq_p16_f32 (float32x4_t a)
{
  return (poly16x8_t) a;
}

DEF: poly16x8
vreinterpretq_p16_u8 (uint8x16_t a)
{
  return (poly16x8_t) a;
}

DEF: poly16x8
vreinterpretq_p16_u16 (uint16x8_t a)
{
  return (poly16x8_t) a;
}

DEF: poly16x8
vreinterpretq_p16_u32 (uint32x4_t a)
{
  return (poly16x8_t) a;
}

DEF: poly16x8
vreinterpretq_p16_u64 (uint64x2_t a)
{
  return (poly16x8_t) a;
}

DEF: poly16x8
vreinterpretq_p16_p8 (poly8x16_t a)
{
  return (poly16x8_t) a;
}

DEF: float32x2
vreinterpret_f32_f64 (float64x1_t a)
{
  return (float32x2_t) a;
}

DEF: float32x2
vreinterpret_f32_s8 (int8x8_t a)
{
  return (float32x2_t) a;
}

DEF: float32x2
vreinterpret_f32_s16 (int16x4_t a)
{
  return (float32x2_t) a;
}

DEF: float32x2
vreinterpret_f32_s32 (int32x2_t a)
{
  return (float32x2_t) a;
}

DEF: float32x2
vreinterpret_f32_s64 (int64x1_t a)
{
  return (float32x2_t) a;
}

DEF: float32x2
vreinterpret_f32_u8 (uint8x8_t a)
{
  return (float32x2_t) a;
}

DEF: float32x2
vreinterpret_f32_u16 (uint16x4_t a)
{
  return (float32x2_t) a;
}

DEF: float32x2
vreinterpret_f32_u32 (uint32x2_t a)
{
  return (float32x2_t) a;
}

DEF: float32x2
vreinterpret_f32_u64 (uint64x1_t a)
{
  return (float32x2_t) a;
}

DEF: float32x2
vreinterpret_f32_p8 (poly8x8_t a)
{
  return (float32x2_t) a;
}

DEF: float32x2
vreinterpret_f32_p16 (poly16x4_t a)
{
  return (float32x2_t) a;
}

DEF: float32x4
vreinterpretq_f32_f64 (float64x2_t a)
{
  return (float32x4_t) a;
}

DEF: float32x4
vreinterpretq_f32_s8 (int8x16_t a)
{
  return (float32x4_t) a;
}

DEF: float32x4
vreinterpretq_f32_s16 (int16x8_t a)
{
  return (float32x4_t) a;
}

DEF: float32x4
vreinterpretq_f32_s32 (int32x4_t a)
{
  return (float32x4_t) a;
}

DEF: float32x4
vreinterpretq_f32_s64 (int64x2_t a)
{
  return (float32x4_t) a;
}

DEF: float32x4
vreinterpretq_f32_u8 (uint8x16_t a)
{
  return (float32x4_t) a;
}

DEF: float32x4
vreinterpretq_f32_u16 (uint16x8_t a)
{
  return (float32x4_t) a;
}

DEF: float32x4
vreinterpretq_f32_u32 (uint32x4_t a)
{
  return (float32x4_t) a;
}

DEF: float32x4
vreinterpretq_f32_u64 (uint64x2_t a)
{
  return (float32x4_t) a;
}

DEF: float32x4
vreinterpretq_f32_p8 (poly8x16_t a)
{
  return (float32x4_t) a;
}

DEF: float32x4
vreinterpretq_f32_p16 (poly16x8_t a)
{
  return (float32x4_t) a;
}

DEF: float64x1_t 
vreinterpret_f64_f32 (float32x2_t a)
{
  return (float64x1_t) a;
}

DEF: float64x1_t 
vreinterpret_f64_p8 (poly8x8_t a)
{
  return (float64x1_t) a;
}

DEF: float64x1_t 
vreinterpret_f64_p16 (poly16x4_t a)
{
  return (float64x1_t) a;
}

DEF: float64x1_t 
vreinterpret_f64_s8 (int8x8_t a)
{
  return (float64x1_t) a;
}

DEF: float64x1_t 
vreinterpret_f64_s16 (int16x4_t a)
{
  return (float64x1_t) a;
}

DEF: float64x1_t 
vreinterpret_f64_s32 (int32x2_t a)
{
  return (float64x1_t) a;
}

DEF: float64x1_t 
vreinterpret_f64_s64 (int64x1_t a)
{
  return (float64x1_t) a;
}

DEF: float64x1_t 
vreinterpret_f64_u8 (uint8x8_t a)
{
  return (float64x1_t) a;
}

DEF: float64x1_t 
vreinterpret_f64_u16 (uint16x4_t a)
{
  return (float64x1_t) a;
}

DEF: float64x1_t 
vreinterpret_f64_u32 (uint32x2_t a)
{
  return (float64x1_t) a;
}

DEF: float64x1_t 
vreinterpret_f64_u64 (uint64x1_t a)
{
  return (float64x1_t) a;
}

DEF: float64x2_t 
vreinterpretq_f64_f32 (float32x4_t a)
{
  return (float64x2_t) a;
}

DEF: float64x2_t 
vreinterpretq_f64_p8 (poly8x16_t a)
{
  return (float64x2_t) a;
}

DEF: float64x2_t 
vreinterpretq_f64_p16 (poly16x8_t a)
{
  return (float64x2_t) a;
}

DEF: float64x2_t 
vreinterpretq_f64_s8 (int8x16_t a)
{
  return (float64x2_t) a;
}

DEF: float64x2_t 
vreinterpretq_f64_s16 (int16x8_t a)
{
  return (float64x2_t) a;
}

DEF: float64x2_t 
vreinterpretq_f64_s32 (int32x4_t a)
{
  return (float64x2_t) a;
}

DEF: float64x2_t 
vreinterpretq_f64_s64 (int64x2_t a)
{
  return (float64x2_t) a;
}

DEF: float64x2_t 
vreinterpretq_f64_u8 (uint8x16_t a)
{
  return (float64x2_t) a;
}

DEF: float64x2_t 
vreinterpretq_f64_u16 (uint16x8_t a)
{
  return (float64x2_t) a;
}

DEF: float64x2_t 
vreinterpretq_f64_u32 (uint32x4_t a)
{
  return (float64x2_t) a;
}

DEF: float64x2_t 
vreinterpretq_f64_u64 (uint64x2_t a)
{
  return (float64x2_t) a;
}

DEF: int64x1
vreinterpret_s64_f64 (float64x1_t a)
{
  return (int64x1_t) a;
}

DEF: int64x1
vreinterpret_s64_s8 (int8x8_t a)
{
  return (int64x1_t) a;
}

DEF: int64x1
vreinterpret_s64_s16 (int16x4_t a)
{
  return (int64x1_t) a;
}

DEF: int64x1
vreinterpret_s64_s32 (int32x2_t a)
{
  return (int64x1_t) a;
}

DEF: int64x1
vreinterpret_s64_f32 (float32x2_t a)
{
  return (int64x1_t) a;
}

DEF: int64x1
vreinterpret_s64_u8 (uint8x8_t a)
{
  return (int64x1_t) a;
}

DEF: int64x1
vreinterpret_s64_u16 (uint16x4_t a)
{
  return (int64x1_t) a;
}

DEF: int64x1
vreinterpret_s64_u32 (uint32x2_t a)
{
  return (int64x1_t) a;
}

DEF: int64x1
vreinterpret_s64_u64 (uint64x1_t a)
{
  return (int64x1_t) a;
}

DEF: int64x1
vreinterpret_s64_p8 (poly8x8_t a)
{
  return (int64x1_t) a;
}

DEF: int64x1
vreinterpret_s64_p16 (poly16x4_t a)
{
  return (int64x1_t) a;
}

DEF: int64x2
vreinterpretq_s64_f64 (float64x2_t a)
{
  return (int64x2_t) a;
}

DEF: int64x2
vreinterpretq_s64_s8 (int8x16_t a)
{
  return (int64x2_t) a;
}

DEF: int64x2
vreinterpretq_s64_s16 (int16x8_t a)
{
  return (int64x2_t) a;
}

DEF: int64x2
vreinterpretq_s64_s32 (int32x4_t a)
{
  return (int64x2_t) a;
}

DEF: int64x2
vreinterpretq_s64_f32 (float32x4_t a)
{
  return (int64x2_t) a;
}

DEF: int64x2
vreinterpretq_s64_u8 (uint8x16_t a)
{
  return (int64x2_t) a;
}

DEF: int64x2
vreinterpretq_s64_u16 (uint16x8_t a)
{
  return (int64x2_t) a;
}

DEF: int64x2
vreinterpretq_s64_u32 (uint32x4_t a)
{
  return (int64x2_t) a;
}

DEF: int64x2
vreinterpretq_s64_u64 (uint64x2_t a)
{
  return (int64x2_t) a;
}

DEF: int64x2
vreinterpretq_s64_p8 (poly8x16_t a)
{
  return (int64x2_t) a;
}

DEF: int64x2
vreinterpretq_s64_p16 (poly16x8_t a)
{
  return (int64x2_t) a;
}

DEF: uint64x1
vreinterpret_u64_f64 (float64x1_t a)
{
  return (uint64x1_t) a;
}

DEF: uint64x1
vreinterpret_u64_s8 (int8x8_t a)
{
  return (uint64x1_t) a;
}

DEF: uint64x1
vreinterpret_u64_s16 (int16x4_t a)
{
  return (uint64x1_t) a;
}

DEF: uint64x1
vreinterpret_u64_s32 (int32x2_t a)
{
  return (uint64x1_t) a;
}

DEF: uint64x1
vreinterpret_u64_s64 (int64x1_t a)
{
  return (uint64x1_t) a;
}

DEF: uint64x1
vreinterpret_u64_f32 (float32x2_t a)
{
  return (uint64x1_t) a;
}

DEF: uint64x1
vreinterpret_u64_u8 (uint8x8_t a)
{
  return (uint64x1_t) a;
}

DEF: uint64x1
vreinterpret_u64_u16 (uint16x4_t a)
{
  return (uint64x1_t) a;
}

DEF: uint64x1
vreinterpret_u64_u32 (uint32x2_t a)
{
  return (uint64x1_t) a;
}

DEF: uint64x1
vreinterpret_u64_p8 (poly8x8_t a)
{
  return (uint64x1_t) a;
}

DEF: uint64x1
vreinterpret_u64_p16 (poly16x4_t a)
{
  return (uint64x1_t) a;
}

DEF: uint64x2
vreinterpretq_u64_f64 (float64x2_t a)
{
  return (uint64x2_t) a;
}

DEF: uint64x2
vreinterpretq_u64_s8 (int8x16_t a)
{
  return (uint64x2_t) a;
}

DEF: uint64x2
vreinterpretq_u64_s16 (int16x8_t a)
{
  return (uint64x2_t) a;
}

DEF: uint64x2
vreinterpretq_u64_s32 (int32x4_t a)
{
  return (uint64x2_t) a;
}

DEF: uint64x2
vreinterpretq_u64_s64 (int64x2_t a)
{
  return (uint64x2_t) a;
}

DEF: uint64x2
vreinterpretq_u64_f32 (float32x4_t a)
{
  return (uint64x2_t) a;
}

DEF: uint64x2
vreinterpretq_u64_u8 (uint8x16_t a)
{
  return (uint64x2_t) a;
}

DEF: uint64x2
vreinterpretq_u64_u16 (uint16x8_t a)
{
  return (uint64x2_t) a;
}

DEF: uint64x2
vreinterpretq_u64_u32 (uint32x4_t a)
{
  return (uint64x2_t) a;
}

DEF: uint64x2
vreinterpretq_u64_p8 (poly8x16_t a)
{
  return (uint64x2_t) a;
}

DEF: uint64x2
vreinterpretq_u64_p16 (poly16x8_t a)
{
  return (uint64x2_t) a;
}

DEF: int8x8
vreinterpret_s8_f64 (float64x1_t a)
{
  return (int8x8_t) a;
}

DEF: int8x8
vreinterpret_s8_s16 (int16x4_t a)
{
  return (int8x8_t) a;
}

DEF: int8x8
vreinterpret_s8_s32 (int32x2_t a)
{
  return (int8x8_t) a;
}

DEF: int8x8
vreinterpret_s8_s64 (int64x1_t a)
{
  return (int8x8_t) a;
}

DEF: int8x8
vreinterpret_s8_f32 (float32x2_t a)
{
  return (int8x8_t) a;
}

DEF: int8x8
vreinterpret_s8_u8 (uint8x8_t a)
{
  return (int8x8_t) a;
}

DEF: int8x8
vreinterpret_s8_u16 (uint16x4_t a)
{
  return (int8x8_t) a;
}

DEF: int8x8
vreinterpret_s8_u32 (uint32x2_t a)
{
  return (int8x8_t) a;
}

DEF: int8x8
vreinterpret_s8_u64 (uint64x1_t a)
{
  return (int8x8_t) a;
}

DEF: int8x8
vreinterpret_s8_p8 (poly8x8_t a)
{
  return (int8x8_t) a;
}

DEF: int8x8
vreinterpret_s8_p16 (poly16x4_t a)
{
  return (int8x8_t) a;
}

DEF: int8x16
vreinterpretq_s8_f64 (float64x2_t a)
{
  return (int8x16_t) a;
}

DEF: int8x16
vreinterpretq_s8_s16 (int16x8_t a)
{
  return (int8x16_t) a;
}

DEF: int8x16
vreinterpretq_s8_s32 (int32x4_t a)
{
  return (int8x16_t) a;
}

DEF: int8x16
vreinterpretq_s8_s64 (int64x2_t a)
{
  return (int8x16_t) a;
}

DEF: int8x16
vreinterpretq_s8_f32 (float32x4_t a)
{
  return (int8x16_t) a;
}

DEF: int8x16
vreinterpretq_s8_u8 (uint8x16_t a)
{
  return (int8x16_t) a;
}

DEF: int8x16
vreinterpretq_s8_u16 (uint16x8_t a)
{
  return (int8x16_t) a;
}

DEF: int8x16
vreinterpretq_s8_u32 (uint32x4_t a)
{
  return (int8x16_t) a;
}

DEF: int8x16
vreinterpretq_s8_u64 (uint64x2_t a)
{
  return (int8x16_t) a;
}

DEF: int8x16
vreinterpretq_s8_p8 (poly8x16_t a)
{
  return (int8x16_t) a;
}

DEF: int8x16
vreinterpretq_s8_p16 (poly16x8_t a)
{
  return (int8x16_t) a;
}

DEF: int16x4
vreinterpret_s16_f64 (float64x1_t a)
{
  return (int16x4_t) a;
}

DEF: int16x4
vreinterpret_s16_s8 (int8x8_t a)
{
  return (int16x4_t) a;
}

DEF: int16x4
vreinterpret_s16_s32 (int32x2_t a)
{
  return (int16x4_t) a;
}

DEF: int16x4
vreinterpret_s16_s64 (int64x1_t a)
{
  return (int16x4_t) a;
}

DEF: int16x4
vreinterpret_s16_f32 (float32x2_t a)
{
  return (int16x4_t) a;
}

DEF: int16x4
vreinterpret_s16_u8 (uint8x8_t a)
{
  return (int16x4_t) a;
}

DEF: int16x4
vreinterpret_s16_u16 (uint16x4_t a)
{
  return (int16x4_t) a;
}

DEF: int16x4
vreinterpret_s16_u32 (uint32x2_t a)
{
  return (int16x4_t) a;
}

DEF: int16x4
vreinterpret_s16_u64 (uint64x1_t a)
{
  return (int16x4_t) a;
}

DEF: int16x4
vreinterpret_s16_p8 (poly8x8_t a)
{
  return (int16x4_t) a;
}

DEF: int16x4
vreinterpret_s16_p16 (poly16x4_t a)
{
  return (int16x4_t) a;
}

DEF: int16x8
vreinterpretq_s16_f64 (float64x2_t a)
{
  return (int16x8_t) a;
}

DEF: int16x8
vreinterpretq_s16_s8 (int8x16_t a)
{
  return (int16x8_t) a;
}

DEF: int16x8
vreinterpretq_s16_s32 (int32x4_t a)
{
  return (int16x8_t) a;
}

DEF: int16x8
vreinterpretq_s16_s64 (int64x2_t a)
{
  return (int16x8_t) a;
}

DEF: int16x8
vreinterpretq_s16_f32 (float32x4_t a)
{
  return (int16x8_t) a;
}

DEF: int16x8
vreinterpretq_s16_u8 (uint8x16_t a)
{
  return (int16x8_t) a;
}

DEF: int16x8
vreinterpretq_s16_u16 (uint16x8_t a)
{
  return (int16x8_t) a;
}

DEF: int16x8
vreinterpretq_s16_u32 (uint32x4_t a)
{
  return (int16x8_t) a;
}

DEF: int16x8
vreinterpretq_s16_u64 (uint64x2_t a)
{
  return (int16x8_t) a;
}

DEF: int16x8
vreinterpretq_s16_p8 (poly8x16_t a)
{
  return (int16x8_t) a;
}

DEF: int16x8
vreinterpretq_s16_p16 (poly16x8_t a)
{
  return (int16x8_t) a;
}

DEF: int32x2
vreinterpret_s32_f64 (float64x1_t a)
{
  return (int32x2_t) a;
}

DEF: int32x2
vreinterpret_s32_s8 (int8x8_t a)
{
  return (int32x2_t) a;
}

DEF: int32x2
vreinterpret_s32_s16 (int16x4_t a)
{
  return (int32x2_t) a;
}

DEF: int32x2
vreinterpret_s32_s64 (int64x1_t a)
{
  return (int32x2_t) a;
}

DEF: int32x2
vreinterpret_s32_f32 (float32x2_t a)
{
  return (int32x2_t) a;
}

DEF: int32x2
vreinterpret_s32_u8 (uint8x8_t a)
{
  return (int32x2_t) a;
}

DEF: int32x2
vreinterpret_s32_u16 (uint16x4_t a)
{
  return (int32x2_t) a;
}

DEF: int32x2
vreinterpret_s32_u32 (uint32x2_t a)
{
  return (int32x2_t) a;
}

DEF: int32x2
vreinterpret_s32_u64 (uint64x1_t a)
{
  return (int32x2_t) a;
}

DEF: int32x2
vreinterpret_s32_p8 (poly8x8_t a)
{
  return (int32x2_t) a;
}

DEF: int32x2
vreinterpret_s32_p16 (poly16x4_t a)
{
  return (int32x2_t) a;
}

DEF: int32x4
vreinterpretq_s32_f64 (float64x2_t a)
{
  return (int32x4_t) a;
}

DEF: int32x4
vreinterpretq_s32_s8 (int8x16_t a)
{
  return (int32x4_t) a;
}

DEF: int32x4
vreinterpretq_s32_s16 (int16x8_t a)
{
  return (int32x4_t) a;
}

DEF: int32x4
vreinterpretq_s32_s64 (int64x2_t a)
{
  return (int32x4_t) a;
}

DEF: int32x4
vreinterpretq_s32_f32 (float32x4_t a)
{
  return (int32x4_t) a;
}

DEF: int32x4
vreinterpretq_s32_u8 (uint8x16_t a)
{
  return (int32x4_t) a;
}

DEF: int32x4
vreinterpretq_s32_u16 (uint16x8_t a)
{
  return (int32x4_t) a;
}

DEF: int32x4
vreinterpretq_s32_u32 (uint32x4_t a)
{
  return (int32x4_t) a;
}

DEF: int32x4
vreinterpretq_s32_u64 (uint64x2_t a)
{
  return (int32x4_t) a;
}

DEF: int32x4
vreinterpretq_s32_p8 (poly8x16_t a)
{
  return (int32x4_t) a;
}

DEF: int32x4
vreinterpretq_s32_p16 (poly16x8_t a)
{
  return (int32x4_t) a;
}

DEF: uint8x8
vreinterpret_u8_f64 (float64x1_t a)
{
  return (uint8x8_t) a;
}

DEF: uint8x8
vreinterpret_u8_s8 (int8x8_t a)
{
  return (uint8x8_t) a;
}

DEF: uint8x8
vreinterpret_u8_s16 (int16x4_t a)
{
  return (uint8x8_t) a;
}

DEF: uint8x8
vreinterpret_u8_s32 (int32x2_t a)
{
  return (uint8x8_t) a;
}

DEF: uint8x8
vreinterpret_u8_s64 (int64x1_t a)
{
  return (uint8x8_t) a;
}

DEF: uint8x8
vreinterpret_u8_f32 (float32x2_t a)
{
  return (uint8x8_t) a;
}

DEF: uint8x8
vreinterpret_u8_u16 (uint16x4_t a)
{
  return (uint8x8_t) a;
}

DEF: uint8x8
vreinterpret_u8_u32 (uint32x2_t a)
{
  return (uint8x8_t) a;
}

DEF: uint8x8
vreinterpret_u8_u64 (uint64x1_t a)
{
  return (uint8x8_t) a;
}

DEF: uint8x8
vreinterpret_u8_p8 (poly8x8_t a)
{
  return (uint8x8_t) a;
}

DEF: uint8x8
vreinterpret_u8_p16 (poly16x4_t a)
{
  return (uint8x8_t) a;
}

DEF: uint8x16
vreinterpretq_u8_f64 (float64x2_t a)
{
  return (uint8x16_t) a;
}

DEF: uint8x16
vreinterpretq_u8_s8 (int8x16_t a)
{
  return (uint8x16_t) a;
}

DEF: uint8x16
vreinterpretq_u8_s16 (int16x8_t a)
{
  return (uint8x16_t) a;
}

DEF: uint8x16
vreinterpretq_u8_s32 (int32x4_t a)
{
  return (uint8x16_t) a;
}

DEF: uint8x16
vreinterpretq_u8_s64 (int64x2_t a)
{
  return (uint8x16_t) a;
}

DEF: uint8x16
vreinterpretq_u8_f32 (float32x4_t a)
{
  return (uint8x16_t) a;
}

DEF: uint8x16
vreinterpretq_u8_u16 (uint16x8_t a)
{
  return (uint8x16_t) a;
}

DEF: uint8x16
vreinterpretq_u8_u32 (uint32x4_t a)
{
  return (uint8x16_t) a;
}

DEF: uint8x16
vreinterpretq_u8_u64 (uint64x2_t a)
{
  return (uint8x16_t) a;
}

DEF: uint8x16
vreinterpretq_u8_p8 (poly8x16_t a)
{
  return (uint8x16_t) a;
}

DEF: uint8x16
vreinterpretq_u8_p16 (poly16x8_t a)
{
  return (uint8x16_t) a;
}

DEF: uint16x4
vreinterpret_u16_f64 (float64x1_t a)
{
  return (uint16x4_t) a;
}

DEF: uint16x4
vreinterpret_u16_s8 (int8x8_t a)
{
  return (uint16x4_t) a;
}

DEF: uint16x4
vreinterpret_u16_s16 (int16x4_t a)
{
  return (uint16x4_t) a;
}

DEF: uint16x4
vreinterpret_u16_s32 (int32x2_t a)
{
  return (uint16x4_t) a;
}

DEF: uint16x4
vreinterpret_u16_s64 (int64x1_t a)
{
  return (uint16x4_t) a;
}

DEF: uint16x4
vreinterpret_u16_f32 (float32x2_t a)
{
  return (uint16x4_t) a;
}

DEF: uint16x4
vreinterpret_u16_u8 (uint8x8_t a)
{
  return (uint16x4_t) a;
}

DEF: uint16x4
vreinterpret_u16_u32 (uint32x2_t a)
{
  return (uint16x4_t) a;
}

DEF: uint16x4
vreinterpret_u16_u64 (uint64x1_t a)
{
  return (uint16x4_t) a;
}

DEF: uint16x4
vreinterpret_u16_p8 (poly8x8_t a)
{
  return (uint16x4_t) a;
}

DEF: uint16x4
vreinterpret_u16_p16 (poly16x4_t a)
{
  return (uint16x4_t) a;
}

DEF: uint16x8
vreinterpretq_u16_f64 (float64x2_t a)
{
  return (uint16x8_t) a;
}

DEF: uint16x8
vreinterpretq_u16_s8 (int8x16_t a)
{
  return (uint16x8_t) a;
}

DEF: uint16x8
vreinterpretq_u16_s16 (int16x8_t a)
{
  return (uint16x8_t) a;
}

DEF: uint16x8
vreinterpretq_u16_s32 (int32x4_t a)
{
  return (uint16x8_t) a;
}

DEF: uint16x8
vreinterpretq_u16_s64 (int64x2_t a)
{
  return (uint16x8_t) a;
}

DEF: uint16x8
vreinterpretq_u16_f32 (float32x4_t a)
{
  return (uint16x8_t) a;
}

DEF: uint16x8
vreinterpretq_u16_u8 (uint8x16_t a)
{
  return (uint16x8_t) a;
}

DEF: uint16x8
vreinterpretq_u16_u32 (uint32x4_t a)
{
  return (uint16x8_t) a;
}

DEF: uint16x8
vreinterpretq_u16_u64 (uint64x2_t a)
{
  return (uint16x8_t) a;
}

DEF: uint16x8
vreinterpretq_u16_p8 (poly8x16_t a)
{
  return (uint16x8_t) a;
}

DEF: uint16x8
vreinterpretq_u16_p16 (poly16x8_t a)
{
  return (uint16x8_t) a;
}

DEF: uint32x2
vreinterpret_u32_f64 (float64x1_t a)
{
  return (uint32x2_t) a;
}

DEF: uint32x2
vreinterpret_u32_s8 (int8x8_t a)
{
  return (uint32x2_t) a;
}

DEF: uint32x2
vreinterpret_u32_s16 (int16x4_t a)
{
  return (uint32x2_t) a;
}

DEF: uint32x2
vreinterpret_u32_s32 (int32x2_t a)
{
  return (uint32x2_t) a;
}

DEF: uint32x2
vreinterpret_u32_s64 (int64x1_t a)
{
  return (uint32x2_t) a;
}

DEF: uint32x2
vreinterpret_u32_f32 (float32x2_t a)
{
  return (uint32x2_t) a;
}

DEF: uint32x2
vreinterpret_u32_u8 (uint8x8_t a)
{
  return (uint32x2_t) a;
}

DEF: uint32x2
vreinterpret_u32_u16 (uint16x4_t a)
{
  return (uint32x2_t) a;
}

DEF: uint32x2
vreinterpret_u32_u64 (uint64x1_t a)
{
  return (uint32x2_t) a;
}

DEF: uint32x2
vreinterpret_u32_p8 (poly8x8_t a)
{
  return (uint32x2_t) a;
}

DEF: uint32x2
vreinterpret_u32_p16 (poly16x4_t a)
{
  return (uint32x2_t) a;
}

DEF: uint32x4
vreinterpretq_u32_f64 (float64x2_t a)
{
  return (uint32x4_t) a;
}

DEF: uint32x4
vreinterpretq_u32_s8 (int8x16_t a)
{
  return (uint32x4_t) a;
}

DEF: uint32x4
vreinterpretq_u32_s16 (int16x8_t a)
{
  return (uint32x4_t) a;
}

DEF: uint32x4
vreinterpretq_u32_s32 (int32x4_t a)
{
  return (uint32x4_t) a;
}

DEF: uint32x4
vreinterpretq_u32_s64 (int64x2_t a)
{
  return (uint32x4_t) a;
}

DEF: uint32x4
vreinterpretq_u32_f32 (float32x4_t a)
{
  return (uint32x4_t) a;
}

DEF: uint32x4
vreinterpretq_u32_u8 (uint8x16_t a)
{
  return (uint32x4_t) a;
}

DEF: uint32x4
vreinterpretq_u32_u16 (uint16x8_t a)
{
  return (uint32x4_t) a;
}

DEF: uint32x4
vreinterpretq_u32_u64 (uint64x2_t a)
{
  return (uint32x4_t) a;
}

DEF: uint32x4
vreinterpretq_u32_p8 (poly8x16_t a)
{
  return (uint32x4_t) a;
}

DEF: uint32x4
vreinterpretq_u32_p16 (poly16x8_t a)
{
  return (uint32x4_t) a;
}

DEF: float32x2
vset_lane_f32 (float32_t elem, float32x2_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: float64x1
vset_lane_f64 (float64_t elem, float64x1_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: poly8x8
vset_lane_p8 (poly8_t elem, poly8x8_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: poly16x4
vset_lane_p16 (poly16_t elem, poly16x4_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: int8x8
vset_lane_s8 (int8_t elem, int8x8_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: int16x4
vset_lane_s16 (int16_t elem, int16x4_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: int32x2
vset_lane_s32 (int32_t elem, int32x2_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: int64x1
vset_lane_s64 (int64_t elem, int64x1_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: uint8x8
vset_lane_u8 (uint8_t elem, uint8x8_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: uint16x4
vset_lane_u16 (uint16_t elem, uint16x4_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: uint32x2
vset_lane_u32 (uint32_t elem, uint32x2_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: uint64x1
vset_lane_u64 (uint64_t elem, uint64x1_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: float32x4
vsetq_lane_f32 (float32_t elem, float32x4_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: float64x2
vsetq_lane_f64 (float64_t elem, float64x2_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: poly8x16
vsetq_lane_p8 (poly8_t elem, poly8x16_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: poly16x8
vsetq_lane_p16 (poly16_t elem, poly16x8_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: int8x16
vsetq_lane_s8 (int8_t elem, int8x16_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: int16x8
vsetq_lane_s16 (int16_t elem, int16x8_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: int32x4
vsetq_lane_s32 (int32_t elem, int32x4_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: int64x2
vsetq_lane_s64 (int64_t elem, int64x2_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: uint8x16
vsetq_lane_u8 (uint8_t elem, uint8x16_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: uint16x8
vsetq_lane_u16 (uint16_t elem, uint16x8_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: uint32x4
vsetq_lane_u32 (uint32_t elem, uint32x4_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}

DEF: uint64x2
vsetq_lane_u64 (uint64_t elem, uint64x2_t vec, const int index)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), index); vec[index] = elem; vec; });
}


DEF: float32x2
vget_low_f32 (float32x4_t a)
{
  uint64x2_t tmp = vreinterpretq_u64_f32 (a); uint64x1_t lo = vcreate_u64 (vgetq_lane_u64 (tmp, 0)); return vreinterpret_f32_u64 (lo);;
}

DEF: float64x1
vget_low_f64 (float64x2_t a)
{
  return (float64x1_t) {vgetq_lane_f64 (a, 0)};
}

DEF: poly8x8
vget_low_p8 (poly8x16_t a)
{
  uint64x2_t tmp = vreinterpretq_u64_p8 (a); uint64x1_t lo = vcreate_u64 (vgetq_lane_u64 (tmp, 0)); return vreinterpret_p8_u64 (lo);;
}

DEF: poly16x4
vget_low_p16 (poly16x8_t a)
{
  uint64x2_t tmp = vreinterpretq_u64_p16 (a); uint64x1_t lo = vcreate_u64 (vgetq_lane_u64 (tmp, 0)); return vreinterpret_p16_u64 (lo);;
}

DEF: int8x8
vget_low_s8 (int8x16_t a)
{
  uint64x2_t tmp = vreinterpretq_u64_s8 (a); uint64x1_t lo = vcreate_u64 (vgetq_lane_u64 (tmp, 0)); return vreinterpret_s8_u64 (lo);;
}

DEF: int16x4
vget_low_s16 (int16x8_t a)
{
  uint64x2_t tmp = vreinterpretq_u64_s16 (a); uint64x1_t lo = vcreate_u64 (vgetq_lane_u64 (tmp, 0)); return vreinterpret_s16_u64 (lo);;
}

DEF: int32x2
vget_low_s32 (int32x4_t a)
{
  uint64x2_t tmp = vreinterpretq_u64_s32 (a); uint64x1_t lo = vcreate_u64 (vgetq_lane_u64 (tmp, 0)); return vreinterpret_s32_u64 (lo);;
}

DEF: int64x1
vget_low_s64 (int64x2_t a)
{
  uint64x2_t tmp = vreinterpretq_u64_s64 (a); uint64x1_t lo = vcreate_u64 (vgetq_lane_u64 (tmp, 0)); return vreinterpret_s64_u64 (lo);;
}

DEF: uint8x8
vget_low_u8 (uint8x16_t a)
{
  uint64x2_t tmp = vreinterpretq_u64_u8 (a); uint64x1_t lo = vcreate_u64 (vgetq_lane_u64 (tmp, 0)); return vreinterpret_u8_u64 (lo);;
}

DEF: uint16x4
vget_low_u16 (uint16x8_t a)
{
  uint64x2_t tmp = vreinterpretq_u64_u16 (a); uint64x1_t lo = vcreate_u64 (vgetq_lane_u64 (tmp, 0)); return vreinterpret_u16_u64 (lo);;
}

DEF: uint32x2
vget_low_u32 (uint32x4_t a)
{
  uint64x2_t tmp = vreinterpretq_u64_u32 (a); uint64x1_t lo = vcreate_u64 (vgetq_lane_u64 (tmp, 0)); return vreinterpret_u32_u64 (lo);;
}

DEF: uint64x1
vget_low_u64 (uint64x2_t a)
{
  return vcreate_u64 (vgetq_lane_u64 (a, 0));
}
# 4657 "arm_neon.h"
DEF: float32x2
vget_high_f32 (float32x4_t a)
{
  uint64x2_t tmp = vreinterpretq_u64_f32 (a); uint64x1_t hi = vcreate_u64 (vgetq_lane_u64 (tmp, 1)); return vreinterpret_f32_u64 (hi);;
}

DEF: float64x1
vget_high_f64 (float64x2_t a)
{
  uint64x2_t tmp = vreinterpretq_u64_f64 (a); uint64x1_t hi = vcreate_u64 (vgetq_lane_u64 (tmp, 1)); return vreinterpret_f64_u64 (hi);;
}

DEF: poly8x8
vget_high_p8 (poly8x16_t a)
{
  uint64x2_t tmp = vreinterpretq_u64_p8 (a); uint64x1_t hi = vcreate_u64 (vgetq_lane_u64 (tmp, 1)); return vreinterpret_p8_u64 (hi);;
}

DEF: poly16x4
vget_high_p16 (poly16x8_t a)
{
  uint64x2_t tmp = vreinterpretq_u64_p16 (a); uint64x1_t hi = vcreate_u64 (vgetq_lane_u64 (tmp, 1)); return vreinterpret_p16_u64 (hi);;
}

DEF: int8x8
vget_high_s8 (int8x16_t a)
{
  uint64x2_t tmp = vreinterpretq_u64_s8 (a); uint64x1_t hi = vcreate_u64 (vgetq_lane_u64 (tmp, 1)); return vreinterpret_s8_u64 (hi);;
}

DEF: int16x4
vget_high_s16 (int16x8_t a)
{
  uint64x2_t tmp = vreinterpretq_u64_s16 (a); uint64x1_t hi = vcreate_u64 (vgetq_lane_u64 (tmp, 1)); return vreinterpret_s16_u64 (hi);;
}

DEF: int32x2
vget_high_s32 (int32x4_t a)
{
  uint64x2_t tmp = vreinterpretq_u64_s32 (a); uint64x1_t hi = vcreate_u64 (vgetq_lane_u64 (tmp, 1)); return vreinterpret_s32_u64 (hi);;
}

DEF: int64x1
vget_high_s64 (int64x2_t a)
{
  uint64x2_t tmp = vreinterpretq_u64_s64 (a); uint64x1_t hi = vcreate_u64 (vgetq_lane_u64 (tmp, 1)); return vreinterpret_s64_u64 (hi);;
}

DEF: uint8x8
vget_high_u8 (uint8x16_t a)
{
  uint64x2_t tmp = vreinterpretq_u64_u8 (a); uint64x1_t hi = vcreate_u64 (vgetq_lane_u64 (tmp, 1)); return vreinterpret_u8_u64 (hi);;
}

DEF: uint16x4
vget_high_u16 (uint16x8_t a)
{
  uint64x2_t tmp = vreinterpretq_u64_u16 (a); uint64x1_t hi = vcreate_u64 (vgetq_lane_u64 (tmp, 1)); return vreinterpret_u16_u64 (hi);;
}

DEF: uint32x2
vget_high_u32 (uint32x4_t a)
{
  uint64x2_t tmp = vreinterpretq_u64_u32 (a); uint64x1_t hi = vcreate_u64 (vgetq_lane_u64 (tmp, 1)); return vreinterpret_u32_u64 (hi);;
}

DEF: uint64x1
vget_high_u64 (uint64x2_t a)
{
  return vcreate_u64 (vgetq_lane_u64 (a, 1));
}

DEF: int8x16
vcombine_s8 (int8x8_t a, int8x8_t b)
{
  return (int8x16_t) builtin_aarch64_combinev8qi (a, b);
}

DEF: int16x8
vcombine_s16 (int16x4_t a, int16x4_t b)
{
  return (int16x8_t) builtin_aarch64_combinev4hi (a, b);
}

DEF: int32x4
vcombine_s32 (int32x2_t a, int32x2_t b)
{
  return (int32x4_t) builtin_aarch64_combinev2si (a, b);
}

DEF: int64x2
vcombine_s64 (int64x1_t a, int64x1_t b)
{
  return builtin_aarch64_combinedi (a[0], b[0]);
}

DEF: float32x4
vcombine_f32 (float32x2_t a, float32x2_t b)
{
  return (float32x4_t) builtin_aarch64_combinev2sf (a, b);
}

DEF: uint8x16
vcombine_u8 (uint8x8_t a, uint8x8_t b)
{
  return (uint8x16_t) builtin_aarch64_combinev8qi ((int8x8_t) a,
           (int8x8_t) b);
}

DEF: uint16x8
vcombine_u16 (uint16x4_t a, uint16x4_t b)
{
  return (uint16x8_t) builtin_aarch64_combinev4hi ((int16x4_t) a,
           (int16x4_t) b);
}

DEF: uint32x4
vcombine_u32 (uint32x2_t a, uint32x2_t b)
{
  return (uint32x4_t) builtin_aarch64_combinev2si ((int32x2_t) a,
           (int32x2_t) b);
}

DEF: uint64x2
vcombine_u64 (uint64x1_t a, uint64x1_t b)
{
  return (uint64x2_t) builtin_aarch64_combinedi (a[0], b[0]);
}

DEF: float64x2
vcombine_f64 (float64x1_t a, float64x1_t b)
{
  return builtin_aarch64_combinedf (a[0], b[0]);
}

DEF: poly8x16
vcombine_p8 (poly8x8_t a, poly8x8_t b)
{
  return (poly8x16_t) builtin_aarch64_combinev8qi ((int8x8_t) a,
           (int8x8_t) b);
}

DEF: poly16x8
vcombine_p16 (poly16x4_t a, poly16x4_t b)
{
  return (poly16x8_t) builtin_aarch64_combinev4hi ((int16x4_t) a,
           (int16x4_t) b);
}

DEF: int8x8
vaba_s8 (int8x8_t a, int8x8_t b, int8x8_t c)
{
  int8x8_t result;
  asm ("saba %0.8b,%2.8b,%3.8b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int16x4
vaba_s16 (int16x4_t a, int16x4_t b, int16x4_t c)
{
  int16x4_t result;
  asm ("saba %0.4h,%2.4h,%3.4h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int32x2
vaba_s32 (int32x2_t a, int32x2_t b, int32x2_t c)
{
  int32x2_t result;
  asm ("saba %0.2s,%2.2s,%3.2s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint8x8
vaba_u8 (uint8x8_t a, uint8x8_t b, uint8x8_t c)
{
  uint8x8_t result;
  asm ("uaba %0.8b,%2.8b,%3.8b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint16x4
vaba_u16 (uint16x4_t a, uint16x4_t b, uint16x4_t c)
{
  uint16x4_t result;
  asm ("uaba %0.4h,%2.4h,%3.4h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint32x2
vaba_u32 (uint32x2_t a, uint32x2_t b, uint32x2_t c)
{
  uint32x2_t result;
  asm ("uaba %0.2s,%2.2s,%3.2s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int16x8
vabal_high_s8 (int16x8_t a, int8x16_t b, int8x16_t c)
{
  int16x8_t result;
  asm ("sabal2 %0.8h,%2.16b,%3.16b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int32x4
vabal_high_s16 (int32x4_t a, int16x8_t b, int16x8_t c)
{
  int32x4_t result;
  asm ("sabal2 %0.4s,%2.8h,%3.8h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int64x2
vabal_high_s32 (int64x2_t a, int32x4_t b, int32x4_t c)
{
  int64x2_t result;
  asm ("sabal2 %0.2d,%2.4s,%3.4s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint16x8
vabal_high_u8 (uint16x8_t a, uint8x16_t b, uint8x16_t c)
{
  uint16x8_t result;
  asm ("uabal2 %0.8h,%2.16b,%3.16b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint32x4
vabal_high_u16 (uint32x4_t a, uint16x8_t b, uint16x8_t c)
{
  uint32x4_t result;
  asm ("uabal2 %0.4s,%2.8h,%3.8h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint64x2
vabal_high_u32 (uint64x2_t a, uint32x4_t b, uint32x4_t c)
{
  uint64x2_t result;
  asm ("uabal2 %0.2d,%2.4s,%3.4s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int16x8
vabal_s8 (int16x8_t a, int8x8_t b, int8x8_t c)
{
  int16x8_t result;
  asm ("sabal %0.8h,%2.8b,%3.8b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int32x4
vabal_s16 (int32x4_t a, int16x4_t b, int16x4_t c)
{
  int32x4_t result;
  asm ("sabal %0.4s,%2.4h,%3.4h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int64x2
vabal_s32 (int64x2_t a, int32x2_t b, int32x2_t c)
{
  int64x2_t result;
  asm ("sabal %0.2d,%2.2s,%3.2s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint16x8
vabal_u8 (uint16x8_t a, uint8x8_t b, uint8x8_t c)
{
  uint16x8_t result;
  asm ("uabal %0.8h,%2.8b,%3.8b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint32x4
vabal_u16 (uint32x4_t a, uint16x4_t b, uint16x4_t c)
{
  uint32x4_t result;
  asm ("uabal %0.4s,%2.4h,%3.4h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint64x2
vabal_u32 (uint64x2_t a, uint32x2_t b, uint32x2_t c)
{
  uint64x2_t result;
  asm ("uabal %0.2d,%2.2s,%3.2s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int8x16
vabaq_s8 (int8x16_t a, int8x16_t b, int8x16_t c)
{
  int8x16_t result;
  asm ("saba %0.16b,%2.16b,%3.16b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int16x8
vabaq_s16 (int16x8_t a, int16x8_t b, int16x8_t c)
{
  int16x8_t result;
  asm ("saba %0.8h,%2.8h,%3.8h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int32x4
vabaq_s32 (int32x4_t a, int32x4_t b, int32x4_t c)
{
  int32x4_t result;
  asm ("saba %0.4s,%2.4s,%3.4s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint8x16
vabaq_u8 (uint8x16_t a, uint8x16_t b, uint8x16_t c)
{
  uint8x16_t result;
  asm ("uaba %0.16b,%2.16b,%3.16b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint16x8
vabaq_u16 (uint16x8_t a, uint16x8_t b, uint16x8_t c)
{
  uint16x8_t result;
  asm ("uaba %0.8h,%2.8h,%3.8h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint32x4
vabaq_u32 (uint32x4_t a, uint32x4_t b, uint32x4_t c)
{
  uint32x4_t result;
  asm ("uaba %0.4s,%2.4s,%3.4s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: float32x2
vabd_f32 (float32x2_t a, float32x2_t b)
{
  float32x2_t result;
  asm ("fabd %0.2s, %1.2s, %2.2s"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int8x8
vabd_s8 (int8x8_t a, int8x8_t b)
{
  int8x8_t result;
  asm ("sabd %0.8b, %1.8b, %2.8b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int16x4
vabd_s16 (int16x4_t a, int16x4_t b)
{
  int16x4_t result;
  asm ("sabd %0.4h, %1.4h, %2.4h"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int32x2
vabd_s32 (int32x2_t a, int32x2_t b)
{
  int32x2_t result;
  asm ("sabd %0.2s, %1.2s, %2.2s"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint8x8
vabd_u8 (uint8x8_t a, uint8x8_t b)
{
  uint8x8_t result;
  asm ("uabd %0.8b, %1.8b, %2.8b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint16x4
vabd_u16 (uint16x4_t a, uint16x4_t b)
{
  uint16x4_t result;
  asm ("uabd %0.4h, %1.4h, %2.4h"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint32x2
vabd_u32 (uint32x2_t a, uint32x2_t b)
{
  uint32x2_t result;
  asm ("uabd %0.2s, %1.2s, %2.2s"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: float64
vabdd_f64 (float64_t a, float64_t b)
{
  float64_t result;
  asm ("fabd %d0, %d1, %d2"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int16x8
vabdl_high_s8 (int8x16_t a, int8x16_t b)
{
  int16x8_t result;
  asm ("sabdl2 %0.8h,%1.16b,%2.16b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int32x4
vabdl_high_s16 (int16x8_t a, int16x8_t b)
{
  int32x4_t result;
  asm ("sabdl2 %0.4s,%1.8h,%2.8h"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int64x2
vabdl_high_s32 (int32x4_t a, int32x4_t b)
{
  int64x2_t result;
  asm ("sabdl2 %0.2d,%1.4s,%2.4s"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint16x8
vabdl_high_u8 (uint8x16_t a, uint8x16_t b)
{
  uint16x8_t result;
  asm ("uabdl2 %0.8h,%1.16b,%2.16b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint32x4
vabdl_high_u16 (uint16x8_t a, uint16x8_t b)
{
  uint32x4_t result;
  asm ("uabdl2 %0.4s,%1.8h,%2.8h"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint64x2
vabdl_high_u32 (uint32x4_t a, uint32x4_t b)
{
  uint64x2_t result;
  asm ("uabdl2 %0.2d,%1.4s,%2.4s"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int16x8
vabdl_s8 (int8x8_t a, int8x8_t b)
{
  int16x8_t result;
  asm ("sabdl %0.8h, %1.8b, %2.8b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int32x4
vabdl_s16 (int16x4_t a, int16x4_t b)
{
  int32x4_t result;
  asm ("sabdl %0.4s, %1.4h, %2.4h"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int64x2
vabdl_s32 (int32x2_t a, int32x2_t b)
{
  int64x2_t result;
  asm ("sabdl %0.2d, %1.2s, %2.2s"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint16x8
vabdl_u8 (uint8x8_t a, uint8x8_t b)
{
  uint16x8_t result;
  asm ("uabdl %0.8h, %1.8b, %2.8b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint32x4
vabdl_u16 (uint16x4_t a, uint16x4_t b)
{
  uint32x4_t result;
  asm ("uabdl %0.4s, %1.4h, %2.4h"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint64x2
vabdl_u32 (uint32x2_t a, uint32x2_t b)
{
  uint64x2_t result;
  asm ("uabdl %0.2d, %1.2s, %2.2s"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: float32x4
vabdq_f32 (float32x4_t a, float32x4_t b)
{
  float32x4_t result;
  asm ("fabd %0.4s, %1.4s, %2.4s"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: float64x2
vabdq_f64 (float64x2_t a, float64x2_t b)
{
  float64x2_t result;
  asm ("fabd %0.2d, %1.2d, %2.2d"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int8x16
vabdq_s8 (int8x16_t a, int8x16_t b)
{
  int8x16_t result;
  asm ("sabd %0.16b, %1.16b, %2.16b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int16x8
vabdq_s16 (int16x8_t a, int16x8_t b)
{
  int16x8_t result;
  asm ("sabd %0.8h, %1.8h, %2.8h"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int32x4
vabdq_s32 (int32x4_t a, int32x4_t b)
{
  int32x4_t result;
  asm ("sabd %0.4s, %1.4s, %2.4s"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint8x16
vabdq_u8 (uint8x16_t a, uint8x16_t b)
{
  uint8x16_t result;
  asm ("uabd %0.16b, %1.16b, %2.16b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint16x8
vabdq_u16 (uint16x8_t a, uint16x8_t b)
{
  uint16x8_t result;
  asm ("uabd %0.8h, %1.8h, %2.8h"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint32x4
vabdq_u32 (uint32x4_t a, uint32x4_t b)
{
  uint32x4_t result;
  asm ("uabd %0.4s, %1.4s, %2.4s"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: float32
vabds_f32 (float32_t a, float32_t b)
{
  float32_t result;
  asm ("fabd %s0, %s1, %s2"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int16
vaddlv_s8 (int8x8_t a)
{
  int16_t result;
  asm ("saddlv %h0,%1.8b"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int32
vaddlv_s16 (int16x4_t a)
{
  int32_t result;
  asm ("saddlv %s0,%1.4h"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint16
vaddlv_u8 (uint8x8_t a)
{
  uint16_t result;
  asm ("uaddlv %h0,%1.8b"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint32
vaddlv_u16 (uint16x4_t a)
{
  uint32_t result;
  asm ("uaddlv %s0,%1.4h"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int16
vaddlvq_s8 (int8x16_t a)
{
  int16_t result;
  asm ("saddlv %h0,%1.16b"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int32
vaddlvq_s16 (int16x8_t a)
{
  int32_t result;
  asm ("saddlv %s0,%1.8h"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int64
vaddlvq_s32 (int32x4_t a)
{
  int64_t result;
  asm ("saddlv %d0,%1.4s"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint16
vaddlvq_u8 (uint8x16_t a)
{
  uint16_t result;
  asm ("uaddlv %h0,%1.16b"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint32
vaddlvq_u16 (uint16x8_t a)
{
  uint32_t result;
  asm ("uaddlv %s0,%1.8h"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint64
vaddlvq_u32 (uint32x4_t a)
{
  uint64_t result;
  asm ("uaddlv %d0,%1.4s"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}
# 5907 "arm_neon.h"
DEF: float32x2
vcvtx_f32_f64 (float64x2_t a)
{
  float32x2_t result;
  asm ("fcvtxn %0.2s,%1.2d"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: float32x4
vcvtx_high_f32_f64 (float32x2_t a, float64x2_t b)
{
  float32x4_t result;
  asm ("fcvtxn2 %0.4s,%1.2d"
           : "=w"(result)
           : "w" (b), "0"(a)
           : );
  return result;
}

DEF: float32
vcvtxd_f32_f64 (float64_t a)
{
  float32_t result;
  asm ("fcvtxn %s0,%d1"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: float32x2
vmla_n_f32 (float32x2_t a, float32x2_t b, float32_t c)
{
  float32x2_t result;
  float32x2_t t1;
  asm ("fmul %1.2s, %3.2s, %4.s[0]; fadd %0.2s, %0.2s, %1.2s"
           : "=w"(result), "=w"(t1)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int16x4
vmla_n_s16 (int16x4_t a, int16x4_t b, int16_t c)
{
  int16x4_t result;
  asm ("mla %0.4h,%2.4h,%3.h[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "x"(c)
           : );
  return result;
}

DEF: int32x2
vmla_n_s32 (int32x2_t a, int32x2_t b, int32_t c)
{
  int32x2_t result;
  asm ("mla %0.2s,%2.2s,%3.s[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint16x4
vmla_n_u16 (uint16x4_t a, uint16x4_t b, uint16_t c)
{
  uint16x4_t result;
  asm ("mla %0.4h,%2.4h,%3.h[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "x"(c)
           : );
  return result;
}

DEF: uint32x2
vmla_n_u32 (uint32x2_t a, uint32x2_t b, uint32_t c)
{
  uint32x2_t result;
  asm ("mla %0.2s,%2.2s,%3.s[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int8x8
vmla_s8 (int8x8_t a, int8x8_t b, int8x8_t c)
{
  int8x8_t result;
  asm ("mla %0.8b, %2.8b, %3.8b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int16x4
vmla_s16 (int16x4_t a, int16x4_t b, int16x4_t c)
{
  int16x4_t result;
  asm ("mla %0.4h, %2.4h, %3.4h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int32x2
vmla_s32 (int32x2_t a, int32x2_t b, int32x2_t c)
{
  int32x2_t result;
  asm ("mla %0.2s, %2.2s, %3.2s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint8x8
vmla_u8 (uint8x8_t a, uint8x8_t b, uint8x8_t c)
{
  uint8x8_t result;
  asm ("mla %0.8b, %2.8b, %3.8b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint16x4
vmla_u16 (uint16x4_t a, uint16x4_t b, uint16x4_t c)
{
  uint16x4_t result;
  asm ("mla %0.4h, %2.4h, %3.4h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint32x2
vmla_u32 (uint32x2_t a, uint32x2_t b, uint32x2_t c)
{
  uint32x2_t result;
  asm ("mla %0.2s, %2.2s, %3.2s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}
# 6174 "arm_neon.h"
DEF: int32x4
vmlal_high_n_s16 (int32x4_t a, int16x8_t b, int16_t c)
{
  int32x4_t result;
  asm ("smlal2 %0.4s,%2.8h,%3.h[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "x"(c)
           : );
  return result;
}

DEF: int64x2
vmlal_high_n_s32 (int64x2_t a, int32x4_t b, int32_t c)
{
  int64x2_t result;
  asm ("smlal2 %0.2d,%2.4s,%3.s[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint32x4
vmlal_high_n_u16 (uint32x4_t a, uint16x8_t b, uint16_t c)
{
  uint32x4_t result;
  asm ("umlal2 %0.4s,%2.8h,%3.h[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "x"(c)
           : );
  return result;
}

DEF: uint64x2
vmlal_high_n_u32 (uint64x2_t a, uint32x4_t b, uint32_t c)
{
  uint64x2_t result;
  asm ("umlal2 %0.2d,%2.4s,%3.s[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int16x8
vmlal_high_s8 (int16x8_t a, int8x16_t b, int8x16_t c)
{
  int16x8_t result;
  asm ("smlal2 %0.8h,%2.16b,%3.16b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int32x4
vmlal_high_s16 (int32x4_t a, int16x8_t b, int16x8_t c)
{
  int32x4_t result;
  asm ("smlal2 %0.4s,%2.8h,%3.8h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int64x2
vmlal_high_s32 (int64x2_t a, int32x4_t b, int32x4_t c)
{
  int64x2_t result;
  asm ("smlal2 %0.2d,%2.4s,%3.4s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint16x8
vmlal_high_u8 (uint16x8_t a, uint8x16_t b, uint8x16_t c)
{
  uint16x8_t result;
  asm ("umlal2 %0.8h,%2.16b,%3.16b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint32x4
vmlal_high_u16 (uint32x4_t a, uint16x8_t b, uint16x8_t c)
{
  uint32x4_t result;
  asm ("umlal2 %0.4s,%2.8h,%3.8h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint64x2
vmlal_high_u32 (uint64x2_t a, uint32x4_t b, uint32x4_t c)
{
  uint64x2_t result;
  asm ("umlal2 %0.2d,%2.4s,%3.4s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}
# 6396 "arm_neon.h"
DEF: int32x4
vmlal_n_s16 (int32x4_t a, int16x4_t b, int16_t c)
{
  int32x4_t result;
  asm ("smlal %0.4s,%2.4h,%3.h[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "x"(c)
           : );
  return result;
}

DEF: int64x2
vmlal_n_s32 (int64x2_t a, int32x2_t b, int32_t c)
{
  int64x2_t result;
  asm ("smlal %0.2d,%2.2s,%3.s[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint32x4
vmlal_n_u16 (uint32x4_t a, uint16x4_t b, uint16_t c)
{
  uint32x4_t result;
  asm ("umlal %0.4s,%2.4h,%3.h[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "x"(c)
           : );
  return result;
}

DEF: uint64x2
vmlal_n_u32 (uint64x2_t a, uint32x2_t b, uint32_t c)
{
  uint64x2_t result;
  asm ("umlal %0.2d,%2.2s,%3.s[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int16x8
vmlal_s8 (int16x8_t a, int8x8_t b, int8x8_t c)
{
  int16x8_t result;
  asm ("smlal %0.8h,%2.8b,%3.8b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int32x4
vmlal_s16 (int32x4_t a, int16x4_t b, int16x4_t c)
{
  int32x4_t result;
  asm ("smlal %0.4s,%2.4h,%3.4h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int64x2
vmlal_s32 (int64x2_t a, int32x2_t b, int32x2_t c)
{
  int64x2_t result;
  asm ("smlal %0.2d,%2.2s,%3.2s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint16x8
vmlal_u8 (uint16x8_t a, uint8x8_t b, uint8x8_t c)
{
  uint16x8_t result;
  asm ("umlal %0.8h,%2.8b,%3.8b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint32x4
vmlal_u16 (uint32x4_t a, uint16x4_t b, uint16x4_t c)
{
  uint32x4_t result;
  asm ("umlal %0.4s,%2.4h,%3.4h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint64x2
vmlal_u32 (uint64x2_t a, uint32x2_t b, uint32x2_t c)
{
  uint64x2_t result;
  asm ("umlal %0.2d,%2.2s,%3.2s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: float32x4
vmlaq_n_f32 (float32x4_t a, float32x4_t b, float32_t c)
{
  float32x4_t result;
  float32x4_t t1;
  asm ("fmul %1.4s, %3.4s, %4.s[0]; fadd %0.4s, %0.4s, %1.4s"
           : "=w"(result), "=w"(t1)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int16x8
vmlaq_n_s16 (int16x8_t a, int16x8_t b, int16_t c)
{
  int16x8_t result;
  asm ("mla %0.8h,%2.8h,%3.h[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "x"(c)
           : );
  return result;
}

DEF: int32x4
vmlaq_n_s32 (int32x4_t a, int32x4_t b, int32_t c)
{
  int32x4_t result;
  asm ("mla %0.4s,%2.4s,%3.s[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint16x8
vmlaq_n_u16 (uint16x8_t a, uint16x8_t b, uint16_t c)
{
  uint16x8_t result;
  asm ("mla %0.8h,%2.8h,%3.h[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "x"(c)
           : );
  return result;
}

DEF: uint32x4
vmlaq_n_u32 (uint32x4_t a, uint32x4_t b, uint32_t c)
{
  uint32x4_t result;
  asm ("mla %0.4s,%2.4s,%3.s[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int8x16
vmlaq_s8 (int8x16_t a, int8x16_t b, int8x16_t c)
{
  int8x16_t result;
  asm ("mla %0.16b, %2.16b, %3.16b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int16x8
vmlaq_s16 (int16x8_t a, int16x8_t b, int16x8_t c)
{
  int16x8_t result;
  asm ("mla %0.8h, %2.8h, %3.8h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int32x4
vmlaq_s32 (int32x4_t a, int32x4_t b, int32x4_t c)
{
  int32x4_t result;
  asm ("mla %0.4s, %2.4s, %3.4s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint8x16
vmlaq_u8 (uint8x16_t a, uint8x16_t b, uint8x16_t c)
{
  uint8x16_t result;
  asm ("mla %0.16b, %2.16b, %3.16b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint16x8
vmlaq_u16 (uint16x8_t a, uint16x8_t b, uint16x8_t c)
{
  uint16x8_t result;
  asm ("mla %0.8h, %2.8h, %3.8h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint32x4
vmlaq_u32 (uint32x4_t a, uint32x4_t b, uint32x4_t c)
{
  uint32x4_t result;
  asm ("mla %0.4s, %2.4s, %3.4s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: float32x2
vmls_n_f32 (float32x2_t a, float32x2_t b, float32_t c)
{
  float32x2_t result;
  float32x2_t t1;
  asm ("fmul %1.2s, %3.2s, %4.s[0]; fsub %0.2s, %0.2s, %1.2s"
           : "=w"(result), "=w"(t1)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int16x4
vmls_n_s16 (int16x4_t a, int16x4_t b, int16_t c)
{
  int16x4_t result;
  asm ("mls %0.4h, %2.4h, %3.h[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "x"(c)
           : );
  return result;
}

DEF: int32x2
vmls_n_s32 (int32x2_t a, int32x2_t b, int32_t c)
{
  int32x2_t result;
  asm ("mls %0.2s, %2.2s, %3.s[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint16x4
vmls_n_u16 (uint16x4_t a, uint16x4_t b, uint16_t c)
{
  uint16x4_t result;
  asm ("mls %0.4h, %2.4h, %3.h[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "x"(c)
           : );
  return result;
}

DEF: uint32x2
vmls_n_u32 (uint32x2_t a, uint32x2_t b, uint32_t c)
{
  uint32x2_t result;
  asm ("mls %0.2s, %2.2s, %3.s[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int8x8
vmls_s8 (int8x8_t a, int8x8_t b, int8x8_t c)
{
  int8x8_t result;
  asm ("mls %0.8b,%2.8b,%3.8b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int16x4
vmls_s16 (int16x4_t a, int16x4_t b, int16x4_t c)
{
  int16x4_t result;
  asm ("mls %0.4h,%2.4h,%3.4h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int32x2
vmls_s32 (int32x2_t a, int32x2_t b, int32x2_t c)
{
  int32x2_t result;
  asm ("mls %0.2s,%2.2s,%3.2s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint8x8
vmls_u8 (uint8x8_t a, uint8x8_t b, uint8x8_t c)
{
  uint8x8_t result;
  asm ("mls %0.8b,%2.8b,%3.8b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint16x4
vmls_u16 (uint16x4_t a, uint16x4_t b, uint16x4_t c)
{
  uint16x4_t result;
  asm ("mls %0.4h,%2.4h,%3.4h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint32x2
vmls_u32 (uint32x2_t a, uint32x2_t b, uint32x2_t c)
{
  uint32x2_t result;
  asm ("mls %0.2s,%2.2s,%3.2s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}
# 6862 "arm_neon.h"
DEF: int32x4
vmlsl_high_n_s16 (int32x4_t a, int16x8_t b, int16_t c)
{
  int32x4_t result;
  asm ("smlsl2 %0.4s, %2.8h, %3.h[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "x"(c)
           : );
  return result;
}

DEF: int64x2
vmlsl_high_n_s32 (int64x2_t a, int32x4_t b, int32_t c)
{
  int64x2_t result;
  asm ("smlsl2 %0.2d, %2.4s, %3.s[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint32x4
vmlsl_high_n_u16 (uint32x4_t a, uint16x8_t b, uint16_t c)
{
  uint32x4_t result;
  asm ("umlsl2 %0.4s, %2.8h, %3.h[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "x"(c)
           : );
  return result;
}

DEF: uint64x2
vmlsl_high_n_u32 (uint64x2_t a, uint32x4_t b, uint32_t c)
{
  uint64x2_t result;
  asm ("umlsl2 %0.2d, %2.4s, %3.s[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int16x8
vmlsl_high_s8 (int16x8_t a, int8x16_t b, int8x16_t c)
{
  int16x8_t result;
  asm ("smlsl2 %0.8h,%2.16b,%3.16b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int32x4
vmlsl_high_s16 (int32x4_t a, int16x8_t b, int16x8_t c)
{
  int32x4_t result;
  asm ("smlsl2 %0.4s,%2.8h,%3.8h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int64x2
vmlsl_high_s32 (int64x2_t a, int32x4_t b, int32x4_t c)
{
  int64x2_t result;
  asm ("smlsl2 %0.2d,%2.4s,%3.4s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint16x8
vmlsl_high_u8 (uint16x8_t a, uint8x16_t b, uint8x16_t c)
{
  uint16x8_t result;
  asm ("umlsl2 %0.8h,%2.16b,%3.16b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint32x4
vmlsl_high_u16 (uint32x4_t a, uint16x8_t b, uint16x8_t c)
{
  uint32x4_t result;
  asm ("umlsl2 %0.4s,%2.8h,%3.8h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint64x2
vmlsl_high_u32 (uint64x2_t a, uint32x4_t b, uint32x4_t c)
{
  uint64x2_t result;
  asm ("umlsl2 %0.2d,%2.4s,%3.4s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}
# 7084 "arm_neon.h"
DEF: int32x4
vmlsl_n_s16 (int32x4_t a, int16x4_t b, int16_t c)
{
  int32x4_t result;
  asm ("smlsl %0.4s, %2.4h, %3.h[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "x"(c)
           : );
  return result;
}

DEF: int64x2
vmlsl_n_s32 (int64x2_t a, int32x2_t b, int32_t c)
{
  int64x2_t result;
  asm ("smlsl %0.2d, %2.2s, %3.s[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint32x4
vmlsl_n_u16 (uint32x4_t a, uint16x4_t b, uint16_t c)
{
  uint32x4_t result;
  asm ("umlsl %0.4s, %2.4h, %3.h[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "x"(c)
           : );
  return result;
}

DEF: uint64x2
vmlsl_n_u32 (uint64x2_t a, uint32x2_t b, uint32_t c)
{
  uint64x2_t result;
  asm ("umlsl %0.2d, %2.2s, %3.s[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int16x8
vmlsl_s8 (int16x8_t a, int8x8_t b, int8x8_t c)
{
  int16x8_t result;
  asm ("smlsl %0.8h, %2.8b, %3.8b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int32x4
vmlsl_s16 (int32x4_t a, int16x4_t b, int16x4_t c)
{
  int32x4_t result;
  asm ("smlsl %0.4s, %2.4h, %3.4h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int64x2
vmlsl_s32 (int64x2_t a, int32x2_t b, int32x2_t c)
{
  int64x2_t result;
  asm ("smlsl %0.2d, %2.2s, %3.2s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint16x8
vmlsl_u8 (uint16x8_t a, uint8x8_t b, uint8x8_t c)
{
  uint16x8_t result;
  asm ("umlsl %0.8h, %2.8b, %3.8b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint32x4
vmlsl_u16 (uint32x4_t a, uint16x4_t b, uint16x4_t c)
{
  uint32x4_t result;
  asm ("umlsl %0.4s, %2.4h, %3.4h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint64x2
vmlsl_u32 (uint64x2_t a, uint32x2_t b, uint32x2_t c)
{
  uint64x2_t result;
  asm ("umlsl %0.2d, %2.2s, %3.2s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: float32x4
vmlsq_n_f32 (float32x4_t a, float32x4_t b, float32_t c)
{
  float32x4_t result;
  float32x4_t t1;
  asm ("fmul %1.4s, %3.4s, %4.s[0]; fsub %0.4s, %0.4s, %1.4s"
           : "=w"(result), "=w"(t1)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int16x8
vmlsq_n_s16 (int16x8_t a, int16x8_t b, int16_t c)
{
  int16x8_t result;
  asm ("mls %0.8h, %2.8h, %3.h[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "x"(c)
           : );
  return result;
}

DEF: int32x4
vmlsq_n_s32 (int32x4_t a, int32x4_t b, int32_t c)
{
  int32x4_t result;
  asm ("mls %0.4s, %2.4s, %3.s[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint16x8
vmlsq_n_u16 (uint16x8_t a, uint16x8_t b, uint16_t c)
{
  uint16x8_t result;
  asm ("mls %0.8h, %2.8h, %3.h[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "x"(c)
           : );
  return result;
}

DEF: uint32x4
vmlsq_n_u32 (uint32x4_t a, uint32x4_t b, uint32_t c)
{
  uint32x4_t result;
  asm ("mls %0.4s, %2.4s, %3.s[0]"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int8x16
vmlsq_s8 (int8x16_t a, int8x16_t b, int8x16_t c)
{
  int8x16_t result;
  asm ("mls %0.16b,%2.16b,%3.16b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int16x8
vmlsq_s16 (int16x8_t a, int16x8_t b, int16x8_t c)
{
  int16x8_t result;
  asm ("mls %0.8h,%2.8h,%3.8h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int32x4
vmlsq_s32 (int32x4_t a, int32x4_t b, int32x4_t c)
{
  int32x4_t result;
  asm ("mls %0.4s,%2.4s,%3.4s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint8x16
vmlsq_u8 (uint8x16_t a, uint8x16_t b, uint8x16_t c)
{
  uint8x16_t result;
  asm ("mls %0.16b,%2.16b,%3.16b"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint16x8
vmlsq_u16 (uint16x8_t a, uint16x8_t b, uint16x8_t c)
{
  uint16x8_t result;
  asm ("mls %0.8h,%2.8h,%3.8h"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: uint32x4
vmlsq_u32 (uint32x4_t a, uint32x4_t b, uint32x4_t c)
{
  uint32x4_t result;
  asm ("mls %0.4s,%2.4s,%3.4s"
           : "=w"(result)
           : "0"(a), "w"(b), "w"(c)
           : );
  return result;
}

DEF: int16x8
vmovl_high_s8 (int8x16_t a)
{
  int16x8_t result;
  asm ("sshll2 %0.8h,%1.16b,#0"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int32x4
vmovl_high_s16 (int16x8_t a)
{
  int32x4_t result;
  asm ("sshll2 %0.4s,%1.8h,#0"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int64x2
vmovl_high_s32 (int32x4_t a)
{
  int64x2_t result;
  asm ("sshll2 %0.2d,%1.4s,#0"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint16x8
vmovl_high_u8 (uint8x16_t a)
{
  uint16x8_t result;
  asm ("ushll2 %0.8h,%1.16b,#0"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint32x4
vmovl_high_u16 (uint16x8_t a)
{
  uint32x4_t result;
  asm ("ushll2 %0.4s,%1.8h,#0"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint64x2
vmovl_high_u32 (uint32x4_t a)
{
  uint64x2_t result;
  asm ("ushll2 %0.2d,%1.4s,#0"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int16x8
vmovl_s8 (int8x8_t a)
{
  int16x8_t result;
  asm ("sshll %0.8h,%1.8b,#0"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int32x4
vmovl_s16 (int16x4_t a)
{
  int32x4_t result;
  asm ("sshll %0.4s,%1.4h,#0"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int64x2
vmovl_s32 (int32x2_t a)
{
  int64x2_t result;
  asm ("sshll %0.2d,%1.2s,#0"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint16x8
vmovl_u8 (uint8x8_t a)
{
  uint16x8_t result;
  asm ("ushll %0.8h,%1.8b,#0"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint32x4
vmovl_u16 (uint16x4_t a)
{
  uint32x4_t result;
  asm ("ushll %0.4s,%1.4h,#0"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint64x2
vmovl_u32 (uint32x2_t a)
{
  uint64x2_t result;
  asm ("ushll %0.2d,%1.2s,#0"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int8x16
vmovn_high_s16 (int8x8_t a, int16x8_t b)
{
  int8x16_t result = vcombine_s8 (a, vcreate_s8 (((uint64_t) 0x0)));
  asm ("xtn2 %0.16b,%1.8h"
           : "+w"(result)
           : "w"(b)
           : );
  return result;
}

DEF: int16x8
vmovn_high_s32 (int16x4_t a, int32x4_t b)
{
  int16x8_t result = vcombine_s16 (a, vcreate_s16 (((uint64_t) 0x0)));
  asm ("xtn2 %0.8h,%1.4s"
           : "+w"(result)
           : "w"(b)
           : );
  return result;
}

DEF: int32x4
vmovn_high_s64 (int32x2_t a, int64x2_t b)
{
  int32x4_t result = vcombine_s32 (a, vcreate_s32 (((uint64_t) 0x0)));
  asm ("xtn2 %0.4s,%1.2d"
           : "+w"(result)
           : "w"(b)
           : );
  return result;
}

DEF: uint8x16
vmovn_high_u16 (uint8x8_t a, uint16x8_t b)
{
  uint8x16_t result = vcombine_u8 (a, vcreate_u8 (((uint64_t) 0x0)));
  asm ("xtn2 %0.16b,%1.8h"
           : "+w"(result)
           : "w"(b)
           : );
  return result;
}

DEF: uint16x8
vmovn_high_u32 (uint16x4_t a, uint32x4_t b)
{
  uint16x8_t result = vcombine_u16 (a, vcreate_u16 (((uint64_t) 0x0)));
  asm ("xtn2 %0.8h,%1.4s"
           : "+w"(result)
           : "w"(b)
           : );
  return result;
}

DEF: uint32x4
vmovn_high_u64 (uint32x2_t a, uint64x2_t b)
{
  uint32x4_t result = vcombine_u32 (a, vcreate_u32 (((uint64_t) 0x0)));
  asm ("xtn2 %0.4s,%1.2d"
           : "+w"(result)
           : "w"(b)
           : );
  return result;
}

DEF: int8x8
vmovn_s16 (int16x8_t a)
{
  int8x8_t result;
  asm ("xtn %0.8b,%1.8h"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int16x4
vmovn_s32 (int32x4_t a)
{
  int16x4_t result;
  asm ("xtn %0.4h,%1.4s"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int32x2
vmovn_s64 (int64x2_t a)
{
  int32x2_t result;
  asm ("xtn %0.2s,%1.2d"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint8x8
vmovn_u16 (uint16x8_t a)
{
  uint8x8_t result;
  asm ("xtn %0.8b,%1.8h"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint16x4
vmovn_u32 (uint32x4_t a)
{
  uint16x4_t result;
  asm ("xtn %0.4h,%1.4s"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint32x2
vmovn_u64 (uint64x2_t a)
{
  uint32x2_t result;
  asm ("xtn %0.2s,%1.2d"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: float32x2
vmul_n_f32 (float32x2_t a, float32_t b)
{
  float32x2_t result;
  asm ("fmul %0.2s,%1.2s,%2.s[0]"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int16x4
vmul_n_s16 (int16x4_t a, int16_t b)
{
  int16x4_t result;
  asm ("mul %0.4h,%1.4h,%2.h[0]"
           : "=w"(result)
           : "w"(a), "x"(b)
           : );
  return result;
}

DEF: int32x2
vmul_n_s32 (int32x2_t a, int32_t b)
{
  int32x2_t result;
  asm ("mul %0.2s,%1.2s,%2.s[0]"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint16x4
vmul_n_u16 (uint16x4_t a, uint16_t b)
{
  uint16x4_t result;
  asm ("mul %0.4h,%1.4h,%2.h[0]"
           : "=w"(result)
           : "w"(a), "x"(b)
           : );
  return result;
}

DEF: uint32x2
vmul_n_u32 (uint32x2_t a, uint32_t b)
{
  uint32x2_t result;
  asm ("mul %0.2s,%1.2s,%2.s[0]"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}
# 7739 "arm_neon.h"
DEF: int32x4
vmull_high_n_s16 (int16x8_t a, int16_t b)
{
  int32x4_t result;
  asm ("smull2 %0.4s,%1.8h,%2.h[0]"
           : "=w"(result)
           : "w"(a), "x"(b)
           : );
  return result;
}

DEF: int64x2
vmull_high_n_s32 (int32x4_t a, int32_t b)
{
  int64x2_t result;
  asm ("smull2 %0.2d,%1.4s,%2.s[0]"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint32x4
vmull_high_n_u16 (uint16x8_t a, uint16_t b)
{
  uint32x4_t result;
  asm ("umull2 %0.4s,%1.8h,%2.h[0]"
           : "=w"(result)
           : "w"(a), "x"(b)
           : );
  return result;
}

DEF: uint64x2
vmull_high_n_u32 (uint32x4_t a, uint32_t b)
{
  uint64x2_t result;
  asm ("umull2 %0.2d,%1.4s,%2.s[0]"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: poly16x8
vmull_high_p8 (poly8x16_t a, poly8x16_t b)
{
  poly16x8_t result;
  asm ("pmull2 %0.8h,%1.16b,%2.16b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int16x8
vmull_high_s8 (int8x16_t a, int8x16_t b)
{
  int16x8_t result;
  asm ("smull2 %0.8h,%1.16b,%2.16b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int32x4
vmull_high_s16 (int16x8_t a, int16x8_t b)
{
  int32x4_t result;
  asm ("smull2 %0.4s,%1.8h,%2.8h"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int64x2
vmull_high_s32 (int32x4_t a, int32x4_t b)
{
  int64x2_t result;
  asm ("smull2 %0.2d,%1.4s,%2.4s"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint16x8
vmull_high_u8 (uint8x16_t a, uint8x16_t b)
{
  uint16x8_t result;
  asm ("umull2 %0.8h,%1.16b,%2.16b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint32x4
vmull_high_u16 (uint16x8_t a, uint16x8_t b)
{
  uint32x4_t result;
  asm ("umull2 %0.4s,%1.8h,%2.8h"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint64x2
vmull_high_u32 (uint32x4_t a, uint32x4_t b)
{
  uint64x2_t result;
  asm ("umull2 %0.2d,%1.4s,%2.4s"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}
# 7964 "arm_neon.h"
DEF: int32x4
vmull_n_s16 (int16x4_t a, int16_t b)
{
  int32x4_t result;
  asm ("smull %0.4s,%1.4h,%2.h[0]"
           : "=w"(result)
           : "w"(a), "x"(b)
           : );
  return result;
}

DEF: int64x2
vmull_n_s32 (int32x2_t a, int32_t b)
{
  int64x2_t result;
  asm ("smull %0.2d,%1.2s,%2.s[0]"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint32x4
vmull_n_u16 (uint16x4_t a, uint16_t b)
{
  uint32x4_t result;
  asm ("umull %0.4s,%1.4h,%2.h[0]"
           : "=w"(result)
           : "w"(a), "x"(b)
           : );
  return result;
}

DEF: uint64x2
vmull_n_u32 (uint32x2_t a, uint32_t b)
{
  uint64x2_t result;
  asm ("umull %0.2d,%1.2s,%2.s[0]"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: poly16x8
vmull_p8 (poly8x8_t a, poly8x8_t b)
{
  poly16x8_t result;
  asm ("pmull %0.8h, %1.8b, %2.8b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int16x8
vmull_s8 (int8x8_t a, int8x8_t b)
{
  int16x8_t result;
  asm ("smull %0.8h, %1.8b, %2.8b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int32x4
vmull_s16 (int16x4_t a, int16x4_t b)
{
  int32x4_t result;
  asm ("smull %0.4s, %1.4h, %2.4h"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int64x2
vmull_s32 (int32x2_t a, int32x2_t b)
{
  int64x2_t result;
  asm ("smull %0.2d, %1.2s, %2.2s"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint16x8
vmull_u8 (uint8x8_t a, uint8x8_t b)
{
  uint16x8_t result;
  asm ("umull %0.8h, %1.8b, %2.8b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint32x4
vmull_u16 (uint16x4_t a, uint16x4_t b)
{
  uint32x4_t result;
  asm ("umull %0.4s, %1.4h, %2.4h"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint64x2
vmull_u32 (uint32x2_t a, uint32x2_t b)
{
  uint64x2_t result;
  asm ("umull %0.2d, %1.2s, %2.2s"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: float32x4
vmulq_n_f32 (float32x4_t a, float32_t b)
{
  float32x4_t result;
  asm ("fmul %0.4s,%1.4s,%2.s[0]"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: float64x2
vmulq_n_f64 (float64x2_t a, float64_t b)
{
  float64x2_t result;
  asm ("fmul %0.2d,%1.2d,%2.d[0]"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int16x8
vmulq_n_s16 (int16x8_t a, int16_t b)
{
  int16x8_t result;
  asm ("mul %0.8h,%1.8h,%2.h[0]"
           : "=w"(result)
           : "w"(a), "x"(b)
           : );
  return result;
}

DEF: int32x4
vmulq_n_s32 (int32x4_t a, int32_t b)
{
  int32x4_t result;
  asm ("mul %0.4s,%1.4s,%2.s[0]"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint16x8
vmulq_n_u16 (uint16x8_t a, uint16_t b)
{
  uint16x8_t result;
  asm ("mul %0.8h,%1.8h,%2.h[0]"
           : "=w"(result)
           : "w"(a), "x"(b)
           : );
  return result;
}

DEF: uint32x4
vmulq_n_u32 (uint32x4_t a, uint32_t b)
{
  uint32x4_t result;
  asm ("mul %0.4s,%1.4s,%2.s[0]"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: float32x2
vmulx_f32 (float32x2_t a, float32x2_t b)
{
  float32x2_t result;
  asm ("fmulx %0.2s,%1.2s,%2.2s"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}
# 8175 "arm_neon.h"
DEF: float64
vmulxd_f64 (float64_t a, float64_t b)
{
  float64_t result;
  asm ("fmulx %d0, %d1, %d2"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: float32x4
vmulxq_f32 (float32x4_t a, float32x4_t b)
{
  float32x4_t result;
  asm ("fmulx %0.4s,%1.4s,%2.4s"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: float64x2
vmulxq_f64 (float64x2_t a, float64x2_t b)
{
  float64x2_t result;
  asm ("fmulx %0.2d,%1.2d,%2.2d"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}
# 8234 "arm_neon.h"
DEF: float32
vmulxs_f32 (float32_t a, float32_t b)
{
  float32_t result;
  asm ("fmulx %s0, %s1, %s2"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: poly8x8
vmvn_p8 (poly8x8_t a)
{
  poly8x8_t result;
  asm ("mvn %0.8b,%1.8b"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int8x8
vmvn_s8 (int8x8_t a)
{
  int8x8_t result;
  asm ("mvn %0.8b,%1.8b"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int16x4
vmvn_s16 (int16x4_t a)
{
  int16x4_t result;
  asm ("mvn %0.8b,%1.8b"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int32x2
vmvn_s32 (int32x2_t a)
{
  int32x2_t result;
  asm ("mvn %0.8b,%1.8b"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint8x8
vmvn_u8 (uint8x8_t a)
{
  uint8x8_t result;
  asm ("mvn %0.8b,%1.8b"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint16x4
vmvn_u16 (uint16x4_t a)
{
  uint16x4_t result;
  asm ("mvn %0.8b,%1.8b"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint32x2
vmvn_u32 (uint32x2_t a)
{
  uint32x2_t result;
  asm ("mvn %0.8b,%1.8b"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: poly8x16
vmvnq_p8 (poly8x16_t a)
{
  poly8x16_t result;
  asm ("mvn %0.16b,%1.16b"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int8x16
vmvnq_s8 (int8x16_t a)
{
  int8x16_t result;
  asm ("mvn %0.16b,%1.16b"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int16x8
vmvnq_s16 (int16x8_t a)
{
  int16x8_t result;
  asm ("mvn %0.16b,%1.16b"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int32x4
vmvnq_s32 (int32x4_t a)
{
  int32x4_t result;
  asm ("mvn %0.16b,%1.16b"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint8x16
vmvnq_u8 (uint8x16_t a)
{
  uint8x16_t result;
  asm ("mvn %0.16b,%1.16b"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint16x8
vmvnq_u16 (uint16x8_t a)
{
  uint16x8_t result;
  asm ("mvn %0.16b,%1.16b"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint32x4
vmvnq_u32 (uint32x4_t a)
{
  uint32x4_t result;
  asm ("mvn %0.16b,%1.16b"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}
DEF: int16x4
vpadal_s8 (int16x4_t a, int8x8_t b)
{
  int16x4_t result;
  asm ("sadalp %0.4h,%2.8b"
           : "=w"(result)
           : "0"(a), "w"(b)
           : );
  return result;
}

DEF: int32x2
vpadal_s16 (int32x2_t a, int16x4_t b)
{
  int32x2_t result;
  asm ("sadalp %0.2s,%2.4h"
           : "=w"(result)
           : "0"(a), "w"(b)
           : );
  return result;
}

DEF: int64x1
vpadal_s32 (int64x1_t a, int32x2_t b)
{
  int64x1_t result;
  asm ("sadalp %0.1d,%2.2s"
           : "=w"(result)
           : "0"(a), "w"(b)
           : );
  return result;
}

DEF: uint16x4
vpadal_u8 (uint16x4_t a, uint8x8_t b)
{
  uint16x4_t result;
  asm ("uadalp %0.4h,%2.8b"
           : "=w"(result)
           : "0"(a), "w"(b)
           : );
  return result;
}

DEF: uint32x2
vpadal_u16 (uint32x2_t a, uint16x4_t b)
{
  uint32x2_t result;
  asm ("uadalp %0.2s,%2.4h"
           : "=w"(result)
           : "0"(a), "w"(b)
           : );
  return result;
}

DEF: uint64x1
vpadal_u32 (uint64x1_t a, uint32x2_t b)
{
  uint64x1_t result;
  asm ("uadalp %0.1d,%2.2s"
           : "=w"(result)
           : "0"(a), "w"(b)
           : );
  return result;
}

DEF: int16x8
vpadalq_s8 (int16x8_t a, int8x16_t b)
{
  int16x8_t result;
  asm ("sadalp %0.8h,%2.16b"
           : "=w"(result)
           : "0"(a), "w"(b)
           : );
  return result;
}

DEF: int32x4
vpadalq_s16 (int32x4_t a, int16x8_t b)
{
  int32x4_t result;
  asm ("sadalp %0.4s,%2.8h"
           : "=w"(result)
           : "0"(a), "w"(b)
           : );
  return result;
}

DEF: int64x2
vpadalq_s32 (int64x2_t a, int32x4_t b)
{
  int64x2_t result;
  asm ("sadalp %0.2d,%2.4s"
           : "=w"(result)
           : "0"(a), "w"(b)
           : );
  return result;
}

DEF: uint16x8
vpadalq_u8 (uint16x8_t a, uint8x16_t b)
{
  uint16x8_t result;
  asm ("uadalp %0.8h,%2.16b"
           : "=w"(result)
           : "0"(a), "w"(b)
           : );
  return result;
}

DEF: uint32x4
vpadalq_u16 (uint32x4_t a, uint16x8_t b)
{
  uint32x4_t result;
  asm ("uadalp %0.4s,%2.8h"
           : "=w"(result)
           : "0"(a), "w"(b)
           : );
  return result;
}

DEF: uint64x2
vpadalq_u32 (uint64x2_t a, uint32x4_t b)
{
  uint64x2_t result;
  asm ("uadalp %0.2d,%2.4s"
           : "=w"(result)
           : "0"(a), "w"(b)
           : );
  return result;
}

DEF: float32x2
vpadd_f32 (float32x2_t a, float32x2_t b)
{
  float32x2_t result;
  asm ("faddp %0.2s,%1.2s,%2.2s"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int16x4
vpaddl_s8 (int8x8_t a)
{
  int16x4_t result;
  asm ("saddlp %0.4h,%1.8b"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int32x2
vpaddl_s16 (int16x4_t a)
{
  int32x2_t result;
  asm ("saddlp %0.2s,%1.4h"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int64x1
vpaddl_s32 (int32x2_t a)
{
  int64x1_t result;
  asm ("saddlp %0.1d,%1.2s"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint16x4
vpaddl_u8 (uint8x8_t a)
{
  uint16x4_t result;
  asm ("uaddlp %0.4h,%1.8b"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint32x2
vpaddl_u16 (uint16x4_t a)
{
  uint32x2_t result;
  asm ("uaddlp %0.2s,%1.4h"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint64x1
vpaddl_u32 (uint32x2_t a)
{
  uint64x1_t result;
  asm ("uaddlp %0.1d,%1.2s"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int16x8
vpaddlq_s8 (int8x16_t a)
{
  int16x8_t result;
  asm ("saddlp %0.8h,%1.16b"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int32x4
vpaddlq_s16 (int16x8_t a)
{
  int32x4_t result;
  asm ("saddlp %0.4s,%1.8h"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int64x2
vpaddlq_s32 (int32x4_t a)
{
  int64x2_t result;
  asm ("saddlp %0.2d,%1.4s"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint16x8
vpaddlq_u8 (uint8x16_t a)
{
  uint16x8_t result;
  asm ("uaddlp %0.8h,%1.16b"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint32x4
vpaddlq_u16 (uint16x8_t a)
{
  uint32x4_t result;
  asm ("uaddlp %0.4s,%1.8h"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint64x2
vpaddlq_u32 (uint32x4_t a)
{
  uint64x2_t result;
  asm ("uaddlp %0.2d,%1.4s"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: float32x4
vpaddq_f32 (float32x4_t a, float32x4_t b)
{
  float32x4_t result;
  asm ("faddp %0.4s,%1.4s,%2.4s"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: float64x2
vpaddq_f64 (float64x2_t a, float64x2_t b)
{
  float64x2_t result;
  asm ("faddp %0.2d,%1.2d,%2.2d"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int8x16
vpaddq_s8 (int8x16_t a, int8x16_t b)
{
  int8x16_t result;
  asm ("addp %0.16b,%1.16b,%2.16b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int16x8
vpaddq_s16 (int16x8_t a, int16x8_t b)
{
  int16x8_t result;
  asm ("addp %0.8h,%1.8h,%2.8h"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int32x4
vpaddq_s32 (int32x4_t a, int32x4_t b)
{
  int32x4_t result;
  asm ("addp %0.4s,%1.4s,%2.4s"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int64x2
vpaddq_s64 (int64x2_t a, int64x2_t b)
{
  int64x2_t result;
  asm ("addp %0.2d,%1.2d,%2.2d"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint8x16
vpaddq_u8 (uint8x16_t a, uint8x16_t b)
{
  uint8x16_t result;
  asm ("addp %0.16b,%1.16b,%2.16b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint16x8
vpaddq_u16 (uint16x8_t a, uint16x8_t b)
{
  uint16x8_t result;
  asm ("addp %0.8h,%1.8h,%2.8h"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint32x4
vpaddq_u32 (uint32x4_t a, uint32x4_t b)
{
  uint32x4_t result;
  asm ("addp %0.4s,%1.4s,%2.4s"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint64x2
vpaddq_u64 (uint64x2_t a, uint64x2_t b)
{
  uint64x2_t result;
  asm ("addp %0.2d,%1.2d,%2.2d"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: float32
vpadds_f32 (float32x2_t a)
{
  float32_t result;
  asm ("faddp %s0,%1.2s"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: int16x4
vqdmulh_n_s16 (int16x4_t a, int16_t b)
{
  int16x4_t result;
  asm ("sqdmulh %0.4h,%1.4h,%2.h[0]"
           : "=w"(result)
           : "w"(a), "x"(b)
           : );
  return result;
}

DEF: int32x2
vqdmulh_n_s32 (int32x2_t a, int32_t b)
{
  int32x2_t result;
  asm ("sqdmulh %0.2s,%1.2s,%2.s[0]"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int16x8
vqdmulhq_n_s16 (int16x8_t a, int16_t b)
{
  int16x8_t result;
  asm ("sqdmulh %0.8h,%1.8h,%2.h[0]"
           : "=w"(result)
           : "w"(a), "x"(b)
           : );
  return result;
}

DEF: int32x4
vqdmulhq_n_s32 (int32x4_t a, int32_t b)
{
  int32x4_t result;
  asm ("sqdmulh %0.4s,%1.4s,%2.s[0]"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int8x16
vqmovn_high_s16 (int8x8_t a, int16x8_t b)
{
  int8x16_t result = vcombine_s8 (a, vcreate_s8 (((uint64_t) 0x0)));
  asm ("sqxtn2 %0.16b, %1.8h"
           : "+w"(result)
           : "w"(b)
           : );
  return result;
}

DEF: int16x8
vqmovn_high_s32 (int16x4_t a, int32x4_t b)
{
  int16x8_t result = vcombine_s16 (a, vcreate_s16 (((uint64_t) 0x0)));
  asm ("sqxtn2 %0.8h, %1.4s"
           : "+w"(result)
           : "w"(b)
           : );
  return result;
}

DEF: int32x4
vqmovn_high_s64 (int32x2_t a, int64x2_t b)
{
  int32x4_t result = vcombine_s32 (a, vcreate_s32 (((uint64_t) 0x0)));
  asm ("sqxtn2 %0.4s, %1.2d"
           : "+w"(result)
           : "w"(b)
           : );
  return result;
}

DEF: uint8x16
vqmovn_high_u16 (uint8x8_t a, uint16x8_t b)
{
  uint8x16_t result = vcombine_u8 (a, vcreate_u8 (((uint64_t) 0x0)));
  asm ("uqxtn2 %0.16b, %1.8h"
           : "+w"(result)
           : "w"(b)
           : );
  return result;
}

DEF: uint16x8
vqmovn_high_u32 (uint16x4_t a, uint32x4_t b)
{
  uint16x8_t result = vcombine_u16 (a, vcreate_u16 (((uint64_t) 0x0)));
  asm ("uqxtn2 %0.8h, %1.4s"
           : "+w"(result)
           : "w"(b)
           : );
  return result;
}

DEF: uint32x4
vqmovn_high_u64 (uint32x2_t a, uint64x2_t b)
{
  uint32x4_t result = vcombine_u32 (a, vcreate_u32 (((uint64_t) 0x0)));
  asm ("uqxtn2 %0.4s, %1.2d"
           : "+w"(result)
           : "w"(b)
           : );
  return result;
}

DEF: uint8x16
vqmovun_high_s16 (uint8x8_t a, int16x8_t b)
{
  uint8x16_t result = vcombine_u8 (a, vcreate_u8 (((uint64_t) 0x0)));
  asm ("sqxtun2 %0.16b, %1.8h"
           : "+w"(result)
           : "w"(b)
           : );
  return result;
}

DEF: uint16x8
vqmovun_high_s32 (uint16x4_t a, int32x4_t b)
{
  uint16x8_t result = vcombine_u16 (a, vcreate_u16 (((uint64_t) 0x0)));
  asm ("sqxtun2 %0.8h, %1.4s"
           : "+w"(result)
           : "w"(b)
           : );
  return result;
}

DEF: uint32x4
vqmovun_high_s64 (uint32x2_t a, int64x2_t b)
{
  uint32x4_t result = vcombine_u32 (a, vcreate_u32 (((uint64_t) 0x0)));
  asm ("sqxtun2 %0.4s, %1.2d"
           : "+w"(result)
           : "w"(b)
           : );
  return result;
}

DEF: int16x4
vqrdmulh_n_s16 (int16x4_t a, int16_t b)
{
  int16x4_t result;
  asm ("sqrdmulh %0.4h,%1.4h,%2.h[0]"
           : "=w"(result)
           : "w"(a), "x"(b)
           : );
  return result;
}

DEF: int32x2
vqrdmulh_n_s32 (int32x2_t a, int32_t b)
{
  int32x2_t result;
  asm ("sqrdmulh %0.2s,%1.2s,%2.s[0]"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int16x8
vqrdmulhq_n_s16 (int16x8_t a, int16_t b)
{
  int16x8_t result;
  asm ("sqrdmulh %0.8h,%1.8h,%2.h[0]"
           : "=w"(result)
           : "w"(a), "x"(b)
           : );
  return result;
}

DEF: int32x4
vqrdmulhq_n_s32 (int32x4_t a, int32_t b)
{
  int32x4_t result;
  asm ("sqrdmulh %0.4s,%1.4s,%2.s[0]"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}
# 9415 "arm_neon.h"
DEF: float32x2
vrsqrte_f32 (float32x2_t a)
{
  float32x2_t result;
  asm ("frsqrte %0.2s,%1.2s"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: float64x1
vrsqrte_f64 (float64x1_t a)
{
  float64x1_t result;
  asm ("frsqrte %d0,%d1"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint32x2
vrsqrte_u32 (uint32x2_t a)
{
  uint32x2_t result;
  asm ("ursqrte %0.2s,%1.2s"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: float64
vrsqrted_f64 (float64_t a)
{
  float64_t result;
  asm ("frsqrte %d0,%d1"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: float32x4
vrsqrteq_f32 (float32x4_t a)
{
  float32x4_t result;
  asm ("frsqrte %0.4s,%1.4s"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: float64x2
vrsqrteq_f64 (float64x2_t a)
{
  float64x2_t result;
  asm ("frsqrte %0.2d,%1.2d"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: uint32x4
vrsqrteq_u32 (uint32x4_t a)
{
  uint32x4_t result;
  asm ("ursqrte %0.4s,%1.4s"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: float32
vrsqrtes_f32 (float32_t a)
{
  float32_t result;
  asm ("frsqrte %s0,%s1"
           : "=w"(result)
           : "w"(a)
           : );
  return result;
}

DEF: float32x2
vrsqrts_f32 (float32x2_t a, float32x2_t b)
{
  float32x2_t result;
  asm ("frsqrts %0.2s,%1.2s,%2.2s"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: float64
vrsqrtsd_f64 (float64_t a, float64_t b)
{
  float64_t result;
  asm ("frsqrts %d0,%d1,%d2"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: float32x4
vrsqrtsq_f32 (float32x4_t a, float32x4_t b)
{
  float32x4_t result;
  asm ("frsqrts %0.4s,%1.4s,%2.4s"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: float64x2
vrsqrtsq_f64 (float64x2_t a, float64x2_t b)
{
  float64x2_t result;
  asm ("frsqrts %0.2d,%1.2d,%2.2d"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: float32
vrsqrtss_f32 (float32_t a, float32_t b)
{
  float32_t result;
  asm ("frsqrts %s0,%s1,%s2"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}
# 9824 "arm_neon.h"
DEF: uint8x8
vtst_p8 (poly8x8_t a, poly8x8_t b)
{
  uint8x8_t result;
  asm ("cmtst %0.8b, %1.8b, %2.8b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint16x4
vtst_p16 (poly16x4_t a, poly16x4_t b)
{
  uint16x4_t result;
  asm ("cmtst %0.4h, %1.4h, %2.4h"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint8x16
vtstq_p8 (poly8x16_t a, poly8x16_t b)
{
  uint8x16_t result;
  asm ("cmtst %0.16b, %1.16b, %2.16b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint16x8
vtstq_p16 (poly16x8_t a, poly16x8_t b)
{
  uint16x8_t result;
  asm ("cmtst %0.8h, %1.8h, %2.8h"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}
# 9923 "arm_neon.h"
typedef struct int8x2_t { int8_t val[2]; } int8x2_t;
typedef struct int16x2_t { int16_t val[2]; } int16x2_t;
typedef struct uint8x2_t { uint8_t val[2]; } uint8x2_t;
typedef struct uint16x2_t { uint16_t val[2]; } uint16x2_t;
typedef struct poly8x2_t { poly8_t val[2]; } poly8x2_t;
typedef struct poly16x2_t { poly16_t val[2]; } poly16x2_t;

typedef struct int8x3_t { int8_t val[3]; } int8x3_t;
typedef struct int16x3_t { int16_t val[3]; } int16x3_t;
typedef struct int32x3_t { int32_t val[3]; } int32x3_t;
typedef struct int64x3_t { int64_t val[3]; } int64x3_t;
typedef struct uint8x3_t { uint8_t val[3]; } uint8x3_t;
typedef struct uint16x3_t { uint16_t val[3]; } uint16x3_t;
typedef struct uint32x3_t { uint32_t val[3]; } uint32x3_t;
typedef struct uint64x3_t { uint64_t val[3]; } uint64x3_t;
typedef struct float32x3_t { float32_t val[3]; } float32x3_t;
typedef struct float64x3_t { float64_t val[3]; } float64x3_t;
typedef struct poly8x3_t { poly8_t val[3]; } poly8x3_t;
typedef struct poly16x3_t { poly16_t val[3]; } poly16x3_t;

typedef struct int8x4_t { int8_t val[4]; } int8x4_t;
typedef struct int64x4_t { int64_t val[4]; } int64x4_t;
typedef struct uint8x4_t { uint8_t val[4]; } uint8x4_t;
typedef struct uint64x4_t { uint64_t val[4]; } uint64x4_t;
typedef struct poly8x4_t { poly8_t val[4]; } poly8x4_t;
typedef struct float64x4_t { float64_t val[4]; } float64x4_t;
# 9975 "arm_neon.h"
DEF: void vst2_lane_f32 (float32_t *ptr, float32x2x2_t b, const int c) { builtin_aarch64_simd_oi o; float32x4x2_t temp; temp.val[0] = vcombine_f32 (b.val[0], vcreate_f32 (((uint64_t) 0))); temp.val[1] = vcombine_f32 (b.val[1], vcreate_f32 (((uint64_t) 0))); o = builtin_aarch64_set_qregoiv4sf (o, (float32x4_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv4sf (o, (float32x4_t) temp.val[1], 1); builtin_aarch64_st2_lanev2sf ((builtin_aarch64_simd_sf *) ptr, o, c); }

DEF: void vst2_lane_f64 (float64_t *ptr, float64x1x2_t b, const int c) { builtin_aarch64_simd_oi o; float64x2x2_t temp; temp.val[0] = vcombine_f64 (b.val[0], vcreate_f64 (((uint64_t) 0))); temp.val[1] = vcombine_f64 (b.val[1], vcreate_f64 (((uint64_t) 0))); o = builtin_aarch64_set_qregoiv2df (o, (float64x2_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv2df (o, (float64x2_t) temp.val[1], 1); builtin_aarch64_st2_lanedf ((builtin_aarch64_simd_df *) ptr, o, c); }

DEF: void vst2_lane_p8 (poly8_t *ptr, poly8x8x2_t b, const int c) { builtin_aarch64_simd_oi o; poly8x16x2_t temp; temp.val[0] = vcombine_p8 (b.val[0], vcreate_p8 (((uint64_t) 0))); temp.val[1] = vcombine_p8 (b.val[1], vcreate_p8 (((uint64_t) 0))); o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) temp.val[1], 1); builtin_aarch64_st2_lanev8qi ((builtin_aarch64_simd_qi *) ptr, o, c); }

DEF: void vst2_lane_p16 (poly16_t *ptr, poly16x4x2_t b, const int c) { builtin_aarch64_simd_oi o; poly16x8x2_t temp; temp.val[0] = vcombine_p16 (b.val[0], vcreate_p16 (((uint64_t) 0))); temp.val[1] = vcombine_p16 (b.val[1], vcreate_p16 (((uint64_t) 0))); o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) temp.val[1], 1); builtin_aarch64_st2_lanev4hi ((builtin_aarch64_simd_hi *) ptr, o, c); }

DEF: void vst2_lane_s8 (int8_t *ptr, int8x8x2_t b, const int c) { builtin_aarch64_simd_oi o; int8x16x2_t temp; temp.val[0] = vcombine_s8 (b.val[0], vcreate_s8 (((uint64_t) 0))); temp.val[1] = vcombine_s8 (b.val[1], vcreate_s8 (((uint64_t) 0))); o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) temp.val[1], 1); builtin_aarch64_st2_lanev8qi ((builtin_aarch64_simd_qi *) ptr, o, c); }

DEF: void vst2_lane_s16 (int16_t *ptr, int16x4x2_t b, const int c) { builtin_aarch64_simd_oi o; int16x8x2_t temp; temp.val[0] = vcombine_s16 (b.val[0], vcreate_s16 (((uint64_t) 0))); temp.val[1] = vcombine_s16 (b.val[1], vcreate_s16 (((uint64_t) 0))); o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) temp.val[1], 1); builtin_aarch64_st2_lanev4hi ((builtin_aarch64_simd_hi *) ptr, o, c); }

DEF: void vst2_lane_s32 (int32_t *ptr, int32x2x2_t b, const int c) { builtin_aarch64_simd_oi o; int32x4x2_t temp; temp.val[0] = vcombine_s32 (b.val[0], vcreate_s32 (((uint64_t) 0))); temp.val[1] = vcombine_s32 (b.val[1], vcreate_s32 (((uint64_t) 0))); o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) temp.val[1], 1); builtin_aarch64_st2_lanev2si ((builtin_aarch64_simd_si *) ptr, o, c); }

DEF: void vst2_lane_s64 (int64_t *ptr, int64x1x2_t b, const int c) { builtin_aarch64_simd_oi o; int64x2x2_t temp; temp.val[0] = vcombine_s64 (b.val[0], vcreate_s64 (((uint64_t) 0))); temp.val[1] = vcombine_s64 (b.val[1], vcreate_s64 (((uint64_t) 0))); o = builtin_aarch64_set_qregoiv2di (o, (int64x2_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv2di (o, (int64x2_t) temp.val[1], 1); builtin_aarch64_st2_lanedi ((builtin_aarch64_simd_di *) ptr, o, c); }

DEF: void vst2_lane_u8 (uint8_t *ptr, uint8x8x2_t b, const int c) { builtin_aarch64_simd_oi o; uint8x16x2_t temp; temp.val[0] = vcombine_u8 (b.val[0], vcreate_u8 (((uint64_t) 0))); temp.val[1] = vcombine_u8 (b.val[1], vcreate_u8 (((uint64_t) 0))); o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) temp.val[1], 1); builtin_aarch64_st2_lanev8qi ((builtin_aarch64_simd_qi *) ptr, o, c); }

DEF: void vst2_lane_u16 (uint16_t *ptr, uint16x4x2_t b, const int c) { builtin_aarch64_simd_oi o; uint16x8x2_t temp; temp.val[0] = vcombine_u16 (b.val[0], vcreate_u16 (((uint64_t) 0))); temp.val[1] = vcombine_u16 (b.val[1], vcreate_u16 (((uint64_t) 0))); o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) temp.val[1], 1); builtin_aarch64_st2_lanev4hi ((builtin_aarch64_simd_hi *) ptr, o, c); }

DEF: void vst2_lane_u32 (uint32_t *ptr, uint32x2x2_t b, const int c) { builtin_aarch64_simd_oi o; uint32x4x2_t temp; temp.val[0] = vcombine_u32 (b.val[0], vcreate_u32 (((uint64_t) 0))); temp.val[1] = vcombine_u32 (b.val[1], vcreate_u32 (((uint64_t) 0))); o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) temp.val[1], 1); builtin_aarch64_st2_lanev2si ((builtin_aarch64_simd_si *) ptr, o, c); }

DEF: void vst2_lane_u64 (uint64_t *ptr, uint64x1x2_t b, const int c) { builtin_aarch64_simd_oi o; uint64x2x2_t temp; temp.val[0] = vcombine_u64 (b.val[0], vcreate_u64 (((uint64_t) 0))); temp.val[1] = vcombine_u64 (b.val[1], vcreate_u64 (((uint64_t) 0))); o = builtin_aarch64_set_qregoiv2di (o, (int64x2_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv2di (o, (int64x2_t) temp.val[1], 1); builtin_aarch64_st2_lanedi ((builtin_aarch64_simd_di *) ptr, o, c); }
# 10013 "arm_neon.h"
DEF: void vst2q_lane_f32 (float32_t *ptr, float32x4x2_t b, const int c) { union { float32x4x2_t i; builtin_aarch64_simd_oi o; } temp = { b }; builtin_aarch64_st2_lanev4sf ((builtin_aarch64_simd_sf *) ptr, temp.o, c); }
DEF: void vst2q_lane_f64 (float64_t *ptr, float64x2x2_t b, const int c) { union { float64x2x2_t i; builtin_aarch64_simd_oi o; } temp = { b }; builtin_aarch64_st2_lanev2df ((builtin_aarch64_simd_df *) ptr, temp.o, c); }
DEF: void vst2q_lane_p8 (poly8_t *ptr, poly8x16x2_t b, const int c) { union { poly8x16x2_t i; builtin_aarch64_simd_oi o; } temp = { b }; builtin_aarch64_st2_lanev16qi ((builtin_aarch64_simd_qi *) ptr, temp.o, c); }
DEF: void vst2q_lane_p16 (poly16_t *ptr, poly16x8x2_t b, const int c) { union { poly16x8x2_t i; builtin_aarch64_simd_oi o; } temp = { b }; builtin_aarch64_st2_lanev8hi ((builtin_aarch64_simd_hi *) ptr, temp.o, c); }
DEF: void vst2q_lane_s8 (int8_t *ptr, int8x16x2_t b, const int c) { union { int8x16x2_t i; builtin_aarch64_simd_oi o; } temp = { b }; builtin_aarch64_st2_lanev16qi ((builtin_aarch64_simd_qi *) ptr, temp.o, c); }
DEF: void vst2q_lane_s16 (int16_t *ptr, int16x8x2_t b, const int c) { union { int16x8x2_t i; builtin_aarch64_simd_oi o; } temp = { b }; builtin_aarch64_st2_lanev8hi ((builtin_aarch64_simd_hi *) ptr, temp.o, c); }
DEF: void vst2q_lane_s32 (int32_t *ptr, int32x4x2_t b, const int c) { union { int32x4x2_t i; builtin_aarch64_simd_oi o; } temp = { b }; builtin_aarch64_st2_lanev4si ((builtin_aarch64_simd_si *) ptr, temp.o, c); }
DEF: void vst2q_lane_s64 (int64_t *ptr, int64x2x2_t b, const int c) { union { int64x2x2_t i; builtin_aarch64_simd_oi o; } temp = { b }; builtin_aarch64_st2_lanev2di ((builtin_aarch64_simd_di *) ptr, temp.o, c); }
DEF: void vst2q_lane_u8 (uint8_t *ptr, uint8x16x2_t b, const int c) { union { uint8x16x2_t i; builtin_aarch64_simd_oi o; } temp = { b }; builtin_aarch64_st2_lanev16qi ((builtin_aarch64_simd_qi *) ptr, temp.o, c); }
DEF: void vst2q_lane_u16 (uint16_t *ptr, uint16x8x2_t b, const int c) { union { uint16x8x2_t i; builtin_aarch64_simd_oi o; } temp = { b }; builtin_aarch64_st2_lanev8hi ((builtin_aarch64_simd_hi *) ptr, temp.o, c); }
DEF: void vst2q_lane_u32 (uint32_t *ptr, uint32x4x2_t b, const int c) { union { uint32x4x2_t i; builtin_aarch64_simd_oi o; } temp = { b }; builtin_aarch64_st2_lanev4si ((builtin_aarch64_simd_si *) ptr, temp.o, c); }
DEF: void vst2q_lane_u64 (uint64_t *ptr, uint64x2x2_t b, const int c) { union { uint64x2x2_t i; builtin_aarch64_simd_oi o; } temp = { b }; builtin_aarch64_st2_lanev2di ((builtin_aarch64_simd_di *) ptr, temp.o, c); }
# 10054 "arm_neon.h"
DEF: void vst3_lane_f32 (float32_t *ptr, float32x2x3_t b, const int c) { builtin_aarch64_simd_ci o; float32x4x3_t temp; temp.val[0] = vcombine_f32 (b.val[0], vcreate_f32 (((uint64_t) 0))); temp.val[1] = vcombine_f32 (b.val[1], vcreate_f32 (((uint64_t) 0))); temp.val[2] = vcombine_f32 (b.val[2], vcreate_f32 (((uint64_t) 0))); o = builtin_aarch64_set_qregciv4sf (o, (float32x4_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv4sf (o, (float32x4_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv4sf (o, (float32x4_t) temp.val[2], 2); builtin_aarch64_st3_lanev2sf ((builtin_aarch64_simd_sf *) ptr, o, c); }

DEF: void vst3_lane_f64 (float64_t *ptr, float64x1x3_t b, const int c) { builtin_aarch64_simd_ci o; float64x2x3_t temp; temp.val[0] = vcombine_f64 (b.val[0], vcreate_f64 (((uint64_t) 0))); temp.val[1] = vcombine_f64 (b.val[1], vcreate_f64 (((uint64_t) 0))); temp.val[2] = vcombine_f64 (b.val[2], vcreate_f64 (((uint64_t) 0))); o = builtin_aarch64_set_qregciv2df (o, (float64x2_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv2df (o, (float64x2_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv2df (o, (float64x2_t) temp.val[2], 2); builtin_aarch64_st3_lanedf ((builtin_aarch64_simd_df *) ptr, o, c); }

DEF: void vst3_lane_p8 (poly8_t *ptr, poly8x8x3_t b, const int c) { builtin_aarch64_simd_ci o; poly8x16x3_t temp; temp.val[0] = vcombine_p8 (b.val[0], vcreate_p8 (((uint64_t) 0))); temp.val[1] = vcombine_p8 (b.val[1], vcreate_p8 (((uint64_t) 0))); temp.val[2] = vcombine_p8 (b.val[2], vcreate_p8 (((uint64_t) 0))); o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[2], 2); builtin_aarch64_st3_lanev8qi ((builtin_aarch64_simd_qi *) ptr, o, c); }

DEF: void vst3_lane_p16 (poly16_t *ptr, poly16x4x3_t b, const int c) { builtin_aarch64_simd_ci o; poly16x8x3_t temp; temp.val[0] = vcombine_p16 (b.val[0], vcreate_p16 (((uint64_t) 0))); temp.val[1] = vcombine_p16 (b.val[1], vcreate_p16 (((uint64_t) 0))); temp.val[2] = vcombine_p16 (b.val[2], vcreate_p16 (((uint64_t) 0))); o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[2], 2); builtin_aarch64_st3_lanev4hi ((builtin_aarch64_simd_hi *) ptr, o, c); }

DEF: void vst3_lane_s8 (int8_t *ptr, int8x8x3_t b, const int c) { builtin_aarch64_simd_ci o; int8x16x3_t temp; temp.val[0] = vcombine_s8 (b.val[0], vcreate_s8 (((uint64_t) 0))); temp.val[1] = vcombine_s8 (b.val[1], vcreate_s8 (((uint64_t) 0))); temp.val[2] = vcombine_s8 (b.val[2], vcreate_s8 (((uint64_t) 0))); o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[2], 2); builtin_aarch64_st3_lanev8qi ((builtin_aarch64_simd_qi *) ptr, o, c); }

DEF: void vst3_lane_s16 (int16_t *ptr, int16x4x3_t b, const int c) { builtin_aarch64_simd_ci o; int16x8x3_t temp; temp.val[0] = vcombine_s16 (b.val[0], vcreate_s16 (((uint64_t) 0))); temp.val[1] = vcombine_s16 (b.val[1], vcreate_s16 (((uint64_t) 0))); temp.val[2] = vcombine_s16 (b.val[2], vcreate_s16 (((uint64_t) 0))); o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[2], 2); builtin_aarch64_st3_lanev4hi ((builtin_aarch64_simd_hi *) ptr, o, c); }

DEF: void vst3_lane_s32 (int32_t *ptr, int32x2x3_t b, const int c) { builtin_aarch64_simd_ci o; int32x4x3_t temp; temp.val[0] = vcombine_s32 (b.val[0], vcreate_s32 (((uint64_t) 0))); temp.val[1] = vcombine_s32 (b.val[1], vcreate_s32 (((uint64_t) 0))); temp.val[2] = vcombine_s32 (b.val[2], vcreate_s32 (((uint64_t) 0))); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) temp.val[2], 2); builtin_aarch64_st3_lanev2si ((builtin_aarch64_simd_si *) ptr, o, c); }

DEF: void vst3_lane_s64 (int64_t *ptr, int64x1x3_t b, const int c) { builtin_aarch64_simd_ci o; int64x2x3_t temp; temp.val[0] = vcombine_s64 (b.val[0], vcreate_s64 (((uint64_t) 0))); temp.val[1] = vcombine_s64 (b.val[1], vcreate_s64 (((uint64_t) 0))); temp.val[2] = vcombine_s64 (b.val[2], vcreate_s64 (((uint64_t) 0))); o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) temp.val[2], 2); builtin_aarch64_st3_lanedi ((builtin_aarch64_simd_di *) ptr, o, c); }

DEF: void vst3_lane_u8 (uint8_t *ptr, uint8x8x3_t b, const int c) { builtin_aarch64_simd_ci o; uint8x16x3_t temp; temp.val[0] = vcombine_u8 (b.val[0], vcreate_u8 (((uint64_t) 0))); temp.val[1] = vcombine_u8 (b.val[1], vcreate_u8 (((uint64_t) 0))); temp.val[2] = vcombine_u8 (b.val[2], vcreate_u8 (((uint64_t) 0))); o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[2], 2); builtin_aarch64_st3_lanev8qi ((builtin_aarch64_simd_qi *) ptr, o, c); }

DEF: void vst3_lane_u16 (uint16_t *ptr, uint16x4x3_t b, const int c) { builtin_aarch64_simd_ci o; uint16x8x3_t temp; temp.val[0] = vcombine_u16 (b.val[0], vcreate_u16 (((uint64_t) 0))); temp.val[1] = vcombine_u16 (b.val[1], vcreate_u16 (((uint64_t) 0))); temp.val[2] = vcombine_u16 (b.val[2], vcreate_u16 (((uint64_t) 0))); o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[2], 2); builtin_aarch64_st3_lanev4hi ((builtin_aarch64_simd_hi *) ptr, o, c); }

DEF: void vst3_lane_u32 (uint32_t *ptr, uint32x2x3_t b, const int c) { builtin_aarch64_simd_ci o; uint32x4x3_t temp; temp.val[0] = vcombine_u32 (b.val[0], vcreate_u32 (((uint64_t) 0))); temp.val[1] = vcombine_u32 (b.val[1], vcreate_u32 (((uint64_t) 0))); temp.val[2] = vcombine_u32 (b.val[2], vcreate_u32 (((uint64_t) 0))); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) temp.val[2], 2); builtin_aarch64_st3_lanev2si ((builtin_aarch64_simd_si *) ptr, o, c); }

DEF: void vst3_lane_u64 (uint64_t *ptr, uint64x1x3_t b, const int c) { builtin_aarch64_simd_ci o; uint64x2x3_t temp; temp.val[0] = vcombine_u64 (b.val[0], vcreate_u64 (((uint64_t) 0))); temp.val[1] = vcombine_u64 (b.val[1], vcreate_u64 (((uint64_t) 0))); temp.val[2] = vcombine_u64 (b.val[2], vcreate_u64 (((uint64_t) 0))); o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) temp.val[2], 2); builtin_aarch64_st3_lanedi ((builtin_aarch64_simd_di *) ptr, o, c); }
# 10092 "arm_neon.h"
DEF: void vst3q_lane_f32 (float32_t *ptr, float32x4x3_t b, const int c) { union { float32x4x3_t i; builtin_aarch64_simd_ci o; } temp = { b }; builtin_aarch64_st3_lanev4sf ((builtin_aarch64_simd_sf *) ptr, temp.o, c); }
DEF: void vst3q_lane_f64 (float64_t *ptr, float64x2x3_t b, const int c) { union { float64x2x3_t i; builtin_aarch64_simd_ci o; } temp = { b }; builtin_aarch64_st3_lanev2df ((builtin_aarch64_simd_df *) ptr, temp.o, c); }
DEF: void vst3q_lane_p8 (poly8_t *ptr, poly8x16x3_t b, const int c) { union { poly8x16x3_t i; builtin_aarch64_simd_ci o; } temp = { b }; builtin_aarch64_st3_lanev16qi ((builtin_aarch64_simd_qi *) ptr, temp.o, c); }
DEF: void vst3q_lane_p16 (poly16_t *ptr, poly16x8x3_t b, const int c) { union { poly16x8x3_t i; builtin_aarch64_simd_ci o; } temp = { b }; builtin_aarch64_st3_lanev8hi ((builtin_aarch64_simd_hi *) ptr, temp.o, c); }
DEF: void vst3q_lane_s8 (int8_t *ptr, int8x16x3_t b, const int c) { union { int8x16x3_t i; builtin_aarch64_simd_ci o; } temp = { b }; builtin_aarch64_st3_lanev16qi ((builtin_aarch64_simd_qi *) ptr, temp.o, c); }
DEF: void vst3q_lane_s16 (int16_t *ptr, int16x8x3_t b, const int c) { union { int16x8x3_t i; builtin_aarch64_simd_ci o; } temp = { b }; builtin_aarch64_st3_lanev8hi ((builtin_aarch64_simd_hi *) ptr, temp.o, c); }
DEF: void vst3q_lane_s32 (int32_t *ptr, int32x4x3_t b, const int c) { union { int32x4x3_t i; builtin_aarch64_simd_ci o; } temp = { b }; builtin_aarch64_st3_lanev4si ((builtin_aarch64_simd_si *) ptr, temp.o, c); }
DEF: void vst3q_lane_s64 (int64_t *ptr, int64x2x3_t b, const int c) { union { int64x2x3_t i; builtin_aarch64_simd_ci o; } temp = { b }; builtin_aarch64_st3_lanev2di ((builtin_aarch64_simd_di *) ptr, temp.o, c); }
DEF: void vst3q_lane_u8 (uint8_t *ptr, uint8x16x3_t b, const int c) { union { uint8x16x3_t i; builtin_aarch64_simd_ci o; } temp = { b }; builtin_aarch64_st3_lanev16qi ((builtin_aarch64_simd_qi *) ptr, temp.o, c); }
DEF: void vst3q_lane_u16 (uint16_t *ptr, uint16x8x3_t b, const int c) { union { uint16x8x3_t i; builtin_aarch64_simd_ci o; } temp = { b }; builtin_aarch64_st3_lanev8hi ((builtin_aarch64_simd_hi *) ptr, temp.o, c); }
DEF: void vst3q_lane_u32 (uint32_t *ptr, uint32x4x3_t b, const int c) { union { uint32x4x3_t i; builtin_aarch64_simd_ci o; } temp = { b }; builtin_aarch64_st3_lanev4si ((builtin_aarch64_simd_si *) ptr, temp.o, c); }
DEF: void vst3q_lane_u64 (uint64_t *ptr, uint64x2x3_t b, const int c) { union { uint64x2x3_t i; builtin_aarch64_simd_ci o; } temp = { b }; builtin_aarch64_st3_lanev2di ((builtin_aarch64_simd_di *) ptr, temp.o, c); }
# 10138 "arm_neon.h"
DEF: void vst4_lane_f32 (float32_t *ptr, float32x2x4_t b, const int c) { builtin_aarch64_simd_xi o; float32x4x4_t temp; temp.val[0] = vcombine_f32 (b.val[0], vcreate_f32 (((uint64_t) 0))); temp.val[1] = vcombine_f32 (b.val[1], vcreate_f32 (((uint64_t) 0))); temp.val[2] = vcombine_f32 (b.val[2], vcreate_f32 (((uint64_t) 0))); temp.val[3] = vcombine_f32 (b.val[3], vcreate_f32 (((uint64_t) 0))); o = builtin_aarch64_set_qregxiv4sf (o, (float32x4_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv4sf (o, (float32x4_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv4sf (o, (float32x4_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv4sf (o, (float32x4_t) temp.val[3], 3); builtin_aarch64_st4_lanev2sf ((builtin_aarch64_simd_sf *) ptr, o, c); }

DEF: void vst4_lane_f64 (float64_t *ptr, float64x1x4_t b, const int c) { builtin_aarch64_simd_xi o; float64x2x4_t temp; temp.val[0] = vcombine_f64 (b.val[0], vcreate_f64 (((uint64_t) 0))); temp.val[1] = vcombine_f64 (b.val[1], vcreate_f64 (((uint64_t) 0))); temp.val[2] = vcombine_f64 (b.val[2], vcreate_f64 (((uint64_t) 0))); temp.val[3] = vcombine_f64 (b.val[3], vcreate_f64 (((uint64_t) 0))); o = builtin_aarch64_set_qregxiv2df (o, (float64x2_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv2df (o, (float64x2_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv2df (o, (float64x2_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv2df (o, (float64x2_t) temp.val[3], 3); builtin_aarch64_st4_lanedf ((builtin_aarch64_simd_df *) ptr, o, c); }

DEF: void vst4_lane_p8 (poly8_t *ptr, poly8x8x4_t b, const int c) { builtin_aarch64_simd_xi o; poly8x16x4_t temp; temp.val[0] = vcombine_p8 (b.val[0], vcreate_p8 (((uint64_t) 0))); temp.val[1] = vcombine_p8 (b.val[1], vcreate_p8 (((uint64_t) 0))); temp.val[2] = vcombine_p8 (b.val[2], vcreate_p8 (((uint64_t) 0))); temp.val[3] = vcombine_p8 (b.val[3], vcreate_p8 (((uint64_t) 0))); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[3], 3); builtin_aarch64_st4_lanev8qi ((builtin_aarch64_simd_qi *) ptr, o, c); }

DEF: void vst4_lane_p16 (poly16_t *ptr, poly16x4x4_t b, const int c) { builtin_aarch64_simd_xi o; poly16x8x4_t temp; temp.val[0] = vcombine_p16 (b.val[0], vcreate_p16 (((uint64_t) 0))); temp.val[1] = vcombine_p16 (b.val[1], vcreate_p16 (((uint64_t) 0))); temp.val[2] = vcombine_p16 (b.val[2], vcreate_p16 (((uint64_t) 0))); temp.val[3] = vcombine_p16 (b.val[3], vcreate_p16 (((uint64_t) 0))); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[3], 3); builtin_aarch64_st4_lanev4hi ((builtin_aarch64_simd_hi *) ptr, o, c); }

DEF: void vst4_lane_s8 (int8_t *ptr, int8x8x4_t b, const int c) { builtin_aarch64_simd_xi o; int8x16x4_t temp; temp.val[0] = vcombine_s8 (b.val[0], vcreate_s8 (((uint64_t) 0))); temp.val[1] = vcombine_s8 (b.val[1], vcreate_s8 (((uint64_t) 0))); temp.val[2] = vcombine_s8 (b.val[2], vcreate_s8 (((uint64_t) 0))); temp.val[3] = vcombine_s8 (b.val[3], vcreate_s8 (((uint64_t) 0))); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[3], 3); builtin_aarch64_st4_lanev8qi ((builtin_aarch64_simd_qi *) ptr, o, c); }

DEF: void vst4_lane_s16 (int16_t *ptr, int16x4x4_t b, const int c) { builtin_aarch64_simd_xi o; int16x8x4_t temp; temp.val[0] = vcombine_s16 (b.val[0], vcreate_s16 (((uint64_t) 0))); temp.val[1] = vcombine_s16 (b.val[1], vcreate_s16 (((uint64_t) 0))); temp.val[2] = vcombine_s16 (b.val[2], vcreate_s16 (((uint64_t) 0))); temp.val[3] = vcombine_s16 (b.val[3], vcreate_s16 (((uint64_t) 0))); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[3], 3); builtin_aarch64_st4_lanev4hi ((builtin_aarch64_simd_hi *) ptr, o, c); }

DEF: void vst4_lane_s32 (int32_t *ptr, int32x2x4_t b, const int c) { builtin_aarch64_simd_xi o; int32x4x4_t temp; temp.val[0] = vcombine_s32 (b.val[0], vcreate_s32 (((uint64_t) 0))); temp.val[1] = vcombine_s32 (b.val[1], vcreate_s32 (((uint64_t) 0))); temp.val[2] = vcombine_s32 (b.val[2], vcreate_s32 (((uint64_t) 0))); temp.val[3] = vcombine_s32 (b.val[3], vcreate_s32 (((uint64_t) 0))); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[3], 3); builtin_aarch64_st4_lanev2si ((builtin_aarch64_simd_si *) ptr, o, c); }

DEF: void vst4_lane_s64 (int64_t *ptr, int64x1x4_t b, const int c) { builtin_aarch64_simd_xi o; int64x2x4_t temp; temp.val[0] = vcombine_s64 (b.val[0], vcreate_s64 (((uint64_t) 0))); temp.val[1] = vcombine_s64 (b.val[1], vcreate_s64 (((uint64_t) 0))); temp.val[2] = vcombine_s64 (b.val[2], vcreate_s64 (((uint64_t) 0))); temp.val[3] = vcombine_s64 (b.val[3], vcreate_s64 (((uint64_t) 0))); o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[3], 3); builtin_aarch64_st4_lanedi ((builtin_aarch64_simd_di *) ptr, o, c); }

DEF: void vst4_lane_u8 (uint8_t *ptr, uint8x8x4_t b, const int c) { builtin_aarch64_simd_xi o; uint8x16x4_t temp; temp.val[0] = vcombine_u8 (b.val[0], vcreate_u8 (((uint64_t) 0))); temp.val[1] = vcombine_u8 (b.val[1], vcreate_u8 (((uint64_t) 0))); temp.val[2] = vcombine_u8 (b.val[2], vcreate_u8 (((uint64_t) 0))); temp.val[3] = vcombine_u8 (b.val[3], vcreate_u8 (((uint64_t) 0))); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[3], 3); builtin_aarch64_st4_lanev8qi ((builtin_aarch64_simd_qi *) ptr, o, c); }

DEF: void vst4_lane_u16 (uint16_t *ptr, uint16x4x4_t b, const int c) { builtin_aarch64_simd_xi o; uint16x8x4_t temp; temp.val[0] = vcombine_u16 (b.val[0], vcreate_u16 (((uint64_t) 0))); temp.val[1] = vcombine_u16 (b.val[1], vcreate_u16 (((uint64_t) 0))); temp.val[2] = vcombine_u16 (b.val[2], vcreate_u16 (((uint64_t) 0))); temp.val[3] = vcombine_u16 (b.val[3], vcreate_u16 (((uint64_t) 0))); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[3], 3); builtin_aarch64_st4_lanev4hi ((builtin_aarch64_simd_hi *) ptr, o, c); }

DEF: void vst4_lane_u32 (uint32_t *ptr, uint32x2x4_t b, const int c) { builtin_aarch64_simd_xi o; uint32x4x4_t temp; temp.val[0] = vcombine_u32 (b.val[0], vcreate_u32 (((uint64_t) 0))); temp.val[1] = vcombine_u32 (b.val[1], vcreate_u32 (((uint64_t) 0))); temp.val[2] = vcombine_u32 (b.val[2], vcreate_u32 (((uint64_t) 0))); temp.val[3] = vcombine_u32 (b.val[3], vcreate_u32 (((uint64_t) 0))); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[3], 3); builtin_aarch64_st4_lanev2si ((builtin_aarch64_simd_si *) ptr, o, c); }

DEF: void vst4_lane_u64 (uint64_t *ptr, uint64x1x4_t b, const int c) { builtin_aarch64_simd_xi o; uint64x2x4_t temp; temp.val[0] = vcombine_u64 (b.val[0], vcreate_u64 (((uint64_t) 0))); temp.val[1] = vcombine_u64 (b.val[1], vcreate_u64 (((uint64_t) 0))); temp.val[2] = vcombine_u64 (b.val[2], vcreate_u64 (((uint64_t) 0))); temp.val[3] = vcombine_u64 (b.val[3], vcreate_u64 (((uint64_t) 0))); o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[3], 3); builtin_aarch64_st4_lanedi ((builtin_aarch64_simd_di *) ptr, o, c); }
# 10176 "arm_neon.h"
DEF: void vst4q_lane_f32 (float32_t *ptr, float32x4x4_t b, const int c) { union { float32x4x4_t i; builtin_aarch64_simd_xi o; } temp = { b }; builtin_aarch64_st4_lanev4sf ((builtin_aarch64_simd_sf *) ptr, temp.o, c); }
DEF: void vst4q_lane_f64 (float64_t *ptr, float64x2x4_t b, const int c) { union { float64x2x4_t i; builtin_aarch64_simd_xi o; } temp = { b }; builtin_aarch64_st4_lanev2df ((builtin_aarch64_simd_df *) ptr, temp.o, c); }
DEF: void vst4q_lane_p8 (poly8_t *ptr, poly8x16x4_t b, const int c) { union { poly8x16x4_t i; builtin_aarch64_simd_xi o; } temp = { b }; builtin_aarch64_st4_lanev16qi ((builtin_aarch64_simd_qi *) ptr, temp.o, c); }
DEF: void vst4q_lane_p16 (poly16_t *ptr, poly16x8x4_t b, const int c) { union { poly16x8x4_t i; builtin_aarch64_simd_xi o; } temp = { b }; builtin_aarch64_st4_lanev8hi ((builtin_aarch64_simd_hi *) ptr, temp.o, c); }
DEF: void vst4q_lane_s8 (int8_t *ptr, int8x16x4_t b, const int c) { union { int8x16x4_t i; builtin_aarch64_simd_xi o; } temp = { b }; builtin_aarch64_st4_lanev16qi ((builtin_aarch64_simd_qi *) ptr, temp.o, c); }
DEF: void vst4q_lane_s16 (int16_t *ptr, int16x8x4_t b, const int c) { union { int16x8x4_t i; builtin_aarch64_simd_xi o; } temp = { b }; builtin_aarch64_st4_lanev8hi ((builtin_aarch64_simd_hi *) ptr, temp.o, c); }
DEF: void vst4q_lane_s32 (int32_t *ptr, int32x4x4_t b, const int c) { union { int32x4x4_t i; builtin_aarch64_simd_xi o; } temp = { b }; builtin_aarch64_st4_lanev4si ((builtin_aarch64_simd_si *) ptr, temp.o, c); }
DEF: void vst4q_lane_s64 (int64_t *ptr, int64x2x4_t b, const int c) { union { int64x2x4_t i; builtin_aarch64_simd_xi o; } temp = { b }; builtin_aarch64_st4_lanev2di ((builtin_aarch64_simd_di *) ptr, temp.o, c); }
DEF: void vst4q_lane_u8 (uint8_t *ptr, uint8x16x4_t b, const int c) { union { uint8x16x4_t i; builtin_aarch64_simd_xi o; } temp = { b }; builtin_aarch64_st4_lanev16qi ((builtin_aarch64_simd_qi *) ptr, temp.o, c); }
DEF: void vst4q_lane_u16 (uint16_t *ptr, uint16x8x4_t b, const int c) { union { uint16x8x4_t i; builtin_aarch64_simd_xi o; } temp = { b }; builtin_aarch64_st4_lanev8hi ((builtin_aarch64_simd_hi *) ptr, temp.o, c); }
DEF: void vst4q_lane_u32 (uint32_t *ptr, uint32x4x4_t b, const int c) { union { uint32x4x4_t i; builtin_aarch64_simd_xi o; } temp = { b }; builtin_aarch64_st4_lanev4si ((builtin_aarch64_simd_si *) ptr, temp.o, c); }
DEF: void vst4q_lane_u64 (uint64_t *ptr, uint64x2x4_t b, const int c) { union { uint64x2x4_t i; builtin_aarch64_simd_xi o; } temp = { b }; builtin_aarch64_st4_lanev2di ((builtin_aarch64_simd_di *) ptr, temp.o, c); }

DEF: int64
vaddlv_s32 (int32x2_t a)
{
  int64_t result;
  asm ("saddlp %0.1d, %1.2s" : "=w"(result) : "w"(a) : );
  return result;
}

DEF: uint64
vaddlv_u32 (uint32x2_t a)
{
  uint64_t result;
  asm ("uaddlp %0.1d, %1.2s" : "=w"(result) : "w"(a) : );
  return result;
}

DEF: int16x4
vqdmulh_laneq_s16 (int16x4_t a, int16x8_t b, const int c)
{
  return builtin_aarch64_sqdmulh_laneqv4hi (a, b, c);
}

DEF: int32x2
vqdmulh_laneq_s32 (int32x2_t a, int32x4_t b, const int c)
{
  return builtin_aarch64_sqdmulh_laneqv2si (a, b, c);
}

DEF: int16x8
vqdmulhq_laneq_s16 (int16x8_t a, int16x8_t b, const int c)
{
  return builtin_aarch64_sqdmulh_laneqv8hi (a, b, c);
}

DEF: int32x4
vqdmulhq_laneq_s32 (int32x4_t a, int32x4_t b, const int c)
{
  return builtin_aarch64_sqdmulh_laneqv4si (a, b, c);
}

DEF: int16x4
vqrdmulh_laneq_s16 (int16x4_t a, int16x8_t b, const int c)
{
  return builtin_aarch64_sqrdmulh_laneqv4hi (a, b, c);
}

DEF: int32x2
vqrdmulh_laneq_s32 (int32x2_t a, int32x4_t b, const int c)
{
  return builtin_aarch64_sqrdmulh_laneqv2si (a, b, c);
}

DEF: int16x8
vqrdmulhq_laneq_s16 (int16x8_t a, int16x8_t b, const int c)
{
  return builtin_aarch64_sqrdmulh_laneqv8hi (a, b, c);
}

DEF: int32x4
vqrdmulhq_laneq_s32 (int32x4_t a, int32x4_t b, const int c)
{
  return builtin_aarch64_sqrdmulh_laneqv4si (a, b, c);
}

DEF: poly8x8
vqtbl1_p8 (poly8x16_t a, uint8x8_t b)
{
  poly8x8_t result;
  asm ("tbl %0.8b, {%1.16b}, %2.8b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int8x8
vqtbl1_s8 (int8x16_t a, uint8x8_t b)
{
  int8x8_t result;
  asm ("tbl %0.8b, {%1.16b}, %2.8b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint8x8
vqtbl1_u8 (uint8x16_t a, uint8x8_t b)
{
  uint8x8_t result;
  asm ("tbl %0.8b, {%1.16b}, %2.8b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: poly8x16
vqtbl1q_p8 (poly8x16_t a, uint8x16_t b)
{
  poly8x16_t result;
  asm ("tbl %0.16b, {%1.16b}, %2.16b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int8x16
vqtbl1q_s8 (int8x16_t a, uint8x16_t b)
{
  int8x16_t result;
  asm ("tbl %0.16b, {%1.16b}, %2.16b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: uint8x16
vqtbl1q_u8 (uint8x16_t a, uint8x16_t b)
{
  uint8x16_t result;
  asm ("tbl %0.16b, {%1.16b}, %2.16b"
           : "=w"(result)
           : "w"(a), "w"(b)
           : );
  return result;
}

DEF: int8x8
vqtbl2_s8 (int8x16x2_t tab, uint8x8_t idx)
{
  int8x8_t result;
  asm ("ld1 {v16.16b, v17.16b}, %1\n\t"
    "tbl %0.8b, {v16.16b, v17.16b}, %2.8b\n\t"
    :"=w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17");
  return result;
}

DEF: uint8x8
vqtbl2_u8 (uint8x16x2_t tab, uint8x8_t idx)
{
  uint8x8_t result;
  asm ("ld1 {v16.16b, v17.16b}, %1\n\t"
    "tbl %0.8b, {v16.16b, v17.16b}, %2.8b\n\t"
    :"=w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17");
  return result;
}

DEF: poly8x8
vqtbl2_p8 (poly8x16x2_t tab, uint8x8_t idx)
{
  poly8x8_t result;
  asm ("ld1 {v16.16b, v17.16b}, %1\n\t"
    "tbl %0.8b, {v16.16b, v17.16b}, %2.8b\n\t"
    :"=w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17");
  return result;
}

DEF: int8x16
vqtbl2q_s8 (int8x16x2_t tab, uint8x16_t idx)
{
  int8x16_t result;
  asm ("ld1 {v16.16b, v17.16b}, %1\n\t"
    "tbl %0.16b, {v16.16b, v17.16b}, %2.16b\n\t"
    :"=w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17");
  return result;
}

DEF: uint8x16
vqtbl2q_u8 (uint8x16x2_t tab, uint8x16_t idx)
{
  uint8x16_t result;
  asm ("ld1 {v16.16b, v17.16b}, %1\n\t"
    "tbl %0.16b, {v16.16b, v17.16b}, %2.16b\n\t"
    :"=w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17");
  return result;
}

DEF: poly8x16
vqtbl2q_p8 (poly8x16x2_t tab, uint8x16_t idx)
{
  poly8x16_t result;
  asm ("ld1 {v16.16b, v17.16b}, %1\n\t"
    "tbl %0.16b, {v16.16b, v17.16b}, %2.16b\n\t"
    :"=w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17");
  return result;
}

DEF: int8x8
vqtbl3_s8 (int8x16x3_t tab, uint8x8_t idx)
{
  int8x8_t result;
  asm ("ld1 {v16.16b - v18.16b}, %1\n\t"
    "tbl %0.8b, {v16.16b - v18.16b}, %2.8b\n\t"
    :"=w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18");
  return result;
}

DEF: uint8x8
vqtbl3_u8 (uint8x16x3_t tab, uint8x8_t idx)
{
  uint8x8_t result;
  asm ("ld1 {v16.16b - v18.16b}, %1\n\t"
    "tbl %0.8b, {v16.16b - v18.16b}, %2.8b\n\t"
    :"=w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18");
  return result;
}

DEF: poly8x8
vqtbl3_p8 (poly8x16x3_t tab, uint8x8_t idx)
{
  poly8x8_t result;
  asm ("ld1 {v16.16b - v18.16b}, %1\n\t"
    "tbl %0.8b, {v16.16b - v18.16b}, %2.8b\n\t"
    :"=w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18");
  return result;
}

DEF: int8x16
vqtbl3q_s8 (int8x16x3_t tab, uint8x16_t idx)
{
  int8x16_t result;
  asm ("ld1 {v16.16b - v18.16b}, %1\n\t"
    "tbl %0.16b, {v16.16b - v18.16b}, %2.16b\n\t"
    :"=w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18");
  return result;
}

DEF: uint8x16
vqtbl3q_u8 (uint8x16x3_t tab, uint8x16_t idx)
{
  uint8x16_t result;
  asm ("ld1 {v16.16b - v18.16b}, %1\n\t"
    "tbl %0.16b, {v16.16b - v18.16b}, %2.16b\n\t"
    :"=w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18");
  return result;
}

DEF: poly8x16
vqtbl3q_p8 (poly8x16x3_t tab, uint8x16_t idx)
{
  poly8x16_t result;
  asm ("ld1 {v16.16b - v18.16b}, %1\n\t"
    "tbl %0.16b, {v16.16b - v18.16b}, %2.16b\n\t"
    :"=w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18");
  return result;
}

DEF: int8x8
vqtbl4_s8 (int8x16x4_t tab, uint8x8_t idx)
{
  int8x8_t result;
  asm ("ld1 {v16.16b - v19.16b}, %1\n\t"
    "tbl %0.8b, {v16.16b - v19.16b}, %2.8b\n\t"
    :"=w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18", "v19");
  return result;
}

DEF: uint8x8
vqtbl4_u8 (uint8x16x4_t tab, uint8x8_t idx)
{
  uint8x8_t result;
  asm ("ld1 {v16.16b - v19.16b}, %1\n\t"
    "tbl %0.8b, {v16.16b - v19.16b}, %2.8b\n\t"
    :"=w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18", "v19");
  return result;
}

DEF: poly8x8
vqtbl4_p8 (poly8x16x4_t tab, uint8x8_t idx)
{
  poly8x8_t result;
  asm ("ld1 {v16.16b - v19.16b}, %1\n\t"
    "tbl %0.8b, {v16.16b - v19.16b}, %2.8b\n\t"
    :"=w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18", "v19");
  return result;
}
DEF: int8x16
vqtbl4q_s8 (int8x16x4_t tab, uint8x16_t idx)
{
  int8x16_t result;
  asm ("ld1 {v16.16b - v19.16b}, %1\n\t"
    "tbl %0.16b, {v16.16b - v19.16b}, %2.16b\n\t"
    :"=w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18", "v19");
  return result;
}

DEF: uint8x16
vqtbl4q_u8 (uint8x16x4_t tab, uint8x16_t idx)
{
  uint8x16_t result;
  asm ("ld1 {v16.16b - v19.16b}, %1\n\t"
    "tbl %0.16b, {v16.16b - v19.16b}, %2.16b\n\t"
    :"=w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18", "v19");
  return result;
}

DEF: poly8x16
vqtbl4q_p8 (poly8x16x4_t tab, uint8x16_t idx)
{
  poly8x16_t result;
  asm ("ld1 {v16.16b - v19.16b}, %1\n\t"
    "tbl %0.16b, {v16.16b - v19.16b}, %2.16b\n\t"
    :"=w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18", "v19");
  return result;
}
DEF: int8x8
vqtbx1_s8 (int8x8_t r, int8x16_t tab, uint8x8_t idx)
{
  int8x8_t result = r;
  asm ("tbx %0.8b,{%1.16b},%2.8b"
           : "+w"(result)
           : "w"(tab), "w"(idx)
           : );
  return result;
}

DEF: uint8x8
vqtbx1_u8 (uint8x8_t r, uint8x16_t tab, uint8x8_t idx)
{
  uint8x8_t result = r;
  asm ("tbx %0.8b,{%1.16b},%2.8b"
           : "+w"(result)
           : "w"(tab), "w"(idx)
           : );
  return result;
}

DEF: poly8x8
vqtbx1_p8 (poly8x8_t r, poly8x16_t tab, uint8x8_t idx)
{
  poly8x8_t result = r;
  asm ("tbx %0.8b,{%1.16b},%2.8b"
           : "+w"(result)
           : "w"(tab), "w"(idx)
           : );
  return result;
}

DEF: int8x16
vqtbx1q_s8 (int8x16_t r, int8x16_t tab, uint8x16_t idx)
{
  int8x16_t result = r;
  asm ("tbx %0.16b,{%1.16b},%2.16b"
           : "+w"(result)
           : "w"(tab), "w"(idx)
           : );
  return result;
}

DEF: uint8x16
vqtbx1q_u8 (uint8x16_t r, uint8x16_t tab, uint8x16_t idx)
{
  uint8x16_t result = r;
  asm ("tbx %0.16b,{%1.16b},%2.16b"
           : "+w"(result)
           : "w"(tab), "w"(idx)
           : );
  return result;
}

DEF: poly8x16
vqtbx1q_p8 (poly8x16_t r, poly8x16_t tab, uint8x16_t idx)
{
  poly8x16_t result = r;
  asm ("tbx %0.16b,{%1.16b},%2.16b"
           : "+w"(result)
           : "w"(tab), "w"(idx)
           : );
  return result;
}

DEF: int8x8
vqtbx2_s8 (int8x8_t r, int8x16x2_t tab, uint8x8_t idx)
{
  int8x8_t result = r;
  asm ("ld1 {v16.16b, v17.16b}, %1\n\t"
    "tbx %0.8b, {v16.16b, v17.16b}, %2.8b\n\t"
    :"+w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17");
  return result;
}

DEF: uint8x8
vqtbx2_u8 (uint8x8_t r, uint8x16x2_t tab, uint8x8_t idx)
{
  uint8x8_t result = r;
  asm ("ld1 {v16.16b, v17.16b}, %1\n\t"
    "tbx %0.8b, {v16.16b, v17.16b}, %2.8b\n\t"
    :"+w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17");
  return result;
}

DEF: poly8x8
vqtbx2_p8 (poly8x8_t r, poly8x16x2_t tab, uint8x8_t idx)
{
  poly8x8_t result = r;
  asm ("ld1 {v16.16b, v17.16b}, %1\n\t"
    "tbx %0.8b, {v16.16b, v17.16b}, %2.8b\n\t"
    :"+w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17");
  return result;
}
DEF: int8x16
vqtbx2q_s8 (int8x16_t r, int8x16x2_t tab, uint8x16_t idx)
{
  int8x16_t result = r;
  asm ("ld1 {v16.16b, v17.16b}, %1\n\t"
    "tbx %0.16b, {v16.16b, v17.16b}, %2.16b\n\t"
    :"+w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17");
  return result;
}

DEF: uint8x16
vqtbx2q_u8 (uint8x16_t r, uint8x16x2_t tab, uint8x16_t idx)
{
  uint8x16_t result = r;
  asm ("ld1 {v16.16b, v17.16b}, %1\n\t"
    "tbx %0.16b, {v16.16b, v17.16b}, %2.16b\n\t"
    :"+w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17");
  return result;
}

DEF: poly8x16
vqtbx2q_p8 (poly8x16_t r, poly8x16x2_t tab, uint8x16_t idx)
{
  poly8x16_t result = r;
  asm ("ld1 {v16.16b, v17.16b}, %1\n\t"
    "tbx %0.16b, {v16.16b, v17.16b}, %2.16b\n\t"
    :"+w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17");
  return result;
}
DEF: int8x8
vqtbx3_s8 (int8x8_t r, int8x16x3_t tab, uint8x8_t idx)
{
  int8x8_t result = r;
  asm ("ld1 {v16.16b - v18.16b}, %1\n\t"
    "tbx %0.8b, {v16.16b - v18.16b}, %2.8b\n\t"
    :"+w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18");
  return result;
}

DEF: uint8x8
vqtbx3_u8 (uint8x8_t r, uint8x16x3_t tab, uint8x8_t idx)
{
  uint8x8_t result = r;
  asm ("ld1 {v16.16b - v18.16b}, %1\n\t"
    "tbx %0.8b, {v16.16b - v18.16b}, %2.8b\n\t"
    :"+w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18");
  return result;
}

DEF: poly8x8
vqtbx3_p8 (poly8x8_t r, poly8x16x3_t tab, uint8x8_t idx)
{
  poly8x8_t result = r;
  asm ("ld1 {v16.16b - v18.16b}, %1\n\t"
    "tbx %0.8b, {v16.16b - v18.16b}, %2.8b\n\t"
    :"+w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18");
  return result;
}
DEF: int8x16
vqtbx3q_s8 (int8x16_t r, int8x16x3_t tab, uint8x16_t idx)
{
  int8x16_t result = r;
  asm ("ld1 {v16.16b - v18.16b}, %1\n\t"
    "tbx %0.16b, {v16.16b - v18.16b}, %2.16b\n\t"
    :"+w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18");
  return result;
}

DEF: uint8x16
vqtbx3q_u8 (uint8x16_t r, uint8x16x3_t tab, uint8x16_t idx)
{
  uint8x16_t result = r;
  asm ("ld1 {v16.16b - v18.16b}, %1\n\t"
    "tbx %0.16b, {v16.16b - v18.16b}, %2.16b\n\t"
    :"+w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18");
  return result;
}

DEF: poly8x16
vqtbx3q_p8 (poly8x16_t r, poly8x16x3_t tab, uint8x16_t idx)
{
  poly8x16_t result = r;
  asm ("ld1 {v16.16b - v18.16b}, %1\n\t"
    "tbx %0.16b, {v16.16b - v18.16b}, %2.16b\n\t"
    :"+w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18");
  return result;
}
DEF: int8x8
vqtbx4_s8 (int8x8_t r, int8x16x4_t tab, uint8x8_t idx)
{
  int8x8_t result = r;
  asm ("ld1 {v16.16b - v19.16b}, %1\n\t"
    "tbx %0.8b, {v16.16b - v19.16b}, %2.8b\n\t"
    :"+w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18", "v19");
  return result;
}

DEF: uint8x8
vqtbx4_u8 (uint8x8_t r, uint8x16x4_t tab, uint8x8_t idx)
{
  uint8x8_t result = r;
  asm ("ld1 {v16.16b - v19.16b}, %1\n\t"
    "tbx %0.8b, {v16.16b - v19.16b}, %2.8b\n\t"
    :"+w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18", "v19");
  return result;
}

DEF: poly8x8
vqtbx4_p8 (poly8x8_t r, poly8x16x4_t tab, uint8x8_t idx)
{
  poly8x8_t result = r;
  asm ("ld1 {v16.16b - v19.16b}, %1\n\t"
    "tbx %0.8b, {v16.16b - v19.16b}, %2.8b\n\t"
    :"+w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18", "v19");
  return result;
}
DEF: int8x16
vqtbx4q_s8 (int8x16_t r, int8x16x4_t tab, uint8x16_t idx)
{
  int8x16_t result = r;
  asm ("ld1 {v16.16b - v19.16b}, %1\n\t"
    "tbx %0.16b, {v16.16b - v19.16b}, %2.16b\n\t"
    :"+w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18", "v19");
  return result;
}

DEF: uint8x16
vqtbx4q_u8 (uint8x16_t r, uint8x16x4_t tab, uint8x16_t idx)
{
  uint8x16_t result = r;
  asm ("ld1 {v16.16b - v19.16b}, %1\n\t"
    "tbx %0.16b, {v16.16b - v19.16b}, %2.16b\n\t"
    :"+w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18", "v19");
  return result;
}

DEF: poly8x16
vqtbx4q_p8 (poly8x16_t r, poly8x16x4_t tab, uint8x16_t idx)
{
  poly8x16_t result = r;
  asm ("ld1 {v16.16b - v19.16b}, %1\n\t"
    "tbx %0.16b, {v16.16b - v19.16b}, %2.16b\n\t"
    :"+w"(result)
    :"Q"(tab),"w"(idx)
    :"memory", "v16", "v17", "v18", "v19");
  return result;
}

DEF: int8x8
vtbl1_s8 (int8x8_t tab, int8x8_t idx)
{
  int8x8_t result;
  int8x16_t temp = vcombine_s8 (tab, vcreate_s8 (((uint64_t) 0x0)));
  asm ("tbl %0.8b, {%1.16b}, %2.8b"
           : "=w"(result)
           : "w"(temp), "w"(idx)
           : );
  return result;
}

DEF: uint8x8
vtbl1_u8 (uint8x8_t tab, uint8x8_t idx)
{
  uint8x8_t result;
  uint8x16_t temp = vcombine_u8 (tab, vcreate_u8 (((uint64_t) 0x0)));
  asm ("tbl %0.8b, {%1.16b}, %2.8b"
           : "=w"(result)
           : "w"(temp), "w"(idx)
           : );
  return result;
}

DEF: poly8x8
vtbl1_p8 (poly8x8_t tab, uint8x8_t idx)
{
  poly8x8_t result;
  poly8x16_t temp = vcombine_p8 (tab, vcreate_p8 (((uint64_t) 0x0)));
  asm ("tbl %0.8b, {%1.16b}, %2.8b"
           : "=w"(result)
           : "w"(temp), "w"(idx)
           : );
  return result;
}

DEF: int8x8
vtbl2_s8 (int8x8x2_t tab, int8x8_t idx)
{
  int8x8_t result;
  int8x16_t temp = vcombine_s8 (tab.val[0], tab.val[1]);
  asm ("tbl %0.8b, {%1.16b}, %2.8b"
           : "=w"(result)
           : "w"(temp), "w"(idx)
           : );
  return result;
}

DEF: uint8x8
vtbl2_u8 (uint8x8x2_t tab, uint8x8_t idx)
{
  uint8x8_t result;
  uint8x16_t temp = vcombine_u8 (tab.val[0], tab.val[1]);
  asm ("tbl %0.8b, {%1.16b}, %2.8b"
           : "=w"(result)
           : "w"(temp), "w"(idx)
           : );
  return result;
}

DEF: poly8x8
vtbl2_p8 (poly8x8x2_t tab, uint8x8_t idx)
{
  poly8x8_t result;
  poly8x16_t temp = vcombine_p8 (tab.val[0], tab.val[1]);
  asm ("tbl %0.8b, {%1.16b}, %2.8b"
           : "=w"(result)
           : "w"(temp), "w"(idx)
           : );
  return result;
}

DEF: int8x8
vtbl3_s8 (int8x8x3_t tab, int8x8_t idx)
{
  int8x8_t result;
  int8x16x2_t temp;
  temp.val[0] = vcombine_s8 (tab.val[0], tab.val[1]);
  temp.val[1] = vcombine_s8 (tab.val[2], vcreate_s8 (((uint64_t) 0x0)));
  asm ("ld1 {v16.16b - v17.16b }, %1\n\t"
    "tbl %0.8b, {v16.16b - v17.16b}, %2.8b\n\t"
           : "=w"(result)
           : "Q"(temp), "w"(idx)
           : "v16", "v17", "memory");
  return result;
}

DEF: uint8x8
vtbl3_u8 (uint8x8x3_t tab, uint8x8_t idx)
{
  uint8x8_t result;
  uint8x16x2_t temp;
  temp.val[0] = vcombine_u8 (tab.val[0], tab.val[1]);
  temp.val[1] = vcombine_u8 (tab.val[2], vcreate_u8 (((uint64_t) 0x0)));
  asm ("ld1 {v16.16b - v17.16b }, %1\n\t"
    "tbl %0.8b, {v16.16b - v17.16b}, %2.8b\n\t"
           : "=w"(result)
           : "Q"(temp), "w"(idx)
           : "v16", "v17", "memory");
  return result;
}

DEF: poly8x8
vtbl3_p8 (poly8x8x3_t tab, uint8x8_t idx)
{
  poly8x8_t result;
  poly8x16x2_t temp;
  temp.val[0] = vcombine_p8 (tab.val[0], tab.val[1]);
  temp.val[1] = vcombine_p8 (tab.val[2], vcreate_p8 (((uint64_t) 0x0)));
  asm ("ld1 {v16.16b - v17.16b }, %1\n\t"
    "tbl %0.8b, {v16.16b - v17.16b}, %2.8b\n\t"
           : "=w"(result)
           : "Q"(temp), "w"(idx)
           : "v16", "v17", "memory");
  return result;
}

DEF: int8x8
vtbl4_s8 (int8x8x4_t tab, int8x8_t idx)
{
  int8x8_t result;
  int8x16x2_t temp;
  temp.val[0] = vcombine_s8 (tab.val[0], tab.val[1]);
  temp.val[1] = vcombine_s8 (tab.val[2], tab.val[3]);
  asm ("ld1 {v16.16b - v17.16b }, %1\n\t"
    "tbl %0.8b, {v16.16b - v17.16b}, %2.8b\n\t"
           : "=w"(result)
           : "Q"(temp), "w"(idx)
           : "v16", "v17", "memory");
  return result;
}

DEF: uint8x8
vtbl4_u8 (uint8x8x4_t tab, uint8x8_t idx)
{
  uint8x8_t result;
  uint8x16x2_t temp;
  temp.val[0] = vcombine_u8 (tab.val[0], tab.val[1]);
  temp.val[1] = vcombine_u8 (tab.val[2], tab.val[3]);
  asm ("ld1 {v16.16b - v17.16b }, %1\n\t"
    "tbl %0.8b, {v16.16b - v17.16b}, %2.8b\n\t"
           : "=w"(result)
           : "Q"(temp), "w"(idx)
           : "v16", "v17", "memory");
  return result;
}

DEF: poly8x8
vtbl4_p8 (poly8x8x4_t tab, uint8x8_t idx)
{
  poly8x8_t result;
  poly8x16x2_t temp;
  temp.val[0] = vcombine_p8 (tab.val[0], tab.val[1]);
  temp.val[1] = vcombine_p8 (tab.val[2], tab.val[3]);
  asm ("ld1 {v16.16b - v17.16b }, %1\n\t"
    "tbl %0.8b, {v16.16b - v17.16b}, %2.8b\n\t"
           : "=w"(result)
           : "Q"(temp), "w"(idx)
           : "v16", "v17", "memory");
  return result;
}

DEF: int8x8
vtbx2_s8 (int8x8_t r, int8x8x2_t tab, int8x8_t idx)
{
  int8x8_t result = r;
  int8x16_t temp = vcombine_s8 (tab.val[0], tab.val[1]);
  asm ("tbx %0.8b, {%1.16b}, %2.8b"
           : "+w"(result)
           : "w"(temp), "w"(idx)
           : );
  return result;
}

DEF: uint8x8
vtbx2_u8 (uint8x8_t r, uint8x8x2_t tab, uint8x8_t idx)
{
  uint8x8_t result = r;
  uint8x16_t temp = vcombine_u8 (tab.val[0], tab.val[1]);
  asm ("tbx %0.8b, {%1.16b}, %2.8b"
           : "+w"(result)
           : "w"(temp), "w"(idx)
           : );
  return result;
}

DEF: poly8x8
vtbx2_p8 (poly8x8_t r, poly8x8x2_t tab, uint8x8_t idx)
{
  poly8x8_t result = r;
  poly8x16_t temp = vcombine_p8 (tab.val[0], tab.val[1]);
  asm ("tbx %0.8b, {%1.16b}, %2.8b"
           : "+w"(result)
           : "w"(temp), "w"(idx)
           : );
  return result;
}

DEF: int8x8
vtbx4_s8 (int8x8_t r, int8x8x4_t tab, int8x8_t idx)
{
  int8x8_t result = r;
  int8x16x2_t temp;
  temp.val[0] = vcombine_s8 (tab.val[0], tab.val[1]);
  temp.val[1] = vcombine_s8 (tab.val[2], tab.val[3]);
  asm ("ld1 {v16.16b - v17.16b }, %1\n\t"
    "tbx %0.8b, {v16.16b - v17.16b}, %2.8b\n\t"
           : "+w"(result)
           : "Q"(temp), "w"(idx)
           : "v16", "v17", "memory");
  return result;
}

DEF: uint8x8
vtbx4_u8 (uint8x8_t r, uint8x8x4_t tab, uint8x8_t idx)
{
  uint8x8_t result = r;
  uint8x16x2_t temp;
  temp.val[0] = vcombine_u8 (tab.val[0], tab.val[1]);
  temp.val[1] = vcombine_u8 (tab.val[2], tab.val[3]);
  asm ("ld1 {v16.16b - v17.16b }, %1\n\t"
    "tbx %0.8b, {v16.16b - v17.16b}, %2.8b\n\t"
           : "+w"(result)
           : "Q"(temp), "w"(idx)
           : "v16", "v17", "memory");
  return result;
}

DEF: poly8x8
vtbx4_p8 (poly8x8_t r, poly8x8x4_t tab, uint8x8_t idx)
{
  poly8x8_t result = r;
  poly8x16x2_t temp;
  temp.val[0] = vcombine_p8 (tab.val[0], tab.val[1]);
  temp.val[1] = vcombine_p8 (tab.val[2], tab.val[3]);
  asm ("ld1 {v16.16b - v17.16b }, %1\n\t"
    "tbx %0.8b, {v16.16b - v17.16b}, %2.8b\n\t"
           : "+w"(result)
           : "Q"(temp), "w"(idx)
           : "v16", "v17", "memory");
  return result;
}



DEF: float32x2
vabs_f32 (float32x2_t a)
{
  return builtin_aarch64_absv2sf (a);
}

DEF: float64x1
vabs_f64 (float64x1_t a)
{
  return (float64x1_t) {builtin_fabs (a[0])};
}

DEF: int8x8
vabs_s8 (int8x8_t a)
{
  return builtin_aarch64_absv8qi (a);
}

DEF: int16x4
vabs_s16 (int16x4_t a)
{
  return builtin_aarch64_absv4hi (a);
}

DEF: int32x2
vabs_s32 (int32x2_t a)
{
  return builtin_aarch64_absv2si (a);
}

DEF: int64x1
vabs_s64 (int64x1_t a)
{
  return (int64x1_t) {builtin_aarch64_absdi (a[0])};
}

DEF: float32x4
vabsq_f32 (float32x4_t a)
{
  return builtin_aarch64_absv4sf (a);
}

DEF: float64x2
vabsq_f64 (float64x2_t a)
{
  return builtin_aarch64_absv2df (a);
}

DEF: int8x16
vabsq_s8 (int8x16_t a)
{
  return builtin_aarch64_absv16qi (a);
}

DEF: int16x8
vabsq_s16 (int16x8_t a)
{
  return builtin_aarch64_absv8hi (a);
}

DEF: int32x4
vabsq_s32 (int32x4_t a)
{
  return builtin_aarch64_absv4si (a);
}

DEF: int64x2
vabsq_s64 (int64x2_t a)
{
  return builtin_aarch64_absv2di (a);
}

DEF: int64
vaddd_s64 (int64_t a, int64_t b)
{
  return a + b;
}

DEF: uint64
vaddd_u64 (uint64_t a, uint64_t b)
{
  return a + b;
}

DEF: int8
vaddv_s8 (int8x8_t a)
{
  return builtin_aarch64_reduc_plus_scal_v8qi (a);
}

DEF: int16
vaddv_s16 (int16x4_t a)
{
  return builtin_aarch64_reduc_plus_scal_v4hi (a);
}

DEF: int32
vaddv_s32 (int32x2_t a)
{
  return builtin_aarch64_reduc_plus_scal_v2si (a);
}

DEF: uint8
vaddv_u8 (uint8x8_t a)
{
  return (uint8_t) builtin_aarch64_reduc_plus_scal_v8qi ((int8x8_t) a);
}

DEF: uint16
vaddv_u16 (uint16x4_t a)
{
  return (uint16_t) builtin_aarch64_reduc_plus_scal_v4hi ((int16x4_t) a);
}

DEF: uint32
vaddv_u32 (uint32x2_t a)
{
  return (int32_t) builtin_aarch64_reduc_plus_scal_v2si ((int32x2_t) a);
}

DEF: int8
vaddvq_s8 (int8x16_t a)
{
  return builtin_aarch64_reduc_plus_scal_v16qi (a);
}

DEF: int16
vaddvq_s16 (int16x8_t a)
{
  return builtin_aarch64_reduc_plus_scal_v8hi (a);
}

DEF: int32
vaddvq_s32 (int32x4_t a)
{
  return builtin_aarch64_reduc_plus_scal_v4si (a);
}

DEF: int64
vaddvq_s64 (int64x2_t a)
{
  return builtin_aarch64_reduc_plus_scal_v2di (a);
}

DEF: uint8
vaddvq_u8 (uint8x16_t a)
{
  return (uint8_t) builtin_aarch64_reduc_plus_scal_v16qi ((int8x16_t) a);
}

DEF: uint16
vaddvq_u16 (uint16x8_t a)
{
  return (uint16_t) builtin_aarch64_reduc_plus_scal_v8hi ((int16x8_t) a);
}

DEF: uint32
vaddvq_u32 (uint32x4_t a)
{
  return (uint32_t) builtin_aarch64_reduc_plus_scal_v4si ((int32x4_t) a);
}

DEF: uint64
vaddvq_u64 (uint64x2_t a)
{
  return (uint64_t) builtin_aarch64_reduc_plus_scal_v2di ((int64x2_t) a);
}

DEF: float32
vaddv_f32 (float32x2_t a)
{
  return builtin_aarch64_reduc_plus_scal_v2sf (a);
}

DEF: float32
vaddvq_f32 (float32x4_t a)
{
  return builtin_aarch64_reduc_plus_scal_v4sf (a);
}

DEF: float64
vaddvq_f64 (float64x2_t a)
{
  return builtin_aarch64_reduc_plus_scal_v2df (a);
}

DEF: float32x2
vbsl_f32 (uint32x2_t a, float32x2_t b, float32x2_t c)
{
  return builtin_aarch64_simd_bslv2sf_suss (a, b, c);
}

DEF: float64x1
vbsl_f64 (uint64x1_t a, float64x1_t b, float64x1_t c)
{
  return (float64x1_t)
    { builtin_aarch64_simd_bsldf_suss (a[0], b[0], c[0]) };
}

DEF: poly8x8
vbsl_p8 (uint8x8_t a, poly8x8_t b, poly8x8_t c)
{
  return builtin_aarch64_simd_bslv8qi_pupp (a, b, c);
}

DEF: poly16x4
vbsl_p16 (uint16x4_t a, poly16x4_t b, poly16x4_t c)
{
  return builtin_aarch64_simd_bslv4hi_pupp (a, b, c);
}

DEF: int8x8
vbsl_s8 (uint8x8_t a, int8x8_t b, int8x8_t c)
{
  return builtin_aarch64_simd_bslv8qi_suss (a, b, c);
}

DEF: int16x4
vbsl_s16 (uint16x4_t a, int16x4_t b, int16x4_t c)
{
  return builtin_aarch64_simd_bslv4hi_suss (a, b, c);
}

DEF: int32x2
vbsl_s32 (uint32x2_t a, int32x2_t b, int32x2_t c)
{
  return builtin_aarch64_simd_bslv2si_suss (a, b, c);
}

DEF: int64x1
vbsl_s64 (uint64x1_t a, int64x1_t b, int64x1_t c)
{
  return (int64x1_t)
      {builtin_aarch64_simd_bsldi_suss (a[0], b[0], c[0])};
}

DEF: uint8x8
vbsl_u8 (uint8x8_t a, uint8x8_t b, uint8x8_t c)
{
  return builtin_aarch64_simd_bslv8qi_uuuu (a, b, c);
}

DEF: uint16x4
vbsl_u16 (uint16x4_t a, uint16x4_t b, uint16x4_t c)
{
  return builtin_aarch64_simd_bslv4hi_uuuu (a, b, c);
}

DEF: uint32x2
vbsl_u32 (uint32x2_t a, uint32x2_t b, uint32x2_t c)
{
  return builtin_aarch64_simd_bslv2si_uuuu (a, b, c);
}

DEF: uint64x1
vbsl_u64 (uint64x1_t a, uint64x1_t b, uint64x1_t c)
{
  return (uint64x1_t)
      {builtin_aarch64_simd_bsldi_uuuu (a[0], b[0], c[0])};
}

DEF: float32x4
vbslq_f32 (uint32x4_t a, float32x4_t b, float32x4_t c)
{
  return builtin_aarch64_simd_bslv4sf_suss (a, b, c);
}

DEF: float64x2
vbslq_f64 (uint64x2_t a, float64x2_t b, float64x2_t c)
{
  return builtin_aarch64_simd_bslv2df_suss (a, b, c);
}

DEF: poly8x16
vbslq_p8 (uint8x16_t a, poly8x16_t b, poly8x16_t c)
{
  return builtin_aarch64_simd_bslv16qi_pupp (a, b, c);
}

DEF: poly16x8
vbslq_p16 (uint16x8_t a, poly16x8_t b, poly16x8_t c)
{
  return builtin_aarch64_simd_bslv8hi_pupp (a, b, c);
}

DEF: int8x16
vbslq_s8 (uint8x16_t a, int8x16_t b, int8x16_t c)
{
  return builtin_aarch64_simd_bslv16qi_suss (a, b, c);
}

DEF: int16x8
vbslq_s16 (uint16x8_t a, int16x8_t b, int16x8_t c)
{
  return builtin_aarch64_simd_bslv8hi_suss (a, b, c);
}

DEF: int32x4
vbslq_s32 (uint32x4_t a, int32x4_t b, int32x4_t c)
{
  return builtin_aarch64_simd_bslv4si_suss (a, b, c);
}

DEF: int64x2
vbslq_s64 (uint64x2_t a, int64x2_t b, int64x2_t c)
{
  return builtin_aarch64_simd_bslv2di_suss (a, b, c);
}

DEF: uint8x16
vbslq_u8 (uint8x16_t a, uint8x16_t b, uint8x16_t c)
{
  return builtin_aarch64_simd_bslv16qi_uuuu (a, b, c);
}

DEF: uint16x8
vbslq_u16 (uint16x8_t a, uint16x8_t b, uint16x8_t c)
{
  return builtin_aarch64_simd_bslv8hi_uuuu (a, b, c);
}

DEF: uint32x4
vbslq_u32 (uint32x4_t a, uint32x4_t b, uint32x4_t c)
{
  return builtin_aarch64_simd_bslv4si_uuuu (a, b, c);
}

DEF: uint64x2
vbslq_u64 (uint64x2_t a, uint64x2_t b, uint64x2_t c)
{
  return builtin_aarch64_simd_bslv2di_uuuu (a, b, c);
}

#pragma GCC push_options
#pragma GCC target ("+nothing+crypto")
DEF: uint8x16
vaeseq_u8 (uint8x16_t data, uint8x16_t key)
{
  return builtin_aarch64_crypto_aesev16qi_uuu (data, key);
}

DEF: uint8x16
vaesdq_u8 (uint8x16_t data, uint8x16_t key)
{
  return builtin_aarch64_crypto_aesdv16qi_uuu (data, key);
}

DEF: uint8x16
vaesmcq_u8 (uint8x16_t data)
{
  return builtin_aarch64_crypto_aesmcv16qi_uu (data);
}

DEF: uint8x16
vaesimcq_u8 (uint8x16_t data)
{
  return builtin_aarch64_crypto_aesimcv16qi_uu (data);
}
#pragma GCC pop_options

DEF: uint64x1
vcage_f64 (float64x1_t a, float64x1_t b)
{
  return vabs_f64 (a) >= vabs_f64 (b);
}

DEF: uint32
vcages_f32 (float32_t a, float32_t b)
{
  return builtin_fabsf (a) >= builtin_fabsf (b) ? -1 : 0;
}

DEF: uint32x2
vcage_f32 (float32x2_t a, float32x2_t b)
{
  return vabs_f32 (a) >= vabs_f32 (b);
}

DEF: uint32x4
vcageq_f32 (float32x4_t a, float32x4_t b)
{
  return vabsq_f32 (a) >= vabsq_f32 (b);
}

DEF: uint64
vcaged_f64 (float64_t a, float64_t b)
{
  return builtin_fabs (a) >= builtin_fabs (b) ? -1 : 0;
}

DEF: uint64x2
vcageq_f64 (float64x2_t a, float64x2_t b)
{
  return vabsq_f64 (a) >= vabsq_f64 (b);
}

DEF: uint32
vcagts_f32 (float32_t a, float32_t b)
{
  return builtin_fabsf (a) > builtin_fabsf (b) ? -1 : 0;
}

DEF: uint32x2
vcagt_f32 (float32x2_t a, float32x2_t b)
{
  return vabs_f32 (a) > vabs_f32 (b);
}

DEF: uint64x1
vcagt_f64 (float64x1_t a, float64x1_t b)
{
  return vabs_f64 (a) > vabs_f64 (b);
}

DEF: uint32x4
vcagtq_f32 (float32x4_t a, float32x4_t b)
{
  return vabsq_f32 (a) > vabsq_f32 (b);
}

DEF: uint64
vcagtd_f64 (float64_t a, float64_t b)
{
  return builtin_fabs (a) > builtin_fabs (b) ? -1 : 0;
}

DEF: uint64x2
vcagtq_f64 (float64x2_t a, float64x2_t b)
{
  return vabsq_f64 (a) > vabsq_f64 (b);
}

DEF: uint32x2
vcale_f32 (float32x2_t a, float32x2_t b)
{
  return vabs_f32 (a) <= vabs_f32 (b);
}

DEF: uint64x1
vcale_f64 (float64x1_t a, float64x1_t b)
{
  return vabs_f64 (a) <= vabs_f64 (b);
}

DEF: uint64
vcaled_f64 (float64_t a, float64_t b)
{
  return builtin_fabs (a) <= builtin_fabs (b) ? -1 : 0;
}

DEF: uint32
vcales_f32 (float32_t a, float32_t b)
{
  return builtin_fabsf (a) <= builtin_fabsf (b) ? -1 : 0;
}

DEF: uint32x4
vcaleq_f32 (float32x4_t a, float32x4_t b)
{
  return vabsq_f32 (a) <= vabsq_f32 (b);
}

DEF: uint64x2
vcaleq_f64 (float64x2_t a, float64x2_t b)
{
  return vabsq_f64 (a) <= vabsq_f64 (b);
}

DEF: uint32x2
vcalt_f32 (float32x2_t a, float32x2_t b)
{
  return vabs_f32 (a) < vabs_f32 (b);
}

DEF: uint64x1
vcalt_f64 (float64x1_t a, float64x1_t b)
{
  return vabs_f64 (a) < vabs_f64 (b);
}

DEF: uint64
vcaltd_f64 (float64_t a, float64_t b)
{
  return builtin_fabs (a) < builtin_fabs (b) ? -1 : 0;
}

DEF: uint32x4
vcaltq_f32 (float32x4_t a, float32x4_t b)
{
  return vabsq_f32 (a) < vabsq_f32 (b);
}

DEF: uint64x2
vcaltq_f64 (float64x2_t a, float64x2_t b)
{
  return vabsq_f64 (a) < vabsq_f64 (b);
}

DEF: uint32
vcalts_f32 (float32_t a, float32_t b)
{
  return builtin_fabsf (a) < builtin_fabsf (b) ? -1 : 0;
}

DEF: uint32x2
vceq_f32 (float32x2_t a, float32x2_t b)
{
  return (uint32x2_t) (a == b);
}

DEF: uint64x1
vceq_f64 (float64x1_t a, float64x1_t b)
{
  return (uint64x1_t) (a == b);
}

DEF: uint8x8
vceq_p8 (poly8x8_t a, poly8x8_t b)
{
  return (uint8x8_t) (a == b);
}

DEF: uint8x8
vceq_s8 (int8x8_t a, int8x8_t b)
{
  return (uint8x8_t) (a == b);
}

DEF: uint16x4
vceq_s16 (int16x4_t a, int16x4_t b)
{
  return (uint16x4_t) (a == b);
}

DEF: uint32x2
vceq_s32 (int32x2_t a, int32x2_t b)
{
  return (uint32x2_t) (a == b);
}

DEF: uint64x1
vceq_s64 (int64x1_t a, int64x1_t b)
{
  return (uint64x1_t) (a == b);
}

DEF: uint8x8
vceq_u8 (uint8x8_t a, uint8x8_t b)
{
  return (a == b);
}

DEF: uint16x4
vceq_u16 (uint16x4_t a, uint16x4_t b)
{
  return (a == b);
}

DEF: uint32x2
vceq_u32 (uint32x2_t a, uint32x2_t b)
{
  return (a == b);
}

DEF: uint64x1
vceq_u64 (uint64x1_t a, uint64x1_t b)
{
  return (a == b);
}

DEF: uint32x4
vceqq_f32 (float32x4_t a, float32x4_t b)
{
  return (uint32x4_t) (a == b);
}

DEF: uint64x2
vceqq_f64 (float64x2_t a, float64x2_t b)
{
  return (uint64x2_t) (a == b);
}

DEF: uint8x16
vceqq_p8 (poly8x16_t a, poly8x16_t b)
{
  return (uint8x16_t) (a == b);
}

DEF: uint8x16
vceqq_s8 (int8x16_t a, int8x16_t b)
{
  return (uint8x16_t) (a == b);
}

DEF: uint16x8
vceqq_s16 (int16x8_t a, int16x8_t b)
{
  return (uint16x8_t) (a == b);
}

DEF: uint32x4
vceqq_s32 (int32x4_t a, int32x4_t b)
{
  return (uint32x4_t) (a == b);
}

DEF: uint64x2
vceqq_s64 (int64x2_t a, int64x2_t b)
{
  return (uint64x2_t) (a == b);
}

DEF: uint8x16
vceqq_u8 (uint8x16_t a, uint8x16_t b)
{
  return (a == b);
}

DEF: uint16x8
vceqq_u16 (uint16x8_t a, uint16x8_t b)
{
  return (a == b);
}

DEF: uint32x4
vceqq_u32 (uint32x4_t a, uint32x4_t b)
{
  return (a == b);
}

DEF: uint64x2
vceqq_u64 (uint64x2_t a, uint64x2_t b)
{
  return (a == b);
}

DEF: uint32
vceqs_f32 (float32_t a, float32_t b)
{
  return a == b ? -1 : 0;
}

DEF: uint64
vceqd_s64 (int64_t a, int64_t b)
{
  return a == b ? -1ll : 0ll;
}

DEF: uint64
vceqd_u64 (uint64_t a, uint64_t b)
{
  return a == b ? -1ll : 0ll;
}

DEF: uint64
vceqd_f64 (float64_t a, float64_t b)
{
  return a == b ? -1ll : 0ll;
}

DEF: uint32x2
vceqz_f32 (float32x2_t a)
{
  return (uint32x2_t) (a == 0.0f);
}

DEF: uint64x1
vceqz_f64 (float64x1_t a)
{
  return (uint64x1_t) (a == (float64x1_t) {0.0});
}

DEF: uint8x8
vceqz_p8 (poly8x8_t a)
{
  return (uint8x8_t) (a == 0);
}

DEF: uint8x8
vceqz_s8 (int8x8_t a)
{
  return (uint8x8_t) (a == 0);
}

DEF: uint16x4
vceqz_s16 (int16x4_t a)
{
  return (uint16x4_t) (a == 0);
}

DEF: uint32x2
vceqz_s32 (int32x2_t a)
{
  return (uint32x2_t) (a == 0);
}

DEF: uint64x1
vceqz_s64 (int64x1_t a)
{
  return (uint64x1_t) (a == ((int64_t) 0));
}

DEF: uint8x8
vceqz_u8 (uint8x8_t a)
{
  return (a == 0);
}

DEF: uint16x4
vceqz_u16 (uint16x4_t a)
{
  return (a == 0);
}

DEF: uint32x2
vceqz_u32 (uint32x2_t a)
{
  return (a == 0);
}

DEF: uint64x1
vceqz_u64 (uint64x1_t a)
{
  return (a == ((uint64_t) 0));
}

DEF: uint32x4
vceqzq_f32 (float32x4_t a)
{
  return (uint32x4_t) (a == 0.0f);
}

DEF: uint64x2
vceqzq_f64 (float64x2_t a)
{
  return (uint64x2_t) (a == 0.0f);
}

DEF: uint8x16
vceqzq_p8 (poly8x16_t a)
{
  return (uint8x16_t) (a == 0);
}

DEF: uint8x16
vceqzq_s8 (int8x16_t a)
{
  return (uint8x16_t) (a == 0);
}

DEF: uint16x8
vceqzq_s16 (int16x8_t a)
{
  return (uint16x8_t) (a == 0);
}

DEF: uint32x4
vceqzq_s32 (int32x4_t a)
{
  return (uint32x4_t) (a == 0);
}

DEF: uint64x2
vceqzq_s64 (int64x2_t a)
{
  return (uint64x2_t) (a == ((int64_t) 0));
}

DEF: uint8x16
vceqzq_u8 (uint8x16_t a)
{
  return (a == 0);
}

DEF: uint16x8
vceqzq_u16 (uint16x8_t a)
{
  return (a == 0);
}

DEF: uint32x4
vceqzq_u32 (uint32x4_t a)
{
  return (a == 0);
}

DEF: uint64x2
vceqzq_u64 (uint64x2_t a)
{
  return (a == ((uint64_t) 0));
}

DEF: uint32
vceqzs_f32 (float32_t a)
{
  return a == 0.0f ? -1 : 0;
}

DEF: uint64
vceqzd_s64 (int64_t a)
{
  return a == 0 ? -1ll : 0ll;
}

DEF: uint64
vceqzd_u64 (uint64_t a)
{
  return a == 0 ? -1ll : 0ll;
}

DEF: uint64
vceqzd_f64 (float64_t a)
{
  return a == 0.0 ? -1ll : 0ll;
}

DEF: uint32x2
vcge_f32 (float32x2_t a, float32x2_t b)
{
  return (uint32x2_t) (a >= b);
}

DEF: uint64x1
vcge_f64 (float64x1_t a, float64x1_t b)
{
  return (uint64x1_t) (a >= b);
}

DEF: uint8x8
vcge_s8 (int8x8_t a, int8x8_t b)
{
  return (uint8x8_t) (a >= b);
}

DEF: uint16x4
vcge_s16 (int16x4_t a, int16x4_t b)
{
  return (uint16x4_t) (a >= b);
}

DEF: uint32x2
vcge_s32 (int32x2_t a, int32x2_t b)
{
  return (uint32x2_t) (a >= b);
}

DEF: uint64x1
vcge_s64 (int64x1_t a, int64x1_t b)
{
  return (uint64x1_t) (a >= b);
}

DEF: uint8x8
vcge_u8 (uint8x8_t a, uint8x8_t b)
{
  return (a >= b);
}

DEF: uint16x4
vcge_u16 (uint16x4_t a, uint16x4_t b)
{
  return (a >= b);
}

DEF: uint32x2
vcge_u32 (uint32x2_t a, uint32x2_t b)
{
  return (a >= b);
}

DEF: uint64x1
vcge_u64 (uint64x1_t a, uint64x1_t b)
{
  return (a >= b);
}

DEF: uint32x4
vcgeq_f32 (float32x4_t a, float32x4_t b)
{
  return (uint32x4_t) (a >= b);
}

DEF: uint64x2
vcgeq_f64 (float64x2_t a, float64x2_t b)
{
  return (uint64x2_t) (a >= b);
}

DEF: uint8x16
vcgeq_s8 (int8x16_t a, int8x16_t b)
{
  return (uint8x16_t) (a >= b);
}

DEF: uint16x8
vcgeq_s16 (int16x8_t a, int16x8_t b)
{
  return (uint16x8_t) (a >= b);
}

DEF: uint32x4
vcgeq_s32 (int32x4_t a, int32x4_t b)
{
  return (uint32x4_t) (a >= b);
}

DEF: uint64x2
vcgeq_s64 (int64x2_t a, int64x2_t b)
{
  return (uint64x2_t) (a >= b);
}

DEF: uint8x16
vcgeq_u8 (uint8x16_t a, uint8x16_t b)
{
  return (a >= b);
}

DEF: uint16x8
vcgeq_u16 (uint16x8_t a, uint16x8_t b)
{
  return (a >= b);
}

DEF: uint32x4
vcgeq_u32 (uint32x4_t a, uint32x4_t b)
{
  return (a >= b);
}

DEF: uint64x2
vcgeq_u64 (uint64x2_t a, uint64x2_t b)
{
  return (a >= b);
}

DEF: uint32
vcges_f32 (float32_t a, float32_t b)
{
  return a >= b ? -1 : 0;
}

DEF: uint64
vcged_s64 (int64_t a, int64_t b)
{
  return a >= b ? -1ll : 0ll;
}

DEF: uint64
vcged_u64 (uint64_t a, uint64_t b)
{
  return a >= b ? -1ll : 0ll;
}

DEF: uint64
vcged_f64 (float64_t a, float64_t b)
{
  return a >= b ? -1ll : 0ll;
}

DEF: uint32x2
vcgez_f32 (float32x2_t a)
{
  return (uint32x2_t) (a >= 0.0f);
}

DEF: uint64x1
vcgez_f64 (float64x1_t a)
{
  return (uint64x1_t) (a[0] >= (float64x1_t) {0.0});
}

DEF: uint8x8
vcgez_s8 (int8x8_t a)
{
  return (uint8x8_t) (a >= 0);
}

DEF: uint16x4
vcgez_s16 (int16x4_t a)
{
  return (uint16x4_t) (a >= 0);
}

DEF: uint32x2
vcgez_s32 (int32x2_t a)
{
  return (uint32x2_t) (a >= 0);
}

DEF: uint64x1
vcgez_s64 (int64x1_t a)
{
  return (uint64x1_t) (a >= ((int64_t) 0));
}

DEF: uint32x4
vcgezq_f32 (float32x4_t a)
{
  return (uint32x4_t) (a >= 0.0f);
}

DEF: uint64x2
vcgezq_f64 (float64x2_t a)
{
  return (uint64x2_t) (a >= 0.0);
}

DEF: uint8x16
vcgezq_s8 (int8x16_t a)
{
  return (uint8x16_t) (a >= 0);
}

DEF: uint16x8
vcgezq_s16 (int16x8_t a)
{
  return (uint16x8_t) (a >= 0);
}

DEF: uint32x4
vcgezq_s32 (int32x4_t a)
{
  return (uint32x4_t) (a >= 0);
}

DEF: uint64x2
vcgezq_s64 (int64x2_t a)
{
  return (uint64x2_t) (a >= ((int64_t) 0));
}

DEF: uint32
vcgezs_f32 (float32_t a)
{
  return a >= 0.0f ? -1 : 0;
}

DEF: uint64
vcgezd_s64 (int64_t a)
{
  return a >= 0 ? -1ll : 0ll;
}

DEF: uint64
vcgezd_f64 (float64_t a)
{
  return a >= 0.0 ? -1ll : 0ll;
}

DEF: uint32x2
vcgt_f32 (float32x2_t a, float32x2_t b)
{
  return (uint32x2_t) (a > b);
}

DEF: uint64x1
vcgt_f64 (float64x1_t a, float64x1_t b)
{
  return (uint64x1_t) (a > b);
}

DEF: uint8x8
vcgt_s8 (int8x8_t a, int8x8_t b)
{
  return (uint8x8_t) (a > b);
}

DEF: uint16x4
vcgt_s16 (int16x4_t a, int16x4_t b)
{
  return (uint16x4_t) (a > b);
}

DEF: uint32x2
vcgt_s32 (int32x2_t a, int32x2_t b)
{
  return (uint32x2_t) (a > b);
}

DEF: uint64x1
vcgt_s64 (int64x1_t a, int64x1_t b)
{
  return (uint64x1_t) (a > b);
}

DEF: uint8x8
vcgt_u8 (uint8x8_t a, uint8x8_t b)
{
  return (a > b);
}

DEF: uint16x4
vcgt_u16 (uint16x4_t a, uint16x4_t b)
{
  return (a > b);
}

DEF: uint32x2
vcgt_u32 (uint32x2_t a, uint32x2_t b)
{
  return (a > b);
}

DEF: uint64x1
vcgt_u64 (uint64x1_t a, uint64x1_t b)
{
  return (a > b);
}

DEF: uint32x4
vcgtq_f32 (float32x4_t a, float32x4_t b)
{
  return (uint32x4_t) (a > b);
}

DEF: uint64x2
vcgtq_f64 (float64x2_t a, float64x2_t b)
{
  return (uint64x2_t) (a > b);
}

DEF: uint8x16
vcgtq_s8 (int8x16_t a, int8x16_t b)
{
  return (uint8x16_t) (a > b);
}

DEF: uint16x8
vcgtq_s16 (int16x8_t a, int16x8_t b)
{
  return (uint16x8_t) (a > b);
}

DEF: uint32x4
vcgtq_s32 (int32x4_t a, int32x4_t b)
{
  return (uint32x4_t) (a > b);
}

DEF: uint64x2
vcgtq_s64 (int64x2_t a, int64x2_t b)
{
  return (uint64x2_t) (a > b);
}

DEF: uint8x16
vcgtq_u8 (uint8x16_t a, uint8x16_t b)
{
  return (a > b);
}

DEF: uint16x8
vcgtq_u16 (uint16x8_t a, uint16x8_t b)
{
  return (a > b);
}

DEF: uint32x4
vcgtq_u32 (uint32x4_t a, uint32x4_t b)
{
  return (a > b);
}

DEF: uint64x2
vcgtq_u64 (uint64x2_t a, uint64x2_t b)
{
  return (a > b);
}

DEF: uint32
vcgts_f32 (float32_t a, float32_t b)
{
  return a > b ? -1 : 0;
}

DEF: uint64
vcgtd_s64 (int64_t a, int64_t b)
{
  return a > b ? -1ll : 0ll;
}

DEF: uint64
vcgtd_u64 (uint64_t a, uint64_t b)
{
  return a > b ? -1ll : 0ll;
}

DEF: uint64
vcgtd_f64 (float64_t a, float64_t b)
{
  return a > b ? -1ll : 0ll;
}

DEF: uint32x2
vcgtz_f32 (float32x2_t a)
{
  return (uint32x2_t) (a > 0.0f);
}

DEF: uint64x1
vcgtz_f64 (float64x1_t a)
{
  return (uint64x1_t) (a > (float64x1_t) {0.0});
}

DEF: uint8x8
vcgtz_s8 (int8x8_t a)
{
  return (uint8x8_t) (a > 0);
}

DEF: uint16x4
vcgtz_s16 (int16x4_t a)
{
  return (uint16x4_t) (a > 0);
}

DEF: uint32x2
vcgtz_s32 (int32x2_t a)
{
  return (uint32x2_t) (a > 0);
}

DEF: uint64x1
vcgtz_s64 (int64x1_t a)
{
  return (uint64x1_t) (a > ((int64_t) 0));
}

DEF: uint32x4
vcgtzq_f32 (float32x4_t a)
{
  return (uint32x4_t) (a > 0.0f);
}

DEF: uint64x2
vcgtzq_f64 (float64x2_t a)
{
    return (uint64x2_t) (a > 0.0);
}

DEF: uint8x16
vcgtzq_s8 (int8x16_t a)
{
  return (uint8x16_t) (a > 0);
}

DEF: uint16x8
vcgtzq_s16 (int16x8_t a)
{
  return (uint16x8_t) (a > 0);
}

DEF: uint32x4
vcgtzq_s32 (int32x4_t a)
{
  return (uint32x4_t) (a > 0);
}

DEF: uint64x2
vcgtzq_s64 (int64x2_t a)
{
  return (uint64x2_t) (a > ((int64_t) 0));
}

DEF: uint32
vcgtzs_f32 (float32_t a)
{
  return a > 0.0f ? -1 : 0;
}

DEF: uint64
vcgtzd_s64 (int64_t a)
{
  return a > 0 ? -1ll : 0ll;
}

DEF: uint64
vcgtzd_f64 (float64_t a)
{
  return a > 0.0 ? -1ll : 0ll;
}

DEF: uint32x2
vcle_f32 (float32x2_t a, float32x2_t b)
{
  return (uint32x2_t) (a <= b);
}

DEF: uint64x1
vcle_f64 (float64x1_t a, float64x1_t b)
{
  return (uint64x1_t) (a <= b);
}

DEF: uint8x8
vcle_s8 (int8x8_t a, int8x8_t b)
{
  return (uint8x8_t) (a <= b);
}

DEF: uint16x4
vcle_s16 (int16x4_t a, int16x4_t b)
{
  return (uint16x4_t) (a <= b);
}

DEF: uint32x2
vcle_s32 (int32x2_t a, int32x2_t b)
{
  return (uint32x2_t) (a <= b);
}

DEF: uint64x1
vcle_s64 (int64x1_t a, int64x1_t b)
{
  return (uint64x1_t) (a <= b);
}

DEF: uint8x8
vcle_u8 (uint8x8_t a, uint8x8_t b)
{
  return (a <= b);
}

DEF: uint16x4
vcle_u16 (uint16x4_t a, uint16x4_t b)
{
  return (a <= b);
}

DEF: uint32x2
vcle_u32 (uint32x2_t a, uint32x2_t b)
{
  return (a <= b);
}

DEF: uint64x1
vcle_u64 (uint64x1_t a, uint64x1_t b)
{
  return (a <= b);
}

DEF: uint32x4
vcleq_f32 (float32x4_t a, float32x4_t b)
{
  return (uint32x4_t) (a <= b);
}

DEF: uint64x2
vcleq_f64 (float64x2_t a, float64x2_t b)
{
  return (uint64x2_t) (a <= b);
}

DEF: uint8x16
vcleq_s8 (int8x16_t a, int8x16_t b)
{
  return (uint8x16_t) (a <= b);
}

DEF: uint16x8
vcleq_s16 (int16x8_t a, int16x8_t b)
{
  return (uint16x8_t) (a <= b);
}

DEF: uint32x4
vcleq_s32 (int32x4_t a, int32x4_t b)
{
  return (uint32x4_t) (a <= b);
}

DEF: uint64x2
vcleq_s64 (int64x2_t a, int64x2_t b)
{
  return (uint64x2_t) (a <= b);
}

DEF: uint8x16
vcleq_u8 (uint8x16_t a, uint8x16_t b)
{
  return (a <= b);
}

DEF: uint16x8
vcleq_u16 (uint16x8_t a, uint16x8_t b)
{
  return (a <= b);
}

DEF: uint32x4
vcleq_u32 (uint32x4_t a, uint32x4_t b)
{
  return (a <= b);
}

DEF: uint64x2
vcleq_u64 (uint64x2_t a, uint64x2_t b)
{
  return (a <= b);
}

DEF: uint32
vcles_f32 (float32_t a, float32_t b)
{
  return a <= b ? -1 : 0;
}

DEF: uint64
vcled_s64 (int64_t a, int64_t b)
{
  return a <= b ? -1ll : 0ll;
}

DEF: uint64
vcled_u64 (uint64_t a, uint64_t b)
{
  return a <= b ? -1ll : 0ll;
}

DEF: uint64
vcled_f64 (float64_t a, float64_t b)
{
  return a <= b ? -1ll : 0ll;
}

DEF: uint32x2
vclez_f32 (float32x2_t a)
{
  return (uint32x2_t) (a <= 0.0f);
}

DEF: uint64x1
vclez_f64 (float64x1_t a)
{
  return (uint64x1_t) (a <= (float64x1_t) {0.0});
}

DEF: uint8x8
vclez_s8 (int8x8_t a)
{
  return (uint8x8_t) (a <= 0);
}

DEF: uint16x4
vclez_s16 (int16x4_t a)
{
  return (uint16x4_t) (a <= 0);
}

DEF: uint32x2
vclez_s32 (int32x2_t a)
{
  return (uint32x2_t) (a <= 0);
}

DEF: uint64x1
vclez_s64 (int64x1_t a)
{
  return (uint64x1_t) (a <= ((int64_t) 0));
}

DEF: uint32x4
vclezq_f32 (float32x4_t a)
{
  return (uint32x4_t) (a <= 0.0f);
}

DEF: uint64x2
vclezq_f64 (float64x2_t a)
{
  return (uint64x2_t) (a <= 0.0);
}

DEF: uint8x16
vclezq_s8 (int8x16_t a)
{
  return (uint8x16_t) (a <= 0);
}

DEF: uint16x8
vclezq_s16 (int16x8_t a)
{
  return (uint16x8_t) (a <= 0);
}

DEF: uint32x4
vclezq_s32 (int32x4_t a)
{
  return (uint32x4_t) (a <= 0);
}

DEF: uint64x2
vclezq_s64 (int64x2_t a)
{
  return (uint64x2_t) (a <= ((int64_t) 0));
}

DEF: uint32
vclezs_f32 (float32_t a)
{
  return a <= 0.0f ? -1 : 0;
}

DEF: uint64
vclezd_s64 (int64_t a)
{
  return a <= 0 ? -1ll : 0ll;
}

DEF: uint64
vclezd_f64 (float64_t a)
{
  return a <= 0.0 ? -1ll : 0ll;
}

DEF: uint32x2
vclt_f32 (float32x2_t a, float32x2_t b)
{
  return (uint32x2_t) (a < b);
}

DEF: uint64x1
vclt_f64 (float64x1_t a, float64x1_t b)
{
  return (uint64x1_t) (a < b);
}

DEF: uint8x8
vclt_s8 (int8x8_t a, int8x8_t b)
{
  return (uint8x8_t) (a < b);
}

DEF: uint16x4
vclt_s16 (int16x4_t a, int16x4_t b)
{
  return (uint16x4_t) (a < b);
}

DEF: uint32x2
vclt_s32 (int32x2_t a, int32x2_t b)
{
  return (uint32x2_t) (a < b);
}

DEF: uint64x1
vclt_s64 (int64x1_t a, int64x1_t b)
{
  return (uint64x1_t) (a < b);
}

DEF: uint8x8
vclt_u8 (uint8x8_t a, uint8x8_t b)
{
  return (a < b);
}

DEF: uint16x4
vclt_u16 (uint16x4_t a, uint16x4_t b)
{
  return (a < b);
}

DEF: uint32x2
vclt_u32 (uint32x2_t a, uint32x2_t b)
{
  return (a < b);
}

DEF: uint64x1
vclt_u64 (uint64x1_t a, uint64x1_t b)
{
  return (a < b);
}

DEF: uint32x4
vcltq_f32 (float32x4_t a, float32x4_t b)
{
  return (uint32x4_t) (a < b);
}

DEF: uint64x2
vcltq_f64 (float64x2_t a, float64x2_t b)
{
  return (uint64x2_t) (a < b);
}

DEF: uint8x16
vcltq_s8 (int8x16_t a, int8x16_t b)
{
  return (uint8x16_t) (a < b);
}

DEF: uint16x8
vcltq_s16 (int16x8_t a, int16x8_t b)
{
  return (uint16x8_t) (a < b);
}

DEF: uint32x4
vcltq_s32 (int32x4_t a, int32x4_t b)
{
  return (uint32x4_t) (a < b);
}

DEF: uint64x2
vcltq_s64 (int64x2_t a, int64x2_t b)
{
  return (uint64x2_t) (a < b);
}

DEF: uint8x16
vcltq_u8 (uint8x16_t a, uint8x16_t b)
{
  return (a < b);
}

DEF: uint16x8
vcltq_u16 (uint16x8_t a, uint16x8_t b)
{
  return (a < b);
}

DEF: uint32x4
vcltq_u32 (uint32x4_t a, uint32x4_t b)
{
  return (a < b);
}

DEF: uint64x2
vcltq_u64 (uint64x2_t a, uint64x2_t b)
{
  return (a < b);
}

DEF: uint32
vclts_f32 (float32_t a, float32_t b)
{
  return a < b ? -1 : 0;
}

DEF: uint64
vcltd_s64 (int64_t a, int64_t b)
{
  return a < b ? -1ll : 0ll;
}

DEF: uint64
vcltd_u64 (uint64_t a, uint64_t b)
{
  return a < b ? -1ll : 0ll;
}

DEF: uint64
vcltd_f64 (float64_t a, float64_t b)
{
  return a < b ? -1ll : 0ll;
}

DEF: uint32x2
vcltz_f32 (float32x2_t a)
{
  return (uint32x2_t) (a < 0.0f);
}

DEF: uint64x1
vcltz_f64 (float64x1_t a)
{
  return (uint64x1_t) (a < (float64x1_t) {0.0});
}

DEF: uint8x8
vcltz_s8 (int8x8_t a)
{
  return (uint8x8_t) (a < 0);
}

DEF: uint16x4
vcltz_s16 (int16x4_t a)
{
  return (uint16x4_t) (a < 0);
}

DEF: uint32x2
vcltz_s32 (int32x2_t a)
{
  return (uint32x2_t) (a < 0);
}

DEF: uint64x1
vcltz_s64 (int64x1_t a)
{
  return (uint64x1_t) (a < ((int64_t) 0));
}

DEF: uint32x4
vcltzq_f32 (float32x4_t a)
{
  return (uint32x4_t) (a < 0.0f);
}

DEF: uint64x2
vcltzq_f64 (float64x2_t a)
{
  return (uint64x2_t) (a < 0.0);
}

DEF: uint8x16
vcltzq_s8 (int8x16_t a)
{
  return (uint8x16_t) (a < 0);
}

DEF: uint16x8
vcltzq_s16 (int16x8_t a)
{
  return (uint16x8_t) (a < 0);
}

DEF: uint32x4
vcltzq_s32 (int32x4_t a)
{
  return (uint32x4_t) (a < 0);
}

DEF: uint64x2
vcltzq_s64 (int64x2_t a)
{
  return (uint64x2_t) (a < ((int64_t) 0));
}

DEF: uint32
vcltzs_f32 (float32_t a)
{
  return a < 0.0f ? -1 : 0;
}

DEF: uint64
vcltzd_s64 (int64_t a)
{
  return a < 0 ? -1ll : 0ll;
}

DEF: uint64
vcltzd_f64 (float64_t a)
{
  return a < 0.0 ? -1ll : 0ll;
}

DEF: int8x8
vcls_s8 (int8x8_t a)
{
  return builtin_aarch64_clrsbv8qi (a);
}

DEF: int16x4
vcls_s16 (int16x4_t a)
{
  return builtin_aarch64_clrsbv4hi (a);
}

DEF: int32x2
vcls_s32 (int32x2_t a)
{
  return builtin_aarch64_clrsbv2si (a);
}

DEF: int8x16
vclsq_s8 (int8x16_t a)
{
  return builtin_aarch64_clrsbv16qi (a);
}

DEF: int16x8
vclsq_s16 (int16x8_t a)
{
  return builtin_aarch64_clrsbv8hi (a);
}

DEF: int32x4
vclsq_s32 (int32x4_t a)
{
  return builtin_aarch64_clrsbv4si (a);
}

DEF: int8x8
vclz_s8 (int8x8_t a)
{
  return builtin_aarch64_clzv8qi (a);
}

DEF: int16x4
vclz_s16 (int16x4_t a)
{
  return builtin_aarch64_clzv4hi (a);
}

DEF: int32x2
vclz_s32 (int32x2_t a)
{
  return builtin_aarch64_clzv2si (a);
}

DEF: uint8x8
vclz_u8 (uint8x8_t a)
{
  return (uint8x8_t)builtin_aarch64_clzv8qi ((int8x8_t)a);
}

DEF: uint16x4
vclz_u16 (uint16x4_t a)
{
  return (uint16x4_t)builtin_aarch64_clzv4hi ((int16x4_t)a);
}

DEF: uint32x2
vclz_u32 (uint32x2_t a)
{
  return (uint32x2_t)builtin_aarch64_clzv2si ((int32x2_t)a);
}

DEF: int8x16
vclzq_s8 (int8x16_t a)
{
  return builtin_aarch64_clzv16qi (a);
}

DEF: int16x8
vclzq_s16 (int16x8_t a)
{
  return builtin_aarch64_clzv8hi (a);
}

DEF: int32x4
vclzq_s32 (int32x4_t a)
{
  return builtin_aarch64_clzv4si (a);
}

DEF: uint8x16
vclzq_u8 (uint8x16_t a)
{
  return (uint8x16_t)builtin_aarch64_clzv16qi ((int8x16_t)a);
}

DEF: uint16x8
vclzq_u16 (uint16x8_t a)
{
  return (uint16x8_t)builtin_aarch64_clzv8hi ((int16x8_t)a);
}

DEF: uint32x4
vclzq_u32 (uint32x4_t a)
{
  return (uint32x4_t)builtin_aarch64_clzv4si ((int32x4_t)a);
}

DEF: poly8x8
vcnt_p8 (poly8x8_t a)
{
  return (poly8x8_t) builtin_aarch64_popcountv8qi ((int8x8_t) a);
}

DEF: int8x8
vcnt_s8 (int8x8_t a)
{
  return builtin_aarch64_popcountv8qi (a);
}

DEF: uint8x8
vcnt_u8 (uint8x8_t a)
{
  return (uint8x8_t) builtin_aarch64_popcountv8qi ((int8x8_t) a);
}

DEF: poly8x16
vcntq_p8 (poly8x16_t a)
{
  return (poly8x16_t) builtin_aarch64_popcountv16qi ((int8x16_t) a);
}

DEF: int8x16
vcntq_s8 (int8x16_t a)
{
  return builtin_aarch64_popcountv16qi (a);
}

DEF: uint8x16
vcntq_u8 (uint8x16_t a)
{
  return (uint8x16_t) builtin_aarch64_popcountv16qi ((int8x16_t) a);
}

DEF: float32x2
vcvt_f32_f64 (float64x2_t a)
{
  return builtin_aarch64_float_truncate_lo_v2sf (a);
}

DEF: float32x4
vcvt_high_f32_f64 (float32x2_t a, float64x2_t b)
{
  return builtin_aarch64_float_truncate_hi_v4sf (a, b);
}

DEF: float64x2
vcvt_f64_f32 (float32x2_t a)
{

  return builtin_aarch64_float_extend_lo_v2df (a);
}

DEF: float64x2
vcvt_high_f64_f32 (float32x4_t a)
{
  return builtin_aarch64_vec_unpacks_hi_v4sf (a);
}

DEF: float64
vcvtd_f64_s64 (int64_t a)
{
  return (float64_t) a;
}

DEF: float64
vcvtd_f64_u64 (uint64_t a)
{
  return (float64_t) a;
}

DEF: float32
vcvts_f32_s32 (int32_t a)
{
  return (float32_t) a;
}

DEF: float32
vcvts_f32_u32 (uint32_t a)
{
  return (float32_t) a;
}

DEF: float32x2
vcvt_f32_s32 (int32x2_t a)
{
  return builtin_aarch64_floatv2siv2sf (a);
}

DEF: float32x2
vcvt_f32_u32 (uint32x2_t a)
{
  return builtin_aarch64_floatunsv2siv2sf ((int32x2_t) a);
}

DEF: float32x4
vcvtq_f32_s32 (int32x4_t a)
{
  return builtin_aarch64_floatv4siv4sf (a);
}

DEF: float32x4
vcvtq_f32_u32 (uint32x4_t a)
{
  return builtin_aarch64_floatunsv4siv4sf ((int32x4_t) a);
}

DEF: float64x2
vcvtq_f64_s64 (int64x2_t a)
{
  return builtin_aarch64_floatv2div2df (a);
}

DEF: float64x2
vcvtq_f64_u64 (uint64x2_t a)
{
  return builtin_aarch64_floatunsv2div2df ((int64x2_t) a);
}

DEF: int64
vcvtd_s64_f64 (float64_t a)
{
  return (int64_t) a;
}

DEF: uint64
vcvtd_u64_f64 (float64_t a)
{
  return (uint64_t) a;
}

DEF: int32
vcvts_s32_f32 (float32_t a)
{
  return (int32_t) a;
}

DEF: uint32
vcvts_u32_f32 (float32_t a)
{
  return (uint32_t) a;
}

DEF: int32x2
vcvt_s32_f32 (float32x2_t a)
{
  return builtin_aarch64_lbtruncv2sfv2si (a);
}

DEF: uint32x2
vcvt_u32_f32 (float32x2_t a)
{
  return (uint32x2_t) builtin_aarch64_lbtruncuv2sfv2si (a);
}

DEF: int32x4
vcvtq_s32_f32 (float32x4_t a)
{
  return builtin_aarch64_lbtruncv4sfv4si (a);
}

DEF: uint32x4
vcvtq_u32_f32 (float32x4_t a)
{
  return (uint32x4_t) builtin_aarch64_lbtruncuv4sfv4si (a);
}

DEF: int64x2
vcvtq_s64_f64 (float64x2_t a)
{
  return builtin_aarch64_lbtruncv2dfv2di (a);
}

DEF: uint64x2
vcvtq_u64_f64 (float64x2_t a)
{
  return (uint64x2_t) builtin_aarch64_lbtruncuv2dfv2di (a);
}

DEF: int64
vcvtad_s64_f64 (float64_t a)
{
  return builtin_aarch64_lrounddfdi (a);
}

DEF: uint64
vcvtad_u64_f64 (float64_t a)
{
  return builtin_aarch64_lroundudfdi (a);
}

DEF: int32
vcvtas_s32_f32 (float32_t a)
{
  return builtin_aarch64_lroundsfsi (a);
}

DEF: uint32
vcvtas_u32_f32 (float32_t a)
{
  return builtin_aarch64_lroundusfsi (a);
}

DEF: int32x2
vcvta_s32_f32 (float32x2_t a)
{
  return builtin_aarch64_lroundv2sfv2si (a);
}

DEF: uint32x2
vcvta_u32_f32 (float32x2_t a)
{
  return (uint32x2_t) builtin_aarch64_lrounduv2sfv2si (a);
}

DEF: int32x4
vcvtaq_s32_f32 (float32x4_t a)
{
  return builtin_aarch64_lroundv4sfv4si (a);
}

DEF: uint32x4
vcvtaq_u32_f32 (float32x4_t a)
{
  return (uint32x4_t) builtin_aarch64_lrounduv4sfv4si (a);
}

DEF: int64x2
vcvtaq_s64_f64 (float64x2_t a)
{
  return builtin_aarch64_lroundv2dfv2di (a);
}

DEF: uint64x2
vcvtaq_u64_f64 (float64x2_t a)
{
  return (uint64x2_t) builtin_aarch64_lrounduv2dfv2di (a);
}

DEF: int64
vcvtmd_s64_f64 (float64_t a)
{
  return builtin_llfloor (a);
}

DEF: uint64
vcvtmd_u64_f64 (float64_t a)
{
  return builtin_aarch64_lfloorudfdi (a);
}

DEF: int32
vcvtms_s32_f32 (float32_t a)
{
  return builtin_ifloorf (a);
}

DEF: uint32
vcvtms_u32_f32 (float32_t a)
{
  return builtin_aarch64_lfloorusfsi (a);
}

DEF: int32x2
vcvtm_s32_f32 (float32x2_t a)
{
  return builtin_aarch64_lfloorv2sfv2si (a);
}

DEF: uint32x2
vcvtm_u32_f32 (float32x2_t a)
{
  return (uint32x2_t) builtin_aarch64_lflooruv2sfv2si (a);
}

DEF: int32x4
vcvtmq_s32_f32 (float32x4_t a)
{
  return builtin_aarch64_lfloorv4sfv4si (a);
}

DEF: uint32x4
vcvtmq_u32_f32 (float32x4_t a)
{
  return (uint32x4_t) builtin_aarch64_lflooruv4sfv4si (a);
}

DEF: int64x2
vcvtmq_s64_f64 (float64x2_t a)
{
  return builtin_aarch64_lfloorv2dfv2di (a);
}

DEF: uint64x2
vcvtmq_u64_f64 (float64x2_t a)
{
  return (uint64x2_t) builtin_aarch64_lflooruv2dfv2di (a);
}

DEF: int64
vcvtnd_s64_f64 (float64_t a)
{
  return builtin_aarch64_lfrintndfdi (a);
}

DEF: uint64
vcvtnd_u64_f64 (float64_t a)
{
  return builtin_aarch64_lfrintnudfdi (a);
}

DEF: int32
vcvtns_s32_f32 (float32_t a)
{
  return builtin_aarch64_lfrintnsfsi (a);
}

DEF: uint32
vcvtns_u32_f32 (float32_t a)
{
  return builtin_aarch64_lfrintnusfsi (a);
}

DEF: int32x2
vcvtn_s32_f32 (float32x2_t a)
{
  return builtin_aarch64_lfrintnv2sfv2si (a);
}

DEF: uint32x2
vcvtn_u32_f32 (float32x2_t a)
{
  return (uint32x2_t) builtin_aarch64_lfrintnuv2sfv2si (a);
}

DEF: int32x4
vcvtnq_s32_f32 (float32x4_t a)
{
  return builtin_aarch64_lfrintnv4sfv4si (a);
}

DEF: uint32x4
vcvtnq_u32_f32 (float32x4_t a)
{
  return (uint32x4_t) builtin_aarch64_lfrintnuv4sfv4si (a);
}

DEF: int64x2
vcvtnq_s64_f64 (float64x2_t a)
{
  return builtin_aarch64_lfrintnv2dfv2di (a);
}

DEF: uint64x2
vcvtnq_u64_f64 (float64x2_t a)
{
  return (uint64x2_t) builtin_aarch64_lfrintnuv2dfv2di (a);
}

DEF: int64
vcvtpd_s64_f64 (float64_t a)
{
  return builtin_llceil (a);
}

DEF: uint64
vcvtpd_u64_f64 (float64_t a)
{
  return builtin_aarch64_lceiludfdi (a);
}

DEF: int32
vcvtps_s32_f32 (float32_t a)
{
  return builtin_iceilf (a);
}

DEF: uint32
vcvtps_u32_f32 (float32_t a)
{
  return builtin_aarch64_lceilusfsi (a);
}

DEF: int32x2
vcvtp_s32_f32 (float32x2_t a)
{
  return builtin_aarch64_lceilv2sfv2si (a);
}

DEF: uint32x2
vcvtp_u32_f32 (float32x2_t a)
{
  return (uint32x2_t) builtin_aarch64_lceiluv2sfv2si (a);
}

DEF: int32x4
vcvtpq_s32_f32 (float32x4_t a)
{
  return builtin_aarch64_lceilv4sfv4si (a);
}

DEF: uint32x4
vcvtpq_u32_f32 (float32x4_t a)
{
  return (uint32x4_t) builtin_aarch64_lceiluv4sfv4si (a);
}

DEF: int64x2
vcvtpq_s64_f64 (float64x2_t a)
{
  return builtin_aarch64_lceilv2dfv2di (a);
}

DEF: uint64x2
vcvtpq_u64_f64 (float64x2_t a)
{
  return (uint64x2_t) builtin_aarch64_lceiluv2dfv2di (a);
}

DEF: float32x2
vdup_n_f32 (float32_t a)
{
  return (float32x2_t) {a, a};
}

DEF: float64x1
vdup_n_f64 (float64_t a)
{
  return (float64x1_t) {a};
}

DEF: poly8x8
vdup_n_p8 (poly8_t a)
{
  return (poly8x8_t) {a, a, a, a, a, a, a, a};
}

DEF: poly16x4
vdup_n_p16 (poly16_t a)
{
  return (poly16x4_t) {a, a, a, a};
}

DEF: int8x8
vdup_n_s8 (int8_t a)
{
  return (int8x8_t) {a, a, a, a, a, a, a, a};
}

DEF: int16x4
vdup_n_s16 (int16_t a)
{
  return (int16x4_t) {a, a, a, a};
}

DEF: int32x2
vdup_n_s32 (int32_t a)
{
  return (int32x2_t) {a, a};
}

DEF: int64x1
vdup_n_s64 (int64_t a)
{
  return (int64x1_t) {a};
}

DEF: uint8x8
vdup_n_u8 (uint8_t a)
{
  return (uint8x8_t) {a, a, a, a, a, a, a, a};
}

DEF: uint16x4
vdup_n_u16 (uint16_t a)
{
  return (uint16x4_t) {a, a, a, a};
}

DEF: uint32x2
vdup_n_u32 (uint32_t a)
{
  return (uint32x2_t) {a, a};
}

DEF: uint64x1
vdup_n_u64 (uint64_t a)
{
  return (uint64x1_t) {a};
}

DEF: float32x4
vdupq_n_f32 (float32_t a)
{
  return (float32x4_t) {a, a, a, a};
}

DEF: float64x2
vdupq_n_f64 (float64_t a)
{
  return (float64x2_t) {a, a};
}

DEF: poly8x16
vdupq_n_p8 (uint32_t a)
{
  return (poly8x16_t) {a, a, a, a, a, a, a, a,
         a, a, a, a, a, a, a, a};
}

DEF: poly16x8
vdupq_n_p16 (uint32_t a)
{
  return (poly16x8_t) {a, a, a, a, a, a, a, a};
}

DEF: int8x16
vdupq_n_s8 (int32_t a)
{
  return (int8x16_t) {a, a, a, a, a, a, a, a,
        a, a, a, a, a, a, a, a};
}

DEF: int16x8
vdupq_n_s16 (int32_t a)
{
  return (int16x8_t) {a, a, a, a, a, a, a, a};
}

DEF: int32x4
vdupq_n_s32 (int32_t a)
{
  return (int32x4_t) {a, a, a, a};
}

DEF: int64x2
vdupq_n_s64 (int64_t a)
{
  return (int64x2_t) {a, a};
}

DEF: uint8x16
vdupq_n_u8 (uint32_t a)
{
  return (uint8x16_t) {a, a, a, a, a, a, a, a,
         a, a, a, a, a, a, a, a};
}

DEF: uint16x8
vdupq_n_u16 (uint32_t a)
{
  return (uint16x8_t) {a, a, a, a, a, a, a, a};
}

DEF: uint32x4
vdupq_n_u32 (uint32_t a)
{
  return (uint32x4_t) {a, a, a, a};
}

DEF: uint64x2
vdupq_n_u64 (uint64_t a)
{
  return (uint64x2_t) {a, a};
}

DEF: float32x2
vdup_lane_f32 (float32x2_t a, const int b)
{
  return vdup_n_f32 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: float64x1
vdup_lane_f64 (float64x1_t a, const int b)
{
  return vdup_n_f64 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: poly8x8
vdup_lane_p8 (poly8x8_t a, const int b)
{
  return vdup_n_p8 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: poly16x4
vdup_lane_p16 (poly16x4_t a, const int b)
{
  return vdup_n_p16 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: int8x8
vdup_lane_s8 (int8x8_t a, const int b)
{
  return vdup_n_s8 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: int16x4
vdup_lane_s16 (int16x4_t a, const int b)
{
  return vdup_n_s16 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: int32x2
vdup_lane_s32 (int32x2_t a, const int b)
{
  return vdup_n_s32 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: int64x1
vdup_lane_s64 (int64x1_t a, const int b)
{
  return vdup_n_s64 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: uint8x8
vdup_lane_u8 (uint8x8_t a, const int b)
{
  return vdup_n_u8 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: uint16x4
vdup_lane_u16 (uint16x4_t a, const int b)
{
  return vdup_n_u16 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: uint32x2
vdup_lane_u32 (uint32x2_t a, const int b)
{
  return vdup_n_u32 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: uint64x1
vdup_lane_u64 (uint64x1_t a, const int b)
{
  return vdup_n_u64 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: float32x2
vdup_laneq_f32 (float32x4_t a, const int b)
{
  return vdup_n_f32 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: float64x1
vdup_laneq_f64 (float64x2_t a, const int b)
{
  return vdup_n_f64 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: poly8x8
vdup_laneq_p8 (poly8x16_t a, const int b)
{
  return vdup_n_p8 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: poly16x4
vdup_laneq_p16 (poly16x8_t a, const int b)
{
  return vdup_n_p16 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: int8x8
vdup_laneq_s8 (int8x16_t a, const int b)
{
  return vdup_n_s8 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: int16x4
vdup_laneq_s16 (int16x8_t a, const int b)
{
  return vdup_n_s16 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: int32x2
vdup_laneq_s32 (int32x4_t a, const int b)
{
  return vdup_n_s32 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: int64x1
vdup_laneq_s64 (int64x2_t a, const int b)
{
  return vdup_n_s64 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: uint8x8
vdup_laneq_u8 (uint8x16_t a, const int b)
{
  return vdup_n_u8 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: uint16x4
vdup_laneq_u16 (uint16x8_t a, const int b)
{
  return vdup_n_u16 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: uint32x2
vdup_laneq_u32 (uint32x4_t a, const int b)
{
  return vdup_n_u32 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: uint64x1
vdup_laneq_u64 (uint64x2_t a, const int b)
{
  return vdup_n_u64 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}
DEF: float32x4
vdupq_lane_f32 (float32x2_t a, const int b)
{
  return vdupq_n_f32 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: float64x2
vdupq_lane_f64 (float64x1_t a, const int b)
{
  return vdupq_n_f64 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: poly8x16
vdupq_lane_p8 (poly8x8_t a, const int b)
{
  return vdupq_n_p8 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: poly16x8
vdupq_lane_p16 (poly16x4_t a, const int b)
{
  return vdupq_n_p16 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: int8x16
vdupq_lane_s8 (int8x8_t a, const int b)
{
  return vdupq_n_s8 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: int16x8
vdupq_lane_s16 (int16x4_t a, const int b)
{
  return vdupq_n_s16 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: int32x4
vdupq_lane_s32 (int32x2_t a, const int b)
{
  return vdupq_n_s32 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: int64x2
vdupq_lane_s64 (int64x1_t a, const int b)
{
  return vdupq_n_s64 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: uint8x16
vdupq_lane_u8 (uint8x8_t a, const int b)
{
  return vdupq_n_u8 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: uint16x8
vdupq_lane_u16 (uint16x4_t a, const int b)
{
  return vdupq_n_u16 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: uint32x4
vdupq_lane_u32 (uint32x2_t a, const int b)
{
  return vdupq_n_u32 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: uint64x2
vdupq_lane_u64 (uint64x1_t a, const int b)
{
  return vdupq_n_u64 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}
DEF: float32x4
vdupq_laneq_f32 (float32x4_t a, const int b)
{
  return vdupq_n_f32 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: float64x2
vdupq_laneq_f64 (float64x2_t a, const int b)
{
  return vdupq_n_f64 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: poly8x16
vdupq_laneq_p8 (poly8x16_t a, const int b)
{
  return vdupq_n_p8 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: poly16x8
vdupq_laneq_p16 (poly16x8_t a, const int b)
{
  return vdupq_n_p16 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: int8x16
vdupq_laneq_s8 (int8x16_t a, const int b)
{
  return vdupq_n_s8 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: int16x8
vdupq_laneq_s16 (int16x8_t a, const int b)
{
  return vdupq_n_s16 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: int32x4
vdupq_laneq_s32 (int32x4_t a, const int b)
{
  return vdupq_n_s32 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: int64x2
vdupq_laneq_s64 (int64x2_t a, const int b)
{
  return vdupq_n_s64 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: uint8x16
vdupq_laneq_u8 (uint8x16_t a, const int b)
{
  return vdupq_n_u8 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: uint16x8
vdupq_laneq_u16 (uint16x8_t a, const int b)
{
  return vdupq_n_u16 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: uint32x4
vdupq_laneq_u32 (uint32x4_t a, const int b)
{
  return vdupq_n_u32 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}

DEF: uint64x2
vdupq_laneq_u64 (uint64x2_t a, const int b)
{
  return vdupq_n_u64 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; }));
}
DEF: poly8
vdupb_lane_p8 (poly8x8_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: int8
vdupb_lane_s8 (int8x8_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: uint8
vdupb_lane_u8 (uint8x8_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}
DEF: poly16
vduph_lane_p16 (poly16x4_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: int16
vduph_lane_s16 (int16x4_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: uint16
vduph_lane_u16 (uint16x4_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}
DEF: float32
vdups_lane_f32 (float32x2_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: int32
vdups_lane_s32 (int32x2_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: uint32
vdups_lane_u32 (uint32x2_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}
DEF: float64
vdupd_lane_f64 (float64x1_t a, const int b)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b);
  return a[0];
}

DEF: int64
vdupd_lane_s64 (int64x1_t a, const int b)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b);
  return a[0];
}

DEF: uint64
vdupd_lane_u64 (uint64x1_t a, const int b)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b);
  return a[0];
}
DEF: poly8
vdupb_laneq_p8 (poly8x16_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: int8
vdupb_laneq_s8 (int8x16_t a, const int  b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: uint8
vdupb_laneq_u8 (uint8x16_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}
DEF: poly16
vduph_laneq_p16 (poly16x8_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: int16
vduph_laneq_s16 (int16x8_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: uint16
vduph_laneq_u16 (uint16x8_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}
DEF: float32
vdups_laneq_f32 (float32x4_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: int32
vdups_laneq_s32 (int32x4_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: uint32
vdups_laneq_u32 (uint32x4_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}
DEF: float64
vdupd_laneq_f64 (float64x2_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: int64
vdupd_laneq_s64 (int64x2_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: uint64
vdupd_laneq_u64 (uint64x2_t a, const int b)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), b); a[b]; });
}

DEF: float32x2
vext_f32 (float32x2_t a, float32x2_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);

  return builtin_shuffle (a, b, (uint32x2_t) {c, c+1});

}

DEF: float64x1
vext_f64 (float64x1_t a, float64x1_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);

  return a;
}
DEF: poly8x8
vext_p8 (poly8x8_t a, poly8x8_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);


  return builtin_shuffle (a, b,
      (uint8x8_t) {c, c+1, c+2, c+3, c+4, c+5, c+6, c+7});

}

DEF: poly16x4
vext_p16 (poly16x4_t a, poly16x4_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);


  return builtin_shuffle (a, b, (uint16x4_t) {c, c+1, c+2, c+3});

}

DEF: int8x8
vext_s8 (int8x8_t a, int8x8_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);


  return builtin_shuffle (a, b,
      (uint8x8_t) {c, c+1, c+2, c+3, c+4, c+5, c+6, c+7});

}

DEF: int16x4
vext_s16 (int16x4_t a, int16x4_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);


  return builtin_shuffle (a, b, (uint16x4_t) {c, c+1, c+2, c+3});

}

DEF: int32x2
vext_s32 (int32x2_t a, int32x2_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);

  return builtin_shuffle (a, b, (uint32x2_t) {c, c+1});

}

DEF: int64x1
vext_s64 (int64x1_t a, int64x1_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);

  return a;
}

DEF: uint8x8
vext_u8 (uint8x8_t a, uint8x8_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);


  return builtin_shuffle (a, b,
      (uint8x8_t) {c, c+1, c+2, c+3, c+4, c+5, c+6, c+7});

}

DEF: uint16x4
vext_u16 (uint16x4_t a, uint16x4_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);


  return builtin_shuffle (a, b, (uint16x4_t) {c, c+1, c+2, c+3});

}

DEF: uint32x2
vext_u32 (uint32x2_t a, uint32x2_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);

  return builtin_shuffle (a, b, (uint32x2_t) {c, c+1});

}

DEF: uint64x1
vext_u64 (uint64x1_t a, uint64x1_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);

  return a;
}

DEF: float32x4
vextq_f32 (float32x4_t a, float32x4_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);


  return builtin_shuffle (a, b, (uint32x4_t) {c, c+1, c+2, c+3});

}

DEF: float64x2
vextq_f64 (float64x2_t a, float64x2_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);

  return builtin_shuffle (a, b, (uint64x2_t) {c, c+1});

}

DEF: poly8x16
vextq_p8 (poly8x16_t a, poly8x16_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);

  return builtin_shuffle (a, b, (uint8x16_t)
      {c, c+1, c+2, c+3, c+4, c+5, c+6, c+7,
       c+8, c+9, c+10, c+11, c+12, c+13, c+14, c+15});

}

DEF: poly16x8
vextq_p16 (poly16x8_t a, poly16x8_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);


  return builtin_shuffle (a, b,
      (uint16x8_t) {c, c+1, c+2, c+3, c+4, c+5, c+6, c+7});

}

DEF: int8x16
vextq_s8 (int8x16_t a, int8x16_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);

  return builtin_shuffle (a, b, (uint8x16_t)
      {c, c+1, c+2, c+3, c+4, c+5, c+6, c+7,
       c+8, c+9, c+10, c+11, c+12, c+13, c+14, c+15});

}

DEF: int16x8
vextq_s16 (int16x8_t a, int16x8_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);


  return builtin_shuffle (a, b,
      (uint16x8_t) {c, c+1, c+2, c+3, c+4, c+5, c+6, c+7});

}

DEF: int32x4
vextq_s32 (int32x4_t a, int32x4_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);


  return builtin_shuffle (a, b, (uint32x4_t) {c, c+1, c+2, c+3});

}

DEF: int64x2
vextq_s64 (int64x2_t a, int64x2_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);

  return builtin_shuffle (a, b, (uint64x2_t) {c, c+1});

}

DEF: uint8x16
vextq_u8 (uint8x16_t a, uint8x16_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);

  return builtin_shuffle (a, b, (uint8x16_t)
      {c, c+1, c+2, c+3, c+4, c+5, c+6, c+7,
       c+8, c+9, c+10, c+11, c+12, c+13, c+14, c+15});

}

DEF: uint16x8
vextq_u16 (uint16x8_t a, uint16x8_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);


  return builtin_shuffle (a, b,
      (uint16x8_t) {c, c+1, c+2, c+3, c+4, c+5, c+6, c+7});

}

DEF: uint32x4
vextq_u32 (uint32x4_t a, uint32x4_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);


  return builtin_shuffle (a, b, (uint32x4_t) {c, c+1, c+2, c+3});

}

DEF: uint64x2
vextq_u64 (uint64x2_t a, uint64x2_t b, const int c)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), c);

  return builtin_shuffle (a, b, (uint64x2_t) {c, c+1});

}

DEF: float64x1
vfma_f64 (float64x1_t a, float64x1_t b, float64x1_t c)
{
  return (float64x1_t) {builtin_fma (b[0], c[0], a[0])};
}

DEF: float32x2
vfma_f32 (float32x2_t a, float32x2_t b, float32x2_t c)
{
  return builtin_aarch64_fmav2sf (b, c, a);
}

DEF: float32x4
vfmaq_f32 (float32x4_t a, float32x4_t b, float32x4_t c)
{
  return builtin_aarch64_fmav4sf (b, c, a);
}

DEF: float64x2
vfmaq_f64 (float64x2_t a, float64x2_t b, float64x2_t c)
{
  return builtin_aarch64_fmav2df (b, c, a);
}

DEF: float32x2
vfma_n_f32 (float32x2_t a, float32x2_t b, float32_t c)
{
  return builtin_aarch64_fmav2sf (b, vdup_n_f32 (c), a);
}

DEF: float32x4
vfmaq_n_f32 (float32x4_t a, float32x4_t b, float32_t c)
{
  return builtin_aarch64_fmav4sf (b, vdupq_n_f32 (c), a);
}

DEF: float64x2
vfmaq_n_f64 (float64x2_t a, float64x2_t b, float64_t c)
{
  return builtin_aarch64_fmav2df (b, vdupq_n_f64 (c), a);
}

DEF: float32x2
vfma_lane_f32 (float32x2_t a, float32x2_t b,
        float32x2_t c, const int lane)
{
  return builtin_aarch64_fmav2sf (b,
        vdup_n_f32 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })),
        a);
}

DEF: float64x1
vfma_lane_f64 (float64x1_t a, float64x1_t b,
        float64x1_t c, const int lane)
{
  return (float64x1_t) {builtin_fma (b[0], c[0], a[0])};
}

DEF: float64
vfmad_lane_f64 (float64_t a, float64_t b,
         float64x1_t c, const int lane)
{
  return builtin_fma (b, c[0], a);
}

DEF: float32
vfmas_lane_f32 (float32_t a, float32_t b,
         float32x2_t c, const int lane)
{
  return builtin_fmaf (b, extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; }), a);
}

DEF: float32x2
vfma_laneq_f32 (float32x2_t a, float32x2_t b,
         float32x4_t c, const int lane)
{
  return builtin_aarch64_fmav2sf (b,
        vdup_n_f32 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })),
        a);
}

DEF: float64x1
vfma_laneq_f64 (float64x1_t a, float64x1_t b,
         float64x2_t c, const int lane)
{
  float64_t c0 = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; });
  return (float64x1_t) {builtin_fma (b[0], c0, a[0])};
}

DEF: float64
vfmad_laneq_f64 (float64_t a, float64_t b,
          float64x2_t c, const int lane)
{
  return builtin_fma (b, extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; }), a);
}

DEF: float32
vfmas_laneq_f32 (float32_t a, float32_t b,
   float32x4_t c, const int lane)
{
  return builtin_fmaf (b, extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; }), a);
}

DEF: float32x4
vfmaq_lane_f32 (float32x4_t a, float32x4_t b,
         float32x2_t c, const int lane)
{
  return builtin_aarch64_fmav4sf (b,
        vdupq_n_f32 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })),
        a);
}

DEF: float64x2
vfmaq_lane_f64 (float64x2_t a, float64x2_t b,
         float64x1_t c, const int lane)
{
  return builtin_aarch64_fmav2df (b, vdupq_n_f64 (c[0]), a);
}

DEF: float32x4
vfmaq_laneq_f32 (float32x4_t a, float32x4_t b,
          float32x4_t c, const int lane)
{
  return builtin_aarch64_fmav4sf (b,
        vdupq_n_f32 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })),
        a);
}

DEF: float64x2
vfmaq_laneq_f64 (float64x2_t a, float64x2_t b,
          float64x2_t c, const int lane)
{
  return builtin_aarch64_fmav2df (b,
        vdupq_n_f64 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })),
        a);
}

DEF: float64x1
vfms_f64 (float64x1_t a, float64x1_t b, float64x1_t c)
{
  return (float64x1_t) {builtin_fma (-b[0], c[0], a[0])};
}

DEF: float32x2
vfms_f32 (float32x2_t a, float32x2_t b, float32x2_t c)
{
  return builtin_aarch64_fmav2sf (-b, c, a);
}

DEF: float32x4
vfmsq_f32 (float32x4_t a, float32x4_t b, float32x4_t c)
{
  return builtin_aarch64_fmav4sf (-b, c, a);
}

DEF: float64x2
vfmsq_f64 (float64x2_t a, float64x2_t b, float64x2_t c)
{
  return builtin_aarch64_fmav2df (-b, c, a);
}


DEF: float32x2
vfms_lane_f32 (float32x2_t a, float32x2_t b,
        float32x2_t c, const int lane)
{
  return builtin_aarch64_fmav2sf (-b,
        vdup_n_f32 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })),
        a);
}

DEF: float64x1
vfms_lane_f64 (float64x1_t a, float64x1_t b,
        float64x1_t c, const int lane)
{
  return (float64x1_t) {builtin_fma (-b[0], c[0], a[0])};
}

DEF: float64
vfmsd_lane_f64 (float64_t a, float64_t b,
         float64x1_t c, const int lane)
{
  return builtin_fma (-b, c[0], a);
}

DEF: float32
vfmss_lane_f32 (float32_t a, float32_t b,
         float32x2_t c, const int lane)
{
  return builtin_fmaf (-b, extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; }), a);
}

DEF: float32x2
vfms_laneq_f32 (float32x2_t a, float32x2_t b,
         float32x4_t c, const int lane)
{
  return builtin_aarch64_fmav2sf (-b,
        vdup_n_f32 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })),
        a);
}

DEF: float64x1
vfms_laneq_f64 (float64x1_t a, float64x1_t b,
         float64x2_t c, const int lane)
{
  float64_t c0 = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; });
  return (float64x1_t) {builtin_fma (-b[0], c0, a[0])};
}

DEF: float64
vfmsd_laneq_f64 (float64_t a, float64_t b,
          float64x2_t c, const int lane)
{
  return builtin_fma (-b, extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; }), a);
}

DEF: float32
vfmss_laneq_f32 (float32_t a, float32_t b,
   float32x4_t c, const int lane)
{
  return builtin_fmaf (-b, extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; }), a);
}

DEF: float32x4
vfmsq_lane_f32 (float32x4_t a, float32x4_t b,
         float32x2_t c, const int lane)
{
  return builtin_aarch64_fmav4sf (-b,
        vdupq_n_f32 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })),
        a);
}

DEF: float64x2
vfmsq_lane_f64 (float64x2_t a, float64x2_t b,
         float64x1_t c, const int lane)
{
  return builtin_aarch64_fmav2df (-b, vdupq_n_f64 (c[0]), a);
}

DEF: float32x4
vfmsq_laneq_f32 (float32x4_t a, float32x4_t b,
          float32x4_t c, const int lane)
{
  return builtin_aarch64_fmav4sf (-b,
        vdupq_n_f32 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })),
        a);
}

DEF: float64x2
vfmsq_laneq_f64 (float64x2_t a, float64x2_t b,
          float64x2_t c, const int lane)
{
  return builtin_aarch64_fmav2df (-b,
        vdupq_n_f64 (extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })),
        a);
}

DEF: float32x2
vld1_f32 (const float32_t *a)
{
  return builtin_aarch64_ld1v2sf ((const builtin_aarch64_simd_sf *) a);
}

DEF: float64x1
vld1_f64 (const float64_t *a)
{
  return (float64x1_t) {*a};
}

DEF: poly8x8
vld1_p8 (const poly8_t *a)
{
  return (poly8x8_t)
    builtin_aarch64_ld1v8qi ((const builtin_aarch64_simd_qi *) a);
}

DEF: poly16x4
vld1_p16 (const poly16_t *a)
{
  return (poly16x4_t)
    builtin_aarch64_ld1v4hi ((const builtin_aarch64_simd_hi *) a);
}

DEF: int8x8
vld1_s8 (const int8_t *a)
{
  return builtin_aarch64_ld1v8qi ((const builtin_aarch64_simd_qi *) a);
}

DEF: int16x4
vld1_s16 (const int16_t *a)
{
  return builtin_aarch64_ld1v4hi ((const builtin_aarch64_simd_hi *) a);
}

DEF: int32x2
vld1_s32 (const int32_t *a)
{
  return builtin_aarch64_ld1v2si ((const builtin_aarch64_simd_si *) a);
}

DEF: int64x1
vld1_s64 (const int64_t *a)
{
  return (int64x1_t) {*a};
}

DEF: uint8x8
vld1_u8 (const uint8_t *a)
{
  return (uint8x8_t)
    builtin_aarch64_ld1v8qi ((const builtin_aarch64_simd_qi *) a);
}

DEF: uint16x4
vld1_u16 (const uint16_t *a)
{
  return (uint16x4_t)
    builtin_aarch64_ld1v4hi ((const builtin_aarch64_simd_hi *) a);
}

DEF: uint32x2
vld1_u32 (const uint32_t *a)
{
  return (uint32x2_t)
    builtin_aarch64_ld1v2si ((const builtin_aarch64_simd_si *) a);
}

DEF: uint64x1
vld1_u64 (const uint64_t *a)
{
  return (uint64x1_t) {*a};
}

DEF: float32x4
vld1q_f32 (const float32_t *a)
{
  return builtin_aarch64_ld1v4sf ((const builtin_aarch64_simd_sf *) a);
}

DEF: float64x2
vld1q_f64 (const float64_t *a)
{
  return builtin_aarch64_ld1v2df ((const builtin_aarch64_simd_df *) a);
}

DEF: poly8x16
vld1q_p8 (const poly8_t *a)
{
  return (poly8x16_t)
    builtin_aarch64_ld1v16qi ((const builtin_aarch64_simd_qi *) a);
}

DEF: poly16x8
vld1q_p16 (const poly16_t *a)
{
  return (poly16x8_t)
    builtin_aarch64_ld1v8hi ((const builtin_aarch64_simd_hi *) a);
}

DEF: int8x16
vld1q_s8 (const int8_t *a)
{
  return builtin_aarch64_ld1v16qi ((const builtin_aarch64_simd_qi *) a);
}

DEF: int16x8
vld1q_s16 (const int16_t *a)
{
  return builtin_aarch64_ld1v8hi ((const builtin_aarch64_simd_hi *) a);
}

DEF: int32x4
vld1q_s32 (const int32_t *a)
{
  return builtin_aarch64_ld1v4si ((const builtin_aarch64_simd_si *) a);
}

DEF: int64x2
vld1q_s64 (const int64_t *a)
{
  return builtin_aarch64_ld1v2di ((const builtin_aarch64_simd_di *) a);
}

DEF: uint8x16
vld1q_u8 (const uint8_t *a)
{
  return (uint8x16_t)
    builtin_aarch64_ld1v16qi ((const builtin_aarch64_simd_qi *) a);
}

DEF: uint16x8
vld1q_u16 (const uint16_t *a)
{
  return (uint16x8_t)
    builtin_aarch64_ld1v8hi ((const builtin_aarch64_simd_hi *) a);
}

DEF: uint32x4
vld1q_u32 (const uint32_t *a)
{
  return (uint32x4_t)
    builtin_aarch64_ld1v4si ((const builtin_aarch64_simd_si *) a);
}

DEF: uint64x2
vld1q_u64 (const uint64_t *a)
{
  return (uint64x2_t)
    builtin_aarch64_ld1v2di ((const builtin_aarch64_simd_di *) a);
}

DEF: float32x2
vld1_dup_f32 (const float32_t* a)
{
  return vdup_n_f32 (*a);
}

DEF: float64x1
vld1_dup_f64 (const float64_t* a)
{
  return vdup_n_f64 (*a);
}

DEF: poly8x8
vld1_dup_p8 (const poly8_t* a)
{
  return vdup_n_p8 (*a);
}

DEF: poly16x4
vld1_dup_p16 (const poly16_t* a)
{
  return vdup_n_p16 (*a);
}

DEF: int8x8
vld1_dup_s8 (const int8_t* a)
{
  return vdup_n_s8 (*a);
}

DEF: int16x4
vld1_dup_s16 (const int16_t* a)
{
  return vdup_n_s16 (*a);
}

DEF: int32x2
vld1_dup_s32 (const int32_t* a)
{
  return vdup_n_s32 (*a);
}

DEF: int64x1
vld1_dup_s64 (const int64_t* a)
{
  return vdup_n_s64 (*a);
}

DEF: uint8x8
vld1_dup_u8 (const uint8_t* a)
{
  return vdup_n_u8 (*a);
}

DEF: uint16x4
vld1_dup_u16 (const uint16_t* a)
{
  return vdup_n_u16 (*a);
}

DEF: uint32x2
vld1_dup_u32 (const uint32_t* a)
{
  return vdup_n_u32 (*a);
}

DEF: uint64x1
vld1_dup_u64 (const uint64_t* a)
{
  return vdup_n_u64 (*a);
}

DEF: float32x4
vld1q_dup_f32 (const float32_t* a)
{
  return vdupq_n_f32 (*a);
}

DEF: float64x2
vld1q_dup_f64 (const float64_t* a)
{
  return vdupq_n_f64 (*a);
}

DEF: poly8x16
vld1q_dup_p8 (const poly8_t* a)
{
  return vdupq_n_p8 (*a);
}

DEF: poly16x8
vld1q_dup_p16 (const poly16_t* a)
{
  return vdupq_n_p16 (*a);
}

DEF: int8x16
vld1q_dup_s8 (const int8_t* a)
{
  return vdupq_n_s8 (*a);
}

DEF: int16x8
vld1q_dup_s16 (const int16_t* a)
{
  return vdupq_n_s16 (*a);
}

DEF: int32x4
vld1q_dup_s32 (const int32_t* a)
{
  return vdupq_n_s32 (*a);
}

DEF: int64x2
vld1q_dup_s64 (const int64_t* a)
{
  return vdupq_n_s64 (*a);
}

DEF: uint8x16
vld1q_dup_u8 (const uint8_t* a)
{
  return vdupq_n_u8 (*a);
}

DEF: uint16x8
vld1q_dup_u16 (const uint16_t* a)
{
  return vdupq_n_u16 (*a);
}

DEF: uint32x4
vld1q_dup_u32 (const uint32_t* a)
{
  return vdupq_n_u32 (*a);
}

DEF: uint64x2
vld1q_dup_u64 (const uint64_t* a)
{
  return vdupq_n_u64 (*a);
}

DEF: float32x2
vld1_lane_f32 (const float32_t *src, float32x2_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: float64x1
vld1_lane_f64 (const float64_t *src, float64x1_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: poly8x8
vld1_lane_p8 (const poly8_t *src, poly8x8_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: poly16x4
vld1_lane_p16 (const poly16_t *src, poly16x4_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: int8x8
vld1_lane_s8 (const int8_t *src, int8x8_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: int16x4
vld1_lane_s16 (const int16_t *src, int16x4_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: int32x2
vld1_lane_s32 (const int32_t *src, int32x2_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: int64x1
vld1_lane_s64 (const int64_t *src, int64x1_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: uint8x8
vld1_lane_u8 (const uint8_t *src, uint8x8_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: uint16x4
vld1_lane_u16 (const uint16_t *src, uint16x4_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: uint32x2
vld1_lane_u32 (const uint32_t *src, uint32x2_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: uint64x1
vld1_lane_u64 (const uint64_t *src, uint64x1_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: float32x4
vld1q_lane_f32 (const float32_t *src, float32x4_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: float64x2
vld1q_lane_f64 (const float64_t *src, float64x2_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: poly8x16
vld1q_lane_p8 (const poly8_t *src, poly8x16_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: poly16x8
vld1q_lane_p16 (const poly16_t *src, poly16x8_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: int8x16
vld1q_lane_s8 (const int8_t *src, int8x16_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: int16x8
vld1q_lane_s16 (const int16_t *src, int16x8_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: int32x4
vld1q_lane_s32 (const int32_t *src, int32x4_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: int64x2
vld1q_lane_s64 (const int64_t *src, int64x2_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: uint8x16
vld1q_lane_u8 (const uint8_t *src, uint8x16_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: uint16x8
vld1q_lane_u16 (const uint16_t *src, uint16x8_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: uint32x4
vld1q_lane_u32 (const uint32_t *src, uint32x4_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: uint64x2
vld1q_lane_u64 (const uint64_t *src, uint64x2_t vec, const int lane)
{
  return extension ({ builtin_aarch64_im_lane_boundsi (sizeof(vec), sizeof(vec[0]), lane); vec[lane] = *src; vec; });
}

DEF: int64x1x2
vld2_s64 (const int64_t * a)
{
  int64x1x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2di ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (int64x1_t) builtin_aarch64_get_dregoidi (o, 0);
  ret.val[1] = (int64x1_t) builtin_aarch64_get_dregoidi (o, 1);
  return ret;
}

DEF: uint64x1x2
vld2_u64 (const uint64_t * a)
{
  uint64x1x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2di ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (uint64x1_t) builtin_aarch64_get_dregoidi (o, 0);
  ret.val[1] = (uint64x1_t) builtin_aarch64_get_dregoidi (o, 1);
  return ret;
}

DEF: float64x1x2
vld2_f64 (const float64_t * a)
{
  float64x1x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2df ((const builtin_aarch64_simd_df *) a);
  ret.val[0] = (float64x1_t) {builtin_aarch64_get_dregoidf (o, 0)};
  ret.val[1] = (float64x1_t) {builtin_aarch64_get_dregoidf (o, 1)};
  return ret;
}

DEF: int8x8x2
vld2_s8 (const int8_t * a)
{
  int8x8x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2v8qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (int8x8_t) builtin_aarch64_get_dregoiv8qi (o, 0);
  ret.val[1] = (int8x8_t) builtin_aarch64_get_dregoiv8qi (o, 1);
  return ret;
}

DEF: poly8x8x2
vld2_p8 (const poly8_t * a)
{
  poly8x8x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2v8qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (poly8x8_t) builtin_aarch64_get_dregoiv8qi (o, 0);
  ret.val[1] = (poly8x8_t) builtin_aarch64_get_dregoiv8qi (o, 1);
  return ret;
}

DEF: int16x4x2
vld2_s16 (const int16_t * a)
{
  int16x4x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2v4hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (int16x4_t) builtin_aarch64_get_dregoiv4hi (o, 0);
  ret.val[1] = (int16x4_t) builtin_aarch64_get_dregoiv4hi (o, 1);
  return ret;
}

DEF: poly16x4x2
vld2_p16 (const poly16_t * a)
{
  poly16x4x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2v4hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (poly16x4_t) builtin_aarch64_get_dregoiv4hi (o, 0);
  ret.val[1] = (poly16x4_t) builtin_aarch64_get_dregoiv4hi (o, 1);
  return ret;
}

DEF: int32x2x2
vld2_s32 (const int32_t * a)
{
  int32x2x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2v2si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (int32x2_t) builtin_aarch64_get_dregoiv2si (o, 0);
  ret.val[1] = (int32x2_t) builtin_aarch64_get_dregoiv2si (o, 1);
  return ret;
}

DEF: uint8x8x2
vld2_u8 (const uint8_t * a)
{
  uint8x8x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2v8qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (uint8x8_t) builtin_aarch64_get_dregoiv8qi (o, 0);
  ret.val[1] = (uint8x8_t) builtin_aarch64_get_dregoiv8qi (o, 1);
  return ret;
}

DEF: uint16x4x2
vld2_u16 (const uint16_t * a)
{
  uint16x4x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2v4hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (uint16x4_t) builtin_aarch64_get_dregoiv4hi (o, 0);
  ret.val[1] = (uint16x4_t) builtin_aarch64_get_dregoiv4hi (o, 1);
  return ret;
}

DEF: uint32x2x2
vld2_u32 (const uint32_t * a)
{
  uint32x2x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2v2si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (uint32x2_t) builtin_aarch64_get_dregoiv2si (o, 0);
  ret.val[1] = (uint32x2_t) builtin_aarch64_get_dregoiv2si (o, 1);
  return ret;
}

DEF: float32x2x2
vld2_f32 (const float32_t * a)
{
  float32x2x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2v2sf ((const builtin_aarch64_simd_sf *) a);
  ret.val[0] = (float32x2_t) builtin_aarch64_get_dregoiv2sf (o, 0);
  ret.val[1] = (float32x2_t) builtin_aarch64_get_dregoiv2sf (o, 1);
  return ret;
}

DEF: int8x16x2
vld2q_s8 (const int8_t * a)
{
  int8x16x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2v16qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (int8x16_t) builtin_aarch64_get_qregoiv16qi (o, 0);
  ret.val[1] = (int8x16_t) builtin_aarch64_get_qregoiv16qi (o, 1);
  return ret;
}

DEF: poly8x16x2
vld2q_p8 (const poly8_t * a)
{
  poly8x16x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2v16qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (poly8x16_t) builtin_aarch64_get_qregoiv16qi (o, 0);
  ret.val[1] = (poly8x16_t) builtin_aarch64_get_qregoiv16qi (o, 1);
  return ret;
}

DEF: int16x8x2
vld2q_s16 (const int16_t * a)
{
  int16x8x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2v8hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (int16x8_t) builtin_aarch64_get_qregoiv8hi (o, 0);
  ret.val[1] = (int16x8_t) builtin_aarch64_get_qregoiv8hi (o, 1);
  return ret;
}

DEF: poly16x8x2
vld2q_p16 (const poly16_t * a)
{
  poly16x8x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2v8hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (poly16x8_t) builtin_aarch64_get_qregoiv8hi (o, 0);
  ret.val[1] = (poly16x8_t) builtin_aarch64_get_qregoiv8hi (o, 1);
  return ret;
}

DEF: int32x4x2
vld2q_s32 (const int32_t * a)
{
  int32x4x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2v4si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (int32x4_t) builtin_aarch64_get_qregoiv4si (o, 0);
  ret.val[1] = (int32x4_t) builtin_aarch64_get_qregoiv4si (o, 1);
  return ret;
}

DEF: int64x2x2
vld2q_s64 (const int64_t * a)
{
  int64x2x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2v2di ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (int64x2_t) builtin_aarch64_get_qregoiv2di (o, 0);
  ret.val[1] = (int64x2_t) builtin_aarch64_get_qregoiv2di (o, 1);
  return ret;
}

DEF: uint8x16x2
vld2q_u8 (const uint8_t * a)
{
  uint8x16x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2v16qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (uint8x16_t) builtin_aarch64_get_qregoiv16qi (o, 0);
  ret.val[1] = (uint8x16_t) builtin_aarch64_get_qregoiv16qi (o, 1);
  return ret;
}

DEF: uint16x8x2
vld2q_u16 (const uint16_t * a)
{
  uint16x8x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2v8hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (uint16x8_t) builtin_aarch64_get_qregoiv8hi (o, 0);
  ret.val[1] = (uint16x8_t) builtin_aarch64_get_qregoiv8hi (o, 1);
  return ret;
}

DEF: uint32x4x2
vld2q_u32 (const uint32_t * a)
{
  uint32x4x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2v4si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (uint32x4_t) builtin_aarch64_get_qregoiv4si (o, 0);
  ret.val[1] = (uint32x4_t) builtin_aarch64_get_qregoiv4si (o, 1);
  return ret;
}

DEF: uint64x2x2
vld2q_u64 (const uint64_t * a)
{
  uint64x2x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2v2di ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (uint64x2_t) builtin_aarch64_get_qregoiv2di (o, 0);
  ret.val[1] = (uint64x2_t) builtin_aarch64_get_qregoiv2di (o, 1);
  return ret;
}

DEF: float32x4x2
vld2q_f32 (const float32_t * a)
{
  float32x4x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2v4sf ((const builtin_aarch64_simd_sf *) a);
  ret.val[0] = (float32x4_t) builtin_aarch64_get_qregoiv4sf (o, 0);
  ret.val[1] = (float32x4_t) builtin_aarch64_get_qregoiv4sf (o, 1);
  return ret;
}

DEF: float64x2x2
vld2q_f64 (const float64_t * a)
{
  float64x2x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2v2df ((const builtin_aarch64_simd_df *) a);
  ret.val[0] = (float64x2_t) builtin_aarch64_get_qregoiv2df (o, 0);
  ret.val[1] = (float64x2_t) builtin_aarch64_get_qregoiv2df (o, 1);
  return ret;
}

DEF: int64x1x3
vld3_s64 (const int64_t * a)
{
  int64x1x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3di ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (int64x1_t) builtin_aarch64_get_dregcidi (o, 0);
  ret.val[1] = (int64x1_t) builtin_aarch64_get_dregcidi (o, 1);
  ret.val[2] = (int64x1_t) builtin_aarch64_get_dregcidi (o, 2);
  return ret;
}

DEF: uint64x1x3
vld3_u64 (const uint64_t * a)
{
  uint64x1x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3di ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (uint64x1_t) builtin_aarch64_get_dregcidi (o, 0);
  ret.val[1] = (uint64x1_t) builtin_aarch64_get_dregcidi (o, 1);
  ret.val[2] = (uint64x1_t) builtin_aarch64_get_dregcidi (o, 2);
  return ret;
}

DEF: float64x1x3
vld3_f64 (const float64_t * a)
{
  float64x1x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3df ((const builtin_aarch64_simd_df *) a);
  ret.val[0] = (float64x1_t) {builtin_aarch64_get_dregcidf (o, 0)};
  ret.val[1] = (float64x1_t) {builtin_aarch64_get_dregcidf (o, 1)};
  ret.val[2] = (float64x1_t) {builtin_aarch64_get_dregcidf (o, 2)};
  return ret;
}

DEF: int8x8x3
vld3_s8 (const int8_t * a)
{
  int8x8x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3v8qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (int8x8_t) builtin_aarch64_get_dregciv8qi (o, 0);
  ret.val[1] = (int8x8_t) builtin_aarch64_get_dregciv8qi (o, 1);
  ret.val[2] = (int8x8_t) builtin_aarch64_get_dregciv8qi (o, 2);
  return ret;
}

DEF: poly8x8x3
vld3_p8 (const poly8_t * a)
{
  poly8x8x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3v8qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (poly8x8_t) builtin_aarch64_get_dregciv8qi (o, 0);
  ret.val[1] = (poly8x8_t) builtin_aarch64_get_dregciv8qi (o, 1);
  ret.val[2] = (poly8x8_t) builtin_aarch64_get_dregciv8qi (o, 2);
  return ret;
}

DEF: int16x4x3
vld3_s16 (const int16_t * a)
{
  int16x4x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3v4hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (int16x4_t) builtin_aarch64_get_dregciv4hi (o, 0);
  ret.val[1] = (int16x4_t) builtin_aarch64_get_dregciv4hi (o, 1);
  ret.val[2] = (int16x4_t) builtin_aarch64_get_dregciv4hi (o, 2);
  return ret;
}

DEF: poly16x4x3
vld3_p16 (const poly16_t * a)
{
  poly16x4x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3v4hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (poly16x4_t) builtin_aarch64_get_dregciv4hi (o, 0);
  ret.val[1] = (poly16x4_t) builtin_aarch64_get_dregciv4hi (o, 1);
  ret.val[2] = (poly16x4_t) builtin_aarch64_get_dregciv4hi (o, 2);
  return ret;
}

DEF: int32x2x3
vld3_s32 (const int32_t * a)
{
  int32x2x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3v2si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (int32x2_t) builtin_aarch64_get_dregciv2si (o, 0);
  ret.val[1] = (int32x2_t) builtin_aarch64_get_dregciv2si (o, 1);
  ret.val[2] = (int32x2_t) builtin_aarch64_get_dregciv2si (o, 2);
  return ret;
}

DEF: uint8x8x3
vld3_u8 (const uint8_t * a)
{
  uint8x8x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3v8qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (uint8x8_t) builtin_aarch64_get_dregciv8qi (o, 0);
  ret.val[1] = (uint8x8_t) builtin_aarch64_get_dregciv8qi (o, 1);
  ret.val[2] = (uint8x8_t) builtin_aarch64_get_dregciv8qi (o, 2);
  return ret;
}

DEF: uint16x4x3
vld3_u16 (const uint16_t * a)
{
  uint16x4x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3v4hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (uint16x4_t) builtin_aarch64_get_dregciv4hi (o, 0);
  ret.val[1] = (uint16x4_t) builtin_aarch64_get_dregciv4hi (o, 1);
  ret.val[2] = (uint16x4_t) builtin_aarch64_get_dregciv4hi (o, 2);
  return ret;
}

DEF: uint32x2x3
vld3_u32 (const uint32_t * a)
{
  uint32x2x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3v2si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (uint32x2_t) builtin_aarch64_get_dregciv2si (o, 0);
  ret.val[1] = (uint32x2_t) builtin_aarch64_get_dregciv2si (o, 1);
  ret.val[2] = (uint32x2_t) builtin_aarch64_get_dregciv2si (o, 2);
  return ret;
}

DEF: float32x2x3
vld3_f32 (const float32_t * a)
{
  float32x2x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3v2sf ((const builtin_aarch64_simd_sf *) a);
  ret.val[0] = (float32x2_t) builtin_aarch64_get_dregciv2sf (o, 0);
  ret.val[1] = (float32x2_t) builtin_aarch64_get_dregciv2sf (o, 1);
  ret.val[2] = (float32x2_t) builtin_aarch64_get_dregciv2sf (o, 2);
  return ret;
}

DEF: int8x16x3
vld3q_s8 (const int8_t * a)
{
  int8x16x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3v16qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (int8x16_t) builtin_aarch64_get_qregciv16qi (o, 0);
  ret.val[1] = (int8x16_t) builtin_aarch64_get_qregciv16qi (o, 1);
  ret.val[2] = (int8x16_t) builtin_aarch64_get_qregciv16qi (o, 2);
  return ret;
}

DEF: poly8x16x3
vld3q_p8 (const poly8_t * a)
{
  poly8x16x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3v16qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (poly8x16_t) builtin_aarch64_get_qregciv16qi (o, 0);
  ret.val[1] = (poly8x16_t) builtin_aarch64_get_qregciv16qi (o, 1);
  ret.val[2] = (poly8x16_t) builtin_aarch64_get_qregciv16qi (o, 2);
  return ret;
}

DEF: int16x8x3
vld3q_s16 (const int16_t * a)
{
  int16x8x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3v8hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (int16x8_t) builtin_aarch64_get_qregciv8hi (o, 0);
  ret.val[1] = (int16x8_t) builtin_aarch64_get_qregciv8hi (o, 1);
  ret.val[2] = (int16x8_t) builtin_aarch64_get_qregciv8hi (o, 2);
  return ret;
}

DEF: poly16x8x3
vld3q_p16 (const poly16_t * a)
{
  poly16x8x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3v8hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (poly16x8_t) builtin_aarch64_get_qregciv8hi (o, 0);
  ret.val[1] = (poly16x8_t) builtin_aarch64_get_qregciv8hi (o, 1);
  ret.val[2] = (poly16x8_t) builtin_aarch64_get_qregciv8hi (o, 2);
  return ret;
}

DEF: int32x4x3
vld3q_s32 (const int32_t * a)
{
  int32x4x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3v4si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (int32x4_t) builtin_aarch64_get_qregciv4si (o, 0);
  ret.val[1] = (int32x4_t) builtin_aarch64_get_qregciv4si (o, 1);
  ret.val[2] = (int32x4_t) builtin_aarch64_get_qregciv4si (o, 2);
  return ret;
}

DEF: int64x2x3
vld3q_s64 (const int64_t * a)
{
  int64x2x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3v2di ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (int64x2_t) builtin_aarch64_get_qregciv2di (o, 0);
  ret.val[1] = (int64x2_t) builtin_aarch64_get_qregciv2di (o, 1);
  ret.val[2] = (int64x2_t) builtin_aarch64_get_qregciv2di (o, 2);
  return ret;
}

DEF: uint8x16x3
vld3q_u8 (const uint8_t * a)
{
  uint8x16x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3v16qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (uint8x16_t) builtin_aarch64_get_qregciv16qi (o, 0);
  ret.val[1] = (uint8x16_t) builtin_aarch64_get_qregciv16qi (o, 1);
  ret.val[2] = (uint8x16_t) builtin_aarch64_get_qregciv16qi (o, 2);
  return ret;
}

DEF: uint16x8x3
vld3q_u16 (const uint16_t * a)
{
  uint16x8x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3v8hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (uint16x8_t) builtin_aarch64_get_qregciv8hi (o, 0);
  ret.val[1] = (uint16x8_t) builtin_aarch64_get_qregciv8hi (o, 1);
  ret.val[2] = (uint16x8_t) builtin_aarch64_get_qregciv8hi (o, 2);
  return ret;
}

DEF: uint32x4x3
vld3q_u32 (const uint32_t * a)
{
  uint32x4x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3v4si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (uint32x4_t) builtin_aarch64_get_qregciv4si (o, 0);
  ret.val[1] = (uint32x4_t) builtin_aarch64_get_qregciv4si (o, 1);
  ret.val[2] = (uint32x4_t) builtin_aarch64_get_qregciv4si (o, 2);
  return ret;
}

DEF: uint64x2x3
vld3q_u64 (const uint64_t * a)
{
  uint64x2x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3v2di ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (uint64x2_t) builtin_aarch64_get_qregciv2di (o, 0);
  ret.val[1] = (uint64x2_t) builtin_aarch64_get_qregciv2di (o, 1);
  ret.val[2] = (uint64x2_t) builtin_aarch64_get_qregciv2di (o, 2);
  return ret;
}

DEF: float32x4x3
vld3q_f32 (const float32_t * a)
{
  float32x4x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3v4sf ((const builtin_aarch64_simd_sf *) a);
  ret.val[0] = (float32x4_t) builtin_aarch64_get_qregciv4sf (o, 0);
  ret.val[1] = (float32x4_t) builtin_aarch64_get_qregciv4sf (o, 1);
  ret.val[2] = (float32x4_t) builtin_aarch64_get_qregciv4sf (o, 2);
  return ret;
}

DEF: float64x2x3
vld3q_f64 (const float64_t * a)
{
  float64x2x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3v2df ((const builtin_aarch64_simd_df *) a);
  ret.val[0] = (float64x2_t) builtin_aarch64_get_qregciv2df (o, 0);
  ret.val[1] = (float64x2_t) builtin_aarch64_get_qregciv2df (o, 1);
  ret.val[2] = (float64x2_t) builtin_aarch64_get_qregciv2df (o, 2);
  return ret;
}

DEF: int64x1x4
vld4_s64 (const int64_t * a)
{
  int64x1x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4di ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (int64x1_t) builtin_aarch64_get_dregxidi (o, 0);
  ret.val[1] = (int64x1_t) builtin_aarch64_get_dregxidi (o, 1);
  ret.val[2] = (int64x1_t) builtin_aarch64_get_dregxidi (o, 2);
  ret.val[3] = (int64x1_t) builtin_aarch64_get_dregxidi (o, 3);
  return ret;
}

DEF: uint64x1x4
vld4_u64 (const uint64_t * a)
{
  uint64x1x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4di ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (uint64x1_t) builtin_aarch64_get_dregxidi (o, 0);
  ret.val[1] = (uint64x1_t) builtin_aarch64_get_dregxidi (o, 1);
  ret.val[2] = (uint64x1_t) builtin_aarch64_get_dregxidi (o, 2);
  ret.val[3] = (uint64x1_t) builtin_aarch64_get_dregxidi (o, 3);
  return ret;
}

DEF: float64x1x4
vld4_f64 (const float64_t * a)
{
  float64x1x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4df ((const builtin_aarch64_simd_df *) a);
  ret.val[0] = (float64x1_t) {builtin_aarch64_get_dregxidf (o, 0)};
  ret.val[1] = (float64x1_t) {builtin_aarch64_get_dregxidf (o, 1)};
  ret.val[2] = (float64x1_t) {builtin_aarch64_get_dregxidf (o, 2)};
  ret.val[3] = (float64x1_t) {builtin_aarch64_get_dregxidf (o, 3)};
  return ret;
}

DEF: int8x8x4
vld4_s8 (const int8_t * a)
{
  int8x8x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4v8qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (int8x8_t) builtin_aarch64_get_dregxiv8qi (o, 0);
  ret.val[1] = (int8x8_t) builtin_aarch64_get_dregxiv8qi (o, 1);
  ret.val[2] = (int8x8_t) builtin_aarch64_get_dregxiv8qi (o, 2);
  ret.val[3] = (int8x8_t) builtin_aarch64_get_dregxiv8qi (o, 3);
  return ret;
}

DEF: poly8x8x4
vld4_p8 (const poly8_t * a)
{
  poly8x8x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4v8qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (poly8x8_t) builtin_aarch64_get_dregxiv8qi (o, 0);
  ret.val[1] = (poly8x8_t) builtin_aarch64_get_dregxiv8qi (o, 1);
  ret.val[2] = (poly8x8_t) builtin_aarch64_get_dregxiv8qi (o, 2);
  ret.val[3] = (poly8x8_t) builtin_aarch64_get_dregxiv8qi (o, 3);
  return ret;
}

DEF: int16x4x4
vld4_s16 (const int16_t * a)
{
  int16x4x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4v4hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (int16x4_t) builtin_aarch64_get_dregxiv4hi (o, 0);
  ret.val[1] = (int16x4_t) builtin_aarch64_get_dregxiv4hi (o, 1);
  ret.val[2] = (int16x4_t) builtin_aarch64_get_dregxiv4hi (o, 2);
  ret.val[3] = (int16x4_t) builtin_aarch64_get_dregxiv4hi (o, 3);
  return ret;
}

DEF: poly16x4x4
vld4_p16 (const poly16_t * a)
{
  poly16x4x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4v4hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (poly16x4_t) builtin_aarch64_get_dregxiv4hi (o, 0);
  ret.val[1] = (poly16x4_t) builtin_aarch64_get_dregxiv4hi (o, 1);
  ret.val[2] = (poly16x4_t) builtin_aarch64_get_dregxiv4hi (o, 2);
  ret.val[3] = (poly16x4_t) builtin_aarch64_get_dregxiv4hi (o, 3);
  return ret;
}

DEF: int32x2x4
vld4_s32 (const int32_t * a)
{
  int32x2x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4v2si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (int32x2_t) builtin_aarch64_get_dregxiv2si (o, 0);
  ret.val[1] = (int32x2_t) builtin_aarch64_get_dregxiv2si (o, 1);
  ret.val[2] = (int32x2_t) builtin_aarch64_get_dregxiv2si (o, 2);
  ret.val[3] = (int32x2_t) builtin_aarch64_get_dregxiv2si (o, 3);
  return ret;
}

DEF: uint8x8x4
vld4_u8 (const uint8_t * a)
{
  uint8x8x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4v8qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (uint8x8_t) builtin_aarch64_get_dregxiv8qi (o, 0);
  ret.val[1] = (uint8x8_t) builtin_aarch64_get_dregxiv8qi (o, 1);
  ret.val[2] = (uint8x8_t) builtin_aarch64_get_dregxiv8qi (o, 2);
  ret.val[3] = (uint8x8_t) builtin_aarch64_get_dregxiv8qi (o, 3);
  return ret;
}

DEF: uint16x4x4
vld4_u16 (const uint16_t * a)
{
  uint16x4x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4v4hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (uint16x4_t) builtin_aarch64_get_dregxiv4hi (o, 0);
  ret.val[1] = (uint16x4_t) builtin_aarch64_get_dregxiv4hi (o, 1);
  ret.val[2] = (uint16x4_t) builtin_aarch64_get_dregxiv4hi (o, 2);
  ret.val[3] = (uint16x4_t) builtin_aarch64_get_dregxiv4hi (o, 3);
  return ret;
}

DEF: uint32x2x4
vld4_u32 (const uint32_t * a)
{
  uint32x2x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4v2si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (uint32x2_t) builtin_aarch64_get_dregxiv2si (o, 0);
  ret.val[1] = (uint32x2_t) builtin_aarch64_get_dregxiv2si (o, 1);
  ret.val[2] = (uint32x2_t) builtin_aarch64_get_dregxiv2si (o, 2);
  ret.val[3] = (uint32x2_t) builtin_aarch64_get_dregxiv2si (o, 3);
  return ret;
}

DEF: float32x2x4
vld4_f32 (const float32_t * a)
{
  float32x2x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4v2sf ((const builtin_aarch64_simd_sf *) a);
  ret.val[0] = (float32x2_t) builtin_aarch64_get_dregxiv2sf (o, 0);
  ret.val[1] = (float32x2_t) builtin_aarch64_get_dregxiv2sf (o, 1);
  ret.val[2] = (float32x2_t) builtin_aarch64_get_dregxiv2sf (o, 2);
  ret.val[3] = (float32x2_t) builtin_aarch64_get_dregxiv2sf (o, 3);
  return ret;
}

DEF: int8x16x4
vld4q_s8 (const int8_t * a)
{
  int8x16x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4v16qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (int8x16_t) builtin_aarch64_get_qregxiv16qi (o, 0);
  ret.val[1] = (int8x16_t) builtin_aarch64_get_qregxiv16qi (o, 1);
  ret.val[2] = (int8x16_t) builtin_aarch64_get_qregxiv16qi (o, 2);
  ret.val[3] = (int8x16_t) builtin_aarch64_get_qregxiv16qi (o, 3);
  return ret;
}

DEF: poly8x16x4
vld4q_p8 (const poly8_t * a)
{
  poly8x16x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4v16qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (poly8x16_t) builtin_aarch64_get_qregxiv16qi (o, 0);
  ret.val[1] = (poly8x16_t) builtin_aarch64_get_qregxiv16qi (o, 1);
  ret.val[2] = (poly8x16_t) builtin_aarch64_get_qregxiv16qi (o, 2);
  ret.val[3] = (poly8x16_t) builtin_aarch64_get_qregxiv16qi (o, 3);
  return ret;
}

DEF: int16x8x4
vld4q_s16 (const int16_t * a)
{
  int16x8x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4v8hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (int16x8_t) builtin_aarch64_get_qregxiv8hi (o, 0);
  ret.val[1] = (int16x8_t) builtin_aarch64_get_qregxiv8hi (o, 1);
  ret.val[2] = (int16x8_t) builtin_aarch64_get_qregxiv8hi (o, 2);
  ret.val[3] = (int16x8_t) builtin_aarch64_get_qregxiv8hi (o, 3);
  return ret;
}

DEF: poly16x8x4
vld4q_p16 (const poly16_t * a)
{
  poly16x8x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4v8hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (poly16x8_t) builtin_aarch64_get_qregxiv8hi (o, 0);
  ret.val[1] = (poly16x8_t) builtin_aarch64_get_qregxiv8hi (o, 1);
  ret.val[2] = (poly16x8_t) builtin_aarch64_get_qregxiv8hi (o, 2);
  ret.val[3] = (poly16x8_t) builtin_aarch64_get_qregxiv8hi (o, 3);
  return ret;
}

DEF: int32x4x4
vld4q_s32 (const int32_t * a)
{
  int32x4x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4v4si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (int32x4_t) builtin_aarch64_get_qregxiv4si (o, 0);
  ret.val[1] = (int32x4_t) builtin_aarch64_get_qregxiv4si (o, 1);
  ret.val[2] = (int32x4_t) builtin_aarch64_get_qregxiv4si (o, 2);
  ret.val[3] = (int32x4_t) builtin_aarch64_get_qregxiv4si (o, 3);
  return ret;
}

DEF: int64x2x4
vld4q_s64 (const int64_t * a)
{
  int64x2x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4v2di ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (int64x2_t) builtin_aarch64_get_qregxiv2di (o, 0);
  ret.val[1] = (int64x2_t) builtin_aarch64_get_qregxiv2di (o, 1);
  ret.val[2] = (int64x2_t) builtin_aarch64_get_qregxiv2di (o, 2);
  ret.val[3] = (int64x2_t) builtin_aarch64_get_qregxiv2di (o, 3);
  return ret;
}

DEF: uint8x16x4
vld4q_u8 (const uint8_t * a)
{
  uint8x16x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4v16qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (uint8x16_t) builtin_aarch64_get_qregxiv16qi (o, 0);
  ret.val[1] = (uint8x16_t) builtin_aarch64_get_qregxiv16qi (o, 1);
  ret.val[2] = (uint8x16_t) builtin_aarch64_get_qregxiv16qi (o, 2);
  ret.val[3] = (uint8x16_t) builtin_aarch64_get_qregxiv16qi (o, 3);
  return ret;
}

DEF: uint16x8x4
vld4q_u16 (const uint16_t * a)
{
  uint16x8x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4v8hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (uint16x8_t) builtin_aarch64_get_qregxiv8hi (o, 0);
  ret.val[1] = (uint16x8_t) builtin_aarch64_get_qregxiv8hi (o, 1);
  ret.val[2] = (uint16x8_t) builtin_aarch64_get_qregxiv8hi (o, 2);
  ret.val[3] = (uint16x8_t) builtin_aarch64_get_qregxiv8hi (o, 3);
  return ret;
}

DEF: uint32x4x4
vld4q_u32 (const uint32_t * a)
{
  uint32x4x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4v4si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (uint32x4_t) builtin_aarch64_get_qregxiv4si (o, 0);
  ret.val[1] = (uint32x4_t) builtin_aarch64_get_qregxiv4si (o, 1);
  ret.val[2] = (uint32x4_t) builtin_aarch64_get_qregxiv4si (o, 2);
  ret.val[3] = (uint32x4_t) builtin_aarch64_get_qregxiv4si (o, 3);
  return ret;
}

DEF: uint64x2x4
vld4q_u64 (const uint64_t * a)
{
  uint64x2x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4v2di ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (uint64x2_t) builtin_aarch64_get_qregxiv2di (o, 0);
  ret.val[1] = (uint64x2_t) builtin_aarch64_get_qregxiv2di (o, 1);
  ret.val[2] = (uint64x2_t) builtin_aarch64_get_qregxiv2di (o, 2);
  ret.val[3] = (uint64x2_t) builtin_aarch64_get_qregxiv2di (o, 3);
  return ret;
}

DEF: float32x4x4
vld4q_f32 (const float32_t * a)
{
  float32x4x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4v4sf ((const builtin_aarch64_simd_sf *) a);
  ret.val[0] = (float32x4_t) builtin_aarch64_get_qregxiv4sf (o, 0);
  ret.val[1] = (float32x4_t) builtin_aarch64_get_qregxiv4sf (o, 1);
  ret.val[2] = (float32x4_t) builtin_aarch64_get_qregxiv4sf (o, 2);
  ret.val[3] = (float32x4_t) builtin_aarch64_get_qregxiv4sf (o, 3);
  return ret;
}

DEF: float64x2x4
vld4q_f64 (const float64_t * a)
{
  float64x2x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4v2df ((const builtin_aarch64_simd_df *) a);
  ret.val[0] = (float64x2_t) builtin_aarch64_get_qregxiv2df (o, 0);
  ret.val[1] = (float64x2_t) builtin_aarch64_get_qregxiv2df (o, 1);
  ret.val[2] = (float64x2_t) builtin_aarch64_get_qregxiv2df (o, 2);
  ret.val[3] = (float64x2_t) builtin_aarch64_get_qregxiv2df (o, 3);
  return ret;
}

DEF: int8x8x2
vld2_dup_s8 (const int8_t * a)
{
  int8x8x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rv8qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (int8x8_t) builtin_aarch64_get_dregoiv8qi (o, 0);
  ret.val[1] = (int8x8_t) builtin_aarch64_get_dregoiv8qi (o, 1);
  return ret;
}

DEF: int16x4x2
vld2_dup_s16 (const int16_t * a)
{
  int16x4x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rv4hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (int16x4_t) builtin_aarch64_get_dregoiv4hi (o, 0);
  ret.val[1] = (int16x4_t) builtin_aarch64_get_dregoiv4hi (o, 1);
  return ret;
}

DEF: int32x2x2
vld2_dup_s32 (const int32_t * a)
{
  int32x2x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rv2si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (int32x2_t) builtin_aarch64_get_dregoiv2si (o, 0);
  ret.val[1] = (int32x2_t) builtin_aarch64_get_dregoiv2si (o, 1);
  return ret;
}

DEF: float32x2x2
vld2_dup_f32 (const float32_t * a)
{
  float32x2x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rv2sf ((const builtin_aarch64_simd_sf *) a);
  ret.val[0] = (float32x2_t) builtin_aarch64_get_dregoiv2sf (o, 0);
  ret.val[1] = (float32x2_t) builtin_aarch64_get_dregoiv2sf (o, 1);
  return ret;
}

DEF: float64x1x2
vld2_dup_f64 (const float64_t * a)
{
  float64x1x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rdf ((const builtin_aarch64_simd_df *) a);
  ret.val[0] = (float64x1_t) {builtin_aarch64_get_dregoidf (o, 0)};
  ret.val[1] = (float64x1_t) {builtin_aarch64_get_dregoidf (o, 1)};
  return ret;
}

DEF: uint8x8x2
vld2_dup_u8 (const uint8_t * a)
{
  uint8x8x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rv8qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (uint8x8_t) builtin_aarch64_get_dregoiv8qi (o, 0);
  ret.val[1] = (uint8x8_t) builtin_aarch64_get_dregoiv8qi (o, 1);
  return ret;
}

DEF: uint16x4x2
vld2_dup_u16 (const uint16_t * a)
{
  uint16x4x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rv4hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (uint16x4_t) builtin_aarch64_get_dregoiv4hi (o, 0);
  ret.val[1] = (uint16x4_t) builtin_aarch64_get_dregoiv4hi (o, 1);
  return ret;
}

DEF: uint32x2x2
vld2_dup_u32 (const uint32_t * a)
{
  uint32x2x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rv2si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (uint32x2_t) builtin_aarch64_get_dregoiv2si (o, 0);
  ret.val[1] = (uint32x2_t) builtin_aarch64_get_dregoiv2si (o, 1);
  return ret;
}

DEF: poly8x8x2
vld2_dup_p8 (const poly8_t * a)
{
  poly8x8x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rv8qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (poly8x8_t) builtin_aarch64_get_dregoiv8qi (o, 0);
  ret.val[1] = (poly8x8_t) builtin_aarch64_get_dregoiv8qi (o, 1);
  return ret;
}

DEF: poly16x4x2
vld2_dup_p16 (const poly16_t * a)
{
  poly16x4x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rv4hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (poly16x4_t) builtin_aarch64_get_dregoiv4hi (o, 0);
  ret.val[1] = (poly16x4_t) builtin_aarch64_get_dregoiv4hi (o, 1);
  return ret;
}

DEF: int64x1x2
vld2_dup_s64 (const int64_t * a)
{
  int64x1x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rdi ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (int64x1_t) builtin_aarch64_get_dregoidi (o, 0);
  ret.val[1] = (int64x1_t) builtin_aarch64_get_dregoidi (o, 1);
  return ret;
}

DEF: uint64x1x2
vld2_dup_u64 (const uint64_t * a)
{
  uint64x1x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rdi ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (uint64x1_t) builtin_aarch64_get_dregoidi (o, 0);
  ret.val[1] = (uint64x1_t) builtin_aarch64_get_dregoidi (o, 1);
  return ret;
}

DEF: int8x16x2
vld2q_dup_s8 (const int8_t * a)
{
  int8x16x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rv16qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (int8x16_t) builtin_aarch64_get_qregoiv16qi (o, 0);
  ret.val[1] = (int8x16_t) builtin_aarch64_get_qregoiv16qi (o, 1);
  return ret;
}

DEF: poly8x16x2
vld2q_dup_p8 (const poly8_t * a)
{
  poly8x16x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rv16qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (poly8x16_t) builtin_aarch64_get_qregoiv16qi (o, 0);
  ret.val[1] = (poly8x16_t) builtin_aarch64_get_qregoiv16qi (o, 1);
  return ret;
}

DEF: int16x8x2
vld2q_dup_s16 (const int16_t * a)
{
  int16x8x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rv8hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (int16x8_t) builtin_aarch64_get_qregoiv8hi (o, 0);
  ret.val[1] = (int16x8_t) builtin_aarch64_get_qregoiv8hi (o, 1);
  return ret;
}

DEF: poly16x8x2
vld2q_dup_p16 (const poly16_t * a)
{
  poly16x8x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rv8hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (poly16x8_t) builtin_aarch64_get_qregoiv8hi (o, 0);
  ret.val[1] = (poly16x8_t) builtin_aarch64_get_qregoiv8hi (o, 1);
  return ret;
}

DEF: int32x4x2
vld2q_dup_s32 (const int32_t * a)
{
  int32x4x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rv4si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (int32x4_t) builtin_aarch64_get_qregoiv4si (o, 0);
  ret.val[1] = (int32x4_t) builtin_aarch64_get_qregoiv4si (o, 1);
  return ret;
}

DEF: int64x2x2
vld2q_dup_s64 (const int64_t * a)
{
  int64x2x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rv2di ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (int64x2_t) builtin_aarch64_get_qregoiv2di (o, 0);
  ret.val[1] = (int64x2_t) builtin_aarch64_get_qregoiv2di (o, 1);
  return ret;
}

DEF: uint8x16x2
vld2q_dup_u8 (const uint8_t * a)
{
  uint8x16x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rv16qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (uint8x16_t) builtin_aarch64_get_qregoiv16qi (o, 0);
  ret.val[1] = (uint8x16_t) builtin_aarch64_get_qregoiv16qi (o, 1);
  return ret;
}

DEF: uint16x8x2
vld2q_dup_u16 (const uint16_t * a)
{
  uint16x8x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rv8hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (uint16x8_t) builtin_aarch64_get_qregoiv8hi (o, 0);
  ret.val[1] = (uint16x8_t) builtin_aarch64_get_qregoiv8hi (o, 1);
  return ret;
}

DEF: uint32x4x2
vld2q_dup_u32 (const uint32_t * a)
{
  uint32x4x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rv4si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (uint32x4_t) builtin_aarch64_get_qregoiv4si (o, 0);
  ret.val[1] = (uint32x4_t) builtin_aarch64_get_qregoiv4si (o, 1);
  return ret;
}

DEF: uint64x2x2
vld2q_dup_u64 (const uint64_t * a)
{
  uint64x2x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rv2di ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (uint64x2_t) builtin_aarch64_get_qregoiv2di (o, 0);
  ret.val[1] = (uint64x2_t) builtin_aarch64_get_qregoiv2di (o, 1);
  return ret;
}

DEF: float32x4x2
vld2q_dup_f32 (const float32_t * a)
{
  float32x4x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rv4sf ((const builtin_aarch64_simd_sf *) a);
  ret.val[0] = (float32x4_t) builtin_aarch64_get_qregoiv4sf (o, 0);
  ret.val[1] = (float32x4_t) builtin_aarch64_get_qregoiv4sf (o, 1);
  return ret;
}

DEF: float64x2x2
vld2q_dup_f64 (const float64_t * a)
{
  float64x2x2_t ret;
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_ld2rv2df ((const builtin_aarch64_simd_df *) a);
  ret.val[0] = (float64x2_t) builtin_aarch64_get_qregoiv2df (o, 0);
  ret.val[1] = (float64x2_t) builtin_aarch64_get_qregoiv2df (o, 1);
  return ret;
}

DEF: int64x1x3
vld3_dup_s64 (const int64_t * a)
{
  int64x1x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rdi ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (int64x1_t) builtin_aarch64_get_dregcidi (o, 0);
  ret.val[1] = (int64x1_t) builtin_aarch64_get_dregcidi (o, 1);
  ret.val[2] = (int64x1_t) builtin_aarch64_get_dregcidi (o, 2);
  return ret;
}

DEF: uint64x1x3
vld3_dup_u64 (const uint64_t * a)
{
  uint64x1x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rdi ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (uint64x1_t) builtin_aarch64_get_dregcidi (o, 0);
  ret.val[1] = (uint64x1_t) builtin_aarch64_get_dregcidi (o, 1);
  ret.val[2] = (uint64x1_t) builtin_aarch64_get_dregcidi (o, 2);
  return ret;
}

DEF: float64x1x3
vld3_dup_f64 (const float64_t * a)
{
  float64x1x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rdf ((const builtin_aarch64_simd_df *) a);
  ret.val[0] = (float64x1_t) {builtin_aarch64_get_dregcidf (o, 0)};
  ret.val[1] = (float64x1_t) {builtin_aarch64_get_dregcidf (o, 1)};
  ret.val[2] = (float64x1_t) {builtin_aarch64_get_dregcidf (o, 2)};
  return ret;
}

DEF: int8x8x3
vld3_dup_s8 (const int8_t * a)
{
  int8x8x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rv8qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (int8x8_t) builtin_aarch64_get_dregciv8qi (o, 0);
  ret.val[1] = (int8x8_t) builtin_aarch64_get_dregciv8qi (o, 1);
  ret.val[2] = (int8x8_t) builtin_aarch64_get_dregciv8qi (o, 2);
  return ret;
}

DEF: poly8x8x3
vld3_dup_p8 (const poly8_t * a)
{
  poly8x8x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rv8qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (poly8x8_t) builtin_aarch64_get_dregciv8qi (o, 0);
  ret.val[1] = (poly8x8_t) builtin_aarch64_get_dregciv8qi (o, 1);
  ret.val[2] = (poly8x8_t) builtin_aarch64_get_dregciv8qi (o, 2);
  return ret;
}

DEF: int16x4x3
vld3_dup_s16 (const int16_t * a)
{
  int16x4x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rv4hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (int16x4_t) builtin_aarch64_get_dregciv4hi (o, 0);
  ret.val[1] = (int16x4_t) builtin_aarch64_get_dregciv4hi (o, 1);
  ret.val[2] = (int16x4_t) builtin_aarch64_get_dregciv4hi (o, 2);
  return ret;
}

DEF: poly16x4x3
vld3_dup_p16 (const poly16_t * a)
{
  poly16x4x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rv4hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (poly16x4_t) builtin_aarch64_get_dregciv4hi (o, 0);
  ret.val[1] = (poly16x4_t) builtin_aarch64_get_dregciv4hi (o, 1);
  ret.val[2] = (poly16x4_t) builtin_aarch64_get_dregciv4hi (o, 2);
  return ret;
}

DEF: int32x2x3
vld3_dup_s32 (const int32_t * a)
{
  int32x2x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rv2si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (int32x2_t) builtin_aarch64_get_dregciv2si (o, 0);
  ret.val[1] = (int32x2_t) builtin_aarch64_get_dregciv2si (o, 1);
  ret.val[2] = (int32x2_t) builtin_aarch64_get_dregciv2si (o, 2);
  return ret;
}

DEF: uint8x8x3
vld3_dup_u8 (const uint8_t * a)
{
  uint8x8x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rv8qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (uint8x8_t) builtin_aarch64_get_dregciv8qi (o, 0);
  ret.val[1] = (uint8x8_t) builtin_aarch64_get_dregciv8qi (o, 1);
  ret.val[2] = (uint8x8_t) builtin_aarch64_get_dregciv8qi (o, 2);
  return ret;
}

DEF: uint16x4x3
vld3_dup_u16 (const uint16_t * a)
{
  uint16x4x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rv4hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (uint16x4_t) builtin_aarch64_get_dregciv4hi (o, 0);
  ret.val[1] = (uint16x4_t) builtin_aarch64_get_dregciv4hi (o, 1);
  ret.val[2] = (uint16x4_t) builtin_aarch64_get_dregciv4hi (o, 2);
  return ret;
}

DEF: uint32x2x3
vld3_dup_u32 (const uint32_t * a)
{
  uint32x2x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rv2si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (uint32x2_t) builtin_aarch64_get_dregciv2si (o, 0);
  ret.val[1] = (uint32x2_t) builtin_aarch64_get_dregciv2si (o, 1);
  ret.val[2] = (uint32x2_t) builtin_aarch64_get_dregciv2si (o, 2);
  return ret;
}

DEF: float32x2x3
vld3_dup_f32 (const float32_t * a)
{
  float32x2x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rv2sf ((const builtin_aarch64_simd_sf *) a);
  ret.val[0] = (float32x2_t) builtin_aarch64_get_dregciv2sf (o, 0);
  ret.val[1] = (float32x2_t) builtin_aarch64_get_dregciv2sf (o, 1);
  ret.val[2] = (float32x2_t) builtin_aarch64_get_dregciv2sf (o, 2);
  return ret;
}

DEF: int8x16x3
vld3q_dup_s8 (const int8_t * a)
{
  int8x16x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rv16qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (int8x16_t) builtin_aarch64_get_qregciv16qi (o, 0);
  ret.val[1] = (int8x16_t) builtin_aarch64_get_qregciv16qi (o, 1);
  ret.val[2] = (int8x16_t) builtin_aarch64_get_qregciv16qi (o, 2);
  return ret;
}

DEF: poly8x16x3
vld3q_dup_p8 (const poly8_t * a)
{
  poly8x16x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rv16qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (poly8x16_t) builtin_aarch64_get_qregciv16qi (o, 0);
  ret.val[1] = (poly8x16_t) builtin_aarch64_get_qregciv16qi (o, 1);
  ret.val[2] = (poly8x16_t) builtin_aarch64_get_qregciv16qi (o, 2);
  return ret;
}

DEF: int16x8x3
vld3q_dup_s16 (const int16_t * a)
{
  int16x8x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rv8hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (int16x8_t) builtin_aarch64_get_qregciv8hi (o, 0);
  ret.val[1] = (int16x8_t) builtin_aarch64_get_qregciv8hi (o, 1);
  ret.val[2] = (int16x8_t) builtin_aarch64_get_qregciv8hi (o, 2);
  return ret;
}

DEF: poly16x8x3
vld3q_dup_p16 (const poly16_t * a)
{
  poly16x8x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rv8hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (poly16x8_t) builtin_aarch64_get_qregciv8hi (o, 0);
  ret.val[1] = (poly16x8_t) builtin_aarch64_get_qregciv8hi (o, 1);
  ret.val[2] = (poly16x8_t) builtin_aarch64_get_qregciv8hi (o, 2);
  return ret;
}

DEF: int32x4x3
vld3q_dup_s32 (const int32_t * a)
{
  int32x4x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rv4si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (int32x4_t) builtin_aarch64_get_qregciv4si (o, 0);
  ret.val[1] = (int32x4_t) builtin_aarch64_get_qregciv4si (o, 1);
  ret.val[2] = (int32x4_t) builtin_aarch64_get_qregciv4si (o, 2);
  return ret;
}

DEF: int64x2x3
vld3q_dup_s64 (const int64_t * a)
{
  int64x2x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rv2di ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (int64x2_t) builtin_aarch64_get_qregciv2di (o, 0);
  ret.val[1] = (int64x2_t) builtin_aarch64_get_qregciv2di (o, 1);
  ret.val[2] = (int64x2_t) builtin_aarch64_get_qregciv2di (o, 2);
  return ret;
}

DEF: uint8x16x3
vld3q_dup_u8 (const uint8_t * a)
{
  uint8x16x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rv16qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (uint8x16_t) builtin_aarch64_get_qregciv16qi (o, 0);
  ret.val[1] = (uint8x16_t) builtin_aarch64_get_qregciv16qi (o, 1);
  ret.val[2] = (uint8x16_t) builtin_aarch64_get_qregciv16qi (o, 2);
  return ret;
}

DEF: uint16x8x3
vld3q_dup_u16 (const uint16_t * a)
{
  uint16x8x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rv8hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (uint16x8_t) builtin_aarch64_get_qregciv8hi (o, 0);
  ret.val[1] = (uint16x8_t) builtin_aarch64_get_qregciv8hi (o, 1);
  ret.val[2] = (uint16x8_t) builtin_aarch64_get_qregciv8hi (o, 2);
  return ret;
}

DEF: uint32x4x3
vld3q_dup_u32 (const uint32_t * a)
{
  uint32x4x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rv4si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (uint32x4_t) builtin_aarch64_get_qregciv4si (o, 0);
  ret.val[1] = (uint32x4_t) builtin_aarch64_get_qregciv4si (o, 1);
  ret.val[2] = (uint32x4_t) builtin_aarch64_get_qregciv4si (o, 2);
  return ret;
}

DEF: uint64x2x3
vld3q_dup_u64 (const uint64_t * a)
{
  uint64x2x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rv2di ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (uint64x2_t) builtin_aarch64_get_qregciv2di (o, 0);
  ret.val[1] = (uint64x2_t) builtin_aarch64_get_qregciv2di (o, 1);
  ret.val[2] = (uint64x2_t) builtin_aarch64_get_qregciv2di (o, 2);
  return ret;
}

DEF: float32x4x3
vld3q_dup_f32 (const float32_t * a)
{
  float32x4x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rv4sf ((const builtin_aarch64_simd_sf *) a);
  ret.val[0] = (float32x4_t) builtin_aarch64_get_qregciv4sf (o, 0);
  ret.val[1] = (float32x4_t) builtin_aarch64_get_qregciv4sf (o, 1);
  ret.val[2] = (float32x4_t) builtin_aarch64_get_qregciv4sf (o, 2);
  return ret;
}

DEF: float64x2x3
vld3q_dup_f64 (const float64_t * a)
{
  float64x2x3_t ret;
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_ld3rv2df ((const builtin_aarch64_simd_df *) a);
  ret.val[0] = (float64x2_t) builtin_aarch64_get_qregciv2df (o, 0);
  ret.val[1] = (float64x2_t) builtin_aarch64_get_qregciv2df (o, 1);
  ret.val[2] = (float64x2_t) builtin_aarch64_get_qregciv2df (o, 2);
  return ret;
}

DEF: int64x1x4
vld4_dup_s64 (const int64_t * a)
{
  int64x1x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rdi ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (int64x1_t) builtin_aarch64_get_dregxidi (o, 0);
  ret.val[1] = (int64x1_t) builtin_aarch64_get_dregxidi (o, 1);
  ret.val[2] = (int64x1_t) builtin_aarch64_get_dregxidi (o, 2);
  ret.val[3] = (int64x1_t) builtin_aarch64_get_dregxidi (o, 3);
  return ret;
}

DEF: uint64x1x4
vld4_dup_u64 (const uint64_t * a)
{
  uint64x1x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rdi ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (uint64x1_t) builtin_aarch64_get_dregxidi (o, 0);
  ret.val[1] = (uint64x1_t) builtin_aarch64_get_dregxidi (o, 1);
  ret.val[2] = (uint64x1_t) builtin_aarch64_get_dregxidi (o, 2);
  ret.val[3] = (uint64x1_t) builtin_aarch64_get_dregxidi (o, 3);
  return ret;
}

DEF: float64x1x4
vld4_dup_f64 (const float64_t * a)
{
  float64x1x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rdf ((const builtin_aarch64_simd_df *) a);
  ret.val[0] = (float64x1_t) {builtin_aarch64_get_dregxidf (o, 0)};
  ret.val[1] = (float64x1_t) {builtin_aarch64_get_dregxidf (o, 1)};
  ret.val[2] = (float64x1_t) {builtin_aarch64_get_dregxidf (o, 2)};
  ret.val[3] = (float64x1_t) {builtin_aarch64_get_dregxidf (o, 3)};
  return ret;
}

DEF: int8x8x4
vld4_dup_s8 (const int8_t * a)
{
  int8x8x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rv8qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (int8x8_t) builtin_aarch64_get_dregxiv8qi (o, 0);
  ret.val[1] = (int8x8_t) builtin_aarch64_get_dregxiv8qi (o, 1);
  ret.val[2] = (int8x8_t) builtin_aarch64_get_dregxiv8qi (o, 2);
  ret.val[3] = (int8x8_t) builtin_aarch64_get_dregxiv8qi (o, 3);
  return ret;
}

DEF: poly8x8x4
vld4_dup_p8 (const poly8_t * a)
{
  poly8x8x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rv8qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (poly8x8_t) builtin_aarch64_get_dregxiv8qi (o, 0);
  ret.val[1] = (poly8x8_t) builtin_aarch64_get_dregxiv8qi (o, 1);
  ret.val[2] = (poly8x8_t) builtin_aarch64_get_dregxiv8qi (o, 2);
  ret.val[3] = (poly8x8_t) builtin_aarch64_get_dregxiv8qi (o, 3);
  return ret;
}

DEF: int16x4x4
vld4_dup_s16 (const int16_t * a)
{
  int16x4x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rv4hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (int16x4_t) builtin_aarch64_get_dregxiv4hi (o, 0);
  ret.val[1] = (int16x4_t) builtin_aarch64_get_dregxiv4hi (o, 1);
  ret.val[2] = (int16x4_t) builtin_aarch64_get_dregxiv4hi (o, 2);
  ret.val[3] = (int16x4_t) builtin_aarch64_get_dregxiv4hi (o, 3);
  return ret;
}

DEF: poly16x4x4
vld4_dup_p16 (const poly16_t * a)
{
  poly16x4x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rv4hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (poly16x4_t) builtin_aarch64_get_dregxiv4hi (o, 0);
  ret.val[1] = (poly16x4_t) builtin_aarch64_get_dregxiv4hi (o, 1);
  ret.val[2] = (poly16x4_t) builtin_aarch64_get_dregxiv4hi (o, 2);
  ret.val[3] = (poly16x4_t) builtin_aarch64_get_dregxiv4hi (o, 3);
  return ret;
}

DEF: int32x2x4
vld4_dup_s32 (const int32_t * a)
{
  int32x2x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rv2si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (int32x2_t) builtin_aarch64_get_dregxiv2si (o, 0);
  ret.val[1] = (int32x2_t) builtin_aarch64_get_dregxiv2si (o, 1);
  ret.val[2] = (int32x2_t) builtin_aarch64_get_dregxiv2si (o, 2);
  ret.val[3] = (int32x2_t) builtin_aarch64_get_dregxiv2si (o, 3);
  return ret;
}

DEF: uint8x8x4
vld4_dup_u8 (const uint8_t * a)
{
  uint8x8x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rv8qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (uint8x8_t) builtin_aarch64_get_dregxiv8qi (o, 0);
  ret.val[1] = (uint8x8_t) builtin_aarch64_get_dregxiv8qi (o, 1);
  ret.val[2] = (uint8x8_t) builtin_aarch64_get_dregxiv8qi (o, 2);
  ret.val[3] = (uint8x8_t) builtin_aarch64_get_dregxiv8qi (o, 3);
  return ret;
}

DEF: uint16x4x4
vld4_dup_u16 (const uint16_t * a)
{
  uint16x4x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rv4hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (uint16x4_t) builtin_aarch64_get_dregxiv4hi (o, 0);
  ret.val[1] = (uint16x4_t) builtin_aarch64_get_dregxiv4hi (o, 1);
  ret.val[2] = (uint16x4_t) builtin_aarch64_get_dregxiv4hi (o, 2);
  ret.val[3] = (uint16x4_t) builtin_aarch64_get_dregxiv4hi (o, 3);
  return ret;
}

DEF: uint32x2x4
vld4_dup_u32 (const uint32_t * a)
{
  uint32x2x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rv2si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (uint32x2_t) builtin_aarch64_get_dregxiv2si (o, 0);
  ret.val[1] = (uint32x2_t) builtin_aarch64_get_dregxiv2si (o, 1);
  ret.val[2] = (uint32x2_t) builtin_aarch64_get_dregxiv2si (o, 2);
  ret.val[3] = (uint32x2_t) builtin_aarch64_get_dregxiv2si (o, 3);
  return ret;
}

DEF: float32x2x4
vld4_dup_f32 (const float32_t * a)
{
  float32x2x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rv2sf ((const builtin_aarch64_simd_sf *) a);
  ret.val[0] = (float32x2_t) builtin_aarch64_get_dregxiv2sf (o, 0);
  ret.val[1] = (float32x2_t) builtin_aarch64_get_dregxiv2sf (o, 1);
  ret.val[2] = (float32x2_t) builtin_aarch64_get_dregxiv2sf (o, 2);
  ret.val[3] = (float32x2_t) builtin_aarch64_get_dregxiv2sf (o, 3);
  return ret;
}

DEF: int8x16x4
vld4q_dup_s8 (const int8_t * a)
{
  int8x16x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rv16qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (int8x16_t) builtin_aarch64_get_qregxiv16qi (o, 0);
  ret.val[1] = (int8x16_t) builtin_aarch64_get_qregxiv16qi (o, 1);
  ret.val[2] = (int8x16_t) builtin_aarch64_get_qregxiv16qi (o, 2);
  ret.val[3] = (int8x16_t) builtin_aarch64_get_qregxiv16qi (o, 3);
  return ret;
}

DEF: poly8x16x4
vld4q_dup_p8 (const poly8_t * a)
{
  poly8x16x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rv16qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (poly8x16_t) builtin_aarch64_get_qregxiv16qi (o, 0);
  ret.val[1] = (poly8x16_t) builtin_aarch64_get_qregxiv16qi (o, 1);
  ret.val[2] = (poly8x16_t) builtin_aarch64_get_qregxiv16qi (o, 2);
  ret.val[3] = (poly8x16_t) builtin_aarch64_get_qregxiv16qi (o, 3);
  return ret;
}

DEF: int16x8x4
vld4q_dup_s16 (const int16_t * a)
{
  int16x8x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rv8hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (int16x8_t) builtin_aarch64_get_qregxiv8hi (o, 0);
  ret.val[1] = (int16x8_t) builtin_aarch64_get_qregxiv8hi (o, 1);
  ret.val[2] = (int16x8_t) builtin_aarch64_get_qregxiv8hi (o, 2);
  ret.val[3] = (int16x8_t) builtin_aarch64_get_qregxiv8hi (o, 3);
  return ret;
}

DEF: poly16x8x4
vld4q_dup_p16 (const poly16_t * a)
{
  poly16x8x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rv8hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (poly16x8_t) builtin_aarch64_get_qregxiv8hi (o, 0);
  ret.val[1] = (poly16x8_t) builtin_aarch64_get_qregxiv8hi (o, 1);
  ret.val[2] = (poly16x8_t) builtin_aarch64_get_qregxiv8hi (o, 2);
  ret.val[3] = (poly16x8_t) builtin_aarch64_get_qregxiv8hi (o, 3);
  return ret;
}

DEF: int32x4x4
vld4q_dup_s32 (const int32_t * a)
{
  int32x4x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rv4si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (int32x4_t) builtin_aarch64_get_qregxiv4si (o, 0);
  ret.val[1] = (int32x4_t) builtin_aarch64_get_qregxiv4si (o, 1);
  ret.val[2] = (int32x4_t) builtin_aarch64_get_qregxiv4si (o, 2);
  ret.val[3] = (int32x4_t) builtin_aarch64_get_qregxiv4si (o, 3);
  return ret;
}

DEF: int64x2x4
vld4q_dup_s64 (const int64_t * a)
{
  int64x2x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rv2di ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (int64x2_t) builtin_aarch64_get_qregxiv2di (o, 0);
  ret.val[1] = (int64x2_t) builtin_aarch64_get_qregxiv2di (o, 1);
  ret.val[2] = (int64x2_t) builtin_aarch64_get_qregxiv2di (o, 2);
  ret.val[3] = (int64x2_t) builtin_aarch64_get_qregxiv2di (o, 3);
  return ret;
}

DEF: uint8x16x4
vld4q_dup_u8 (const uint8_t * a)
{
  uint8x16x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rv16qi ((const builtin_aarch64_simd_qi *) a);
  ret.val[0] = (uint8x16_t) builtin_aarch64_get_qregxiv16qi (o, 0);
  ret.val[1] = (uint8x16_t) builtin_aarch64_get_qregxiv16qi (o, 1);
  ret.val[2] = (uint8x16_t) builtin_aarch64_get_qregxiv16qi (o, 2);
  ret.val[3] = (uint8x16_t) builtin_aarch64_get_qregxiv16qi (o, 3);
  return ret;
}

DEF: uint16x8x4
vld4q_dup_u16 (const uint16_t * a)
{
  uint16x8x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rv8hi ((const builtin_aarch64_simd_hi *) a);
  ret.val[0] = (uint16x8_t) builtin_aarch64_get_qregxiv8hi (o, 0);
  ret.val[1] = (uint16x8_t) builtin_aarch64_get_qregxiv8hi (o, 1);
  ret.val[2] = (uint16x8_t) builtin_aarch64_get_qregxiv8hi (o, 2);
  ret.val[3] = (uint16x8_t) builtin_aarch64_get_qregxiv8hi (o, 3);
  return ret;
}

DEF: uint32x4x4
vld4q_dup_u32 (const uint32_t * a)
{
  uint32x4x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rv4si ((const builtin_aarch64_simd_si *) a);
  ret.val[0] = (uint32x4_t) builtin_aarch64_get_qregxiv4si (o, 0);
  ret.val[1] = (uint32x4_t) builtin_aarch64_get_qregxiv4si (o, 1);
  ret.val[2] = (uint32x4_t) builtin_aarch64_get_qregxiv4si (o, 2);
  ret.val[3] = (uint32x4_t) builtin_aarch64_get_qregxiv4si (o, 3);
  return ret;
}

DEF: uint64x2x4
vld4q_dup_u64 (const uint64_t * a)
{
  uint64x2x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rv2di ((const builtin_aarch64_simd_di *) a);
  ret.val[0] = (uint64x2_t) builtin_aarch64_get_qregxiv2di (o, 0);
  ret.val[1] = (uint64x2_t) builtin_aarch64_get_qregxiv2di (o, 1);
  ret.val[2] = (uint64x2_t) builtin_aarch64_get_qregxiv2di (o, 2);
  ret.val[3] = (uint64x2_t) builtin_aarch64_get_qregxiv2di (o, 3);
  return ret;
}

DEF: float32x4x4
vld4q_dup_f32 (const float32_t * a)
{
  float32x4x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rv4sf ((const builtin_aarch64_simd_sf *) a);
  ret.val[0] = (float32x4_t) builtin_aarch64_get_qregxiv4sf (o, 0);
  ret.val[1] = (float32x4_t) builtin_aarch64_get_qregxiv4sf (o, 1);
  ret.val[2] = (float32x4_t) builtin_aarch64_get_qregxiv4sf (o, 2);
  ret.val[3] = (float32x4_t) builtin_aarch64_get_qregxiv4sf (o, 3);
  return ret;
}

DEF: float64x2x4
vld4q_dup_f64 (const float64_t * a)
{
  float64x2x4_t ret;
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_ld4rv2df ((const builtin_aarch64_simd_df *) a);
  ret.val[0] = (float64x2_t) builtin_aarch64_get_qregxiv2df (o, 0);
  ret.val[1] = (float64x2_t) builtin_aarch64_get_qregxiv2df (o, 1);
  ret.val[2] = (float64x2_t) builtin_aarch64_get_qregxiv2df (o, 2);
  ret.val[3] = (float64x2_t) builtin_aarch64_get_qregxiv2df (o, 3);
  return ret;
}
# 16842 "arm_neon.h"
DEF: float32x2x2 vld2_lane_f32 (const float32_t * ptr, float32x2x2_t b, const int c) { builtin_aarch64_simd_oi o; float32x4x2_t temp; temp.val[0] = vcombine_f32 (b.val[0], vcreate_f32 (0)); temp.val[1] = vcombine_f32 (b.val[1], vcreate_f32 (0)); o = builtin_aarch64_set_qregoiv4sf (o, (float32x4_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv4sf (o, (float32x4_t) temp.val[1], 1); o = builtin_aarch64_ld2_lanev2sf ( (builtin_aarch64_simd_sf *) ptr, o, c); b.val[0] = (float32x2_t) builtin_aarch64_get_dregoidi (o, 0); b.val[1] = (float32x2_t) builtin_aarch64_get_dregoidi (o, 1); return b; }

DEF: float64x1x2 vld2_lane_f64 (const float64_t * ptr, float64x1x2_t b, const int c) { builtin_aarch64_simd_oi o; float64x2x2_t temp; temp.val[0] = vcombine_f64 (b.val[0], vcreate_f64 (0)); temp.val[1] = vcombine_f64 (b.val[1], vcreate_f64 (0)); o = builtin_aarch64_set_qregoiv2df (o, (float64x2_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv2df (o, (float64x2_t) temp.val[1], 1); o = builtin_aarch64_ld2_lanedf ( (builtin_aarch64_simd_df *) ptr, o, c); b.val[0] = (float64x1_t) builtin_aarch64_get_dregoidi (o, 0); b.val[1] = (float64x1_t) builtin_aarch64_get_dregoidi (o, 1); return b; }

DEF: poly8x8x2 vld2_lane_p8 (const poly8_t * ptr, poly8x8x2_t b, const int c) { builtin_aarch64_simd_oi o; poly8x16x2_t temp; temp.val[0] = vcombine_p8 (b.val[0], vcreate_p8 (0)); temp.val[1] = vcombine_p8 (b.val[1], vcreate_p8 (0)); o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) temp.val[1], 1); o = builtin_aarch64_ld2_lanev8qi ( (builtin_aarch64_simd_qi *) ptr, o, c); b.val[0] = (poly8x8_t) builtin_aarch64_get_dregoidi (o, 0); b.val[1] = (poly8x8_t) builtin_aarch64_get_dregoidi (o, 1); return b; }

DEF: poly16x4x2 vld2_lane_p16 (const poly16_t * ptr, poly16x4x2_t b, const int c) { builtin_aarch64_simd_oi o; poly16x8x2_t temp; temp.val[0] = vcombine_p16 (b.val[0], vcreate_p16 (0)); temp.val[1] = vcombine_p16 (b.val[1], vcreate_p16 (0)); o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) temp.val[1], 1); o = builtin_aarch64_ld2_lanev4hi ( (builtin_aarch64_simd_hi *) ptr, o, c); b.val[0] = (poly16x4_t) builtin_aarch64_get_dregoidi (o, 0); b.val[1] = (poly16x4_t) builtin_aarch64_get_dregoidi (o, 1); return b; }

DEF: int8x8x2 vld2_lane_s8 (const int8_t * ptr, int8x8x2_t b, const int c) { builtin_aarch64_simd_oi o; int8x16x2_t temp; temp.val[0] = vcombine_s8 (b.val[0], vcreate_s8 (0)); temp.val[1] = vcombine_s8 (b.val[1], vcreate_s8 (0)); o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) temp.val[1], 1); o = builtin_aarch64_ld2_lanev8qi ( (builtin_aarch64_simd_qi *) ptr, o, c); b.val[0] = (int8x8_t) builtin_aarch64_get_dregoidi (o, 0); b.val[1] = (int8x8_t) builtin_aarch64_get_dregoidi (o, 1); return b; }

DEF: int16x4x2 vld2_lane_s16 (const int16_t * ptr, int16x4x2_t b, const int c) { builtin_aarch64_simd_oi o; int16x8x2_t temp; temp.val[0] = vcombine_s16 (b.val[0], vcreate_s16 (0)); temp.val[1] = vcombine_s16 (b.val[1], vcreate_s16 (0)); o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) temp.val[1], 1); o = builtin_aarch64_ld2_lanev4hi ( (builtin_aarch64_simd_hi *) ptr, o, c); b.val[0] = (int16x4_t) builtin_aarch64_get_dregoidi (o, 0); b.val[1] = (int16x4_t) builtin_aarch64_get_dregoidi (o, 1); return b; }

DEF: int32x2x2 vld2_lane_s32 (const int32_t * ptr, int32x2x2_t b, const int c) { builtin_aarch64_simd_oi o; int32x4x2_t temp; temp.val[0] = vcombine_s32 (b.val[0], vcreate_s32 (0)); temp.val[1] = vcombine_s32 (b.val[1], vcreate_s32 (0)); o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) temp.val[1], 1); o = builtin_aarch64_ld2_lanev2si ( (builtin_aarch64_simd_si *) ptr, o, c); b.val[0] = (int32x2_t) builtin_aarch64_get_dregoidi (o, 0); b.val[1] = (int32x2_t) builtin_aarch64_get_dregoidi (o, 1); return b; }

DEF: int64x1x2 vld2_lane_s64 (const int64_t * ptr, int64x1x2_t b, const int c) { builtin_aarch64_simd_oi o; int64x2x2_t temp; temp.val[0] = vcombine_s64 (b.val[0], vcreate_s64 (0)); temp.val[1] = vcombine_s64 (b.val[1], vcreate_s64 (0)); o = builtin_aarch64_set_qregoiv2di (o, (int64x2_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv2di (o, (int64x2_t) temp.val[1], 1); o = builtin_aarch64_ld2_lanedi ( (builtin_aarch64_simd_di *) ptr, o, c); b.val[0] = (int64x1_t) builtin_aarch64_get_dregoidi (o, 0); b.val[1] = (int64x1_t) builtin_aarch64_get_dregoidi (o, 1); return b; }

DEF: uint8x8x2 vld2_lane_u8 (const uint8_t * ptr, uint8x8x2_t b, const int c) { builtin_aarch64_simd_oi o; uint8x16x2_t temp; temp.val[0] = vcombine_u8 (b.val[0], vcreate_u8 (0)); temp.val[1] = vcombine_u8 (b.val[1], vcreate_u8 (0)); o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) temp.val[1], 1); o = builtin_aarch64_ld2_lanev8qi ( (builtin_aarch64_simd_qi *) ptr, o, c); b.val[0] = (uint8x8_t) builtin_aarch64_get_dregoidi (o, 0); b.val[1] = (uint8x8_t) builtin_aarch64_get_dregoidi (o, 1); return b; }

DEF: uint16x4x2 vld2_lane_u16 (const uint16_t * ptr, uint16x4x2_t b, const int c) { builtin_aarch64_simd_oi o; uint16x8x2_t temp; temp.val[0] = vcombine_u16 (b.val[0], vcreate_u16 (0)); temp.val[1] = vcombine_u16 (b.val[1], vcreate_u16 (0)); o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) temp.val[1], 1); o = builtin_aarch64_ld2_lanev4hi ( (builtin_aarch64_simd_hi *) ptr, o, c); b.val[0] = (uint16x4_t) builtin_aarch64_get_dregoidi (o, 0); b.val[1] = (uint16x4_t) builtin_aarch64_get_dregoidi (o, 1); return b; }

DEF: uint32x2x2 vld2_lane_u32 (const uint32_t * ptr, uint32x2x2_t b, const int c) { builtin_aarch64_simd_oi o; uint32x4x2_t temp; temp.val[0] = vcombine_u32 (b.val[0], vcreate_u32 (0)); temp.val[1] = vcombine_u32 (b.val[1], vcreate_u32 (0)); o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) temp.val[1], 1); o = builtin_aarch64_ld2_lanev2si ( (builtin_aarch64_simd_si *) ptr, o, c); b.val[0] = (uint32x2_t) builtin_aarch64_get_dregoidi (o, 0); b.val[1] = (uint32x2_t) builtin_aarch64_get_dregoidi (o, 1); return b; }

DEF: uint64x1x2 vld2_lane_u64 (const uint64_t * ptr, uint64x1x2_t b, const int c) { builtin_aarch64_simd_oi o; uint64x2x2_t temp; temp.val[0] = vcombine_u64 (b.val[0], vcreate_u64 (0)); temp.val[1] = vcombine_u64 (b.val[1], vcreate_u64 (0)); o = builtin_aarch64_set_qregoiv2di (o, (int64x2_t) temp.val[0], 0); o = builtin_aarch64_set_qregoiv2di (o, (int64x2_t) temp.val[1], 1); o = builtin_aarch64_ld2_lanedi ( (builtin_aarch64_simd_di *) ptr, o, c); b.val[0] = (uint64x1_t) builtin_aarch64_get_dregoidi (o, 0); b.val[1] = (uint64x1_t) builtin_aarch64_get_dregoidi (o, 1); return b; }
# 16886 "arm_neon.h"
DEF: float32x4x2 vld2q_lane_f32 (const float32_t * ptr, float32x4x2_t b, const int c) { builtin_aarch64_simd_oi o; float32x4x2_t ret; o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_ld2_lanev4sf ( (builtin_aarch64_simd_sf *) ptr, o, c); ret.val[0] = (float32x4_t) builtin_aarch64_get_qregoiv4si (o, 0); ret.val[1] = (float32x4_t) builtin_aarch64_get_qregoiv4si (o, 1); return ret; }
DEF: float64x2x2 vld2q_lane_f64 (const float64_t * ptr, float64x2x2_t b, const int c) { builtin_aarch64_simd_oi o; float64x2x2_t ret; o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_ld2_lanev2df ( (builtin_aarch64_simd_df *) ptr, o, c); ret.val[0] = (float64x2_t) builtin_aarch64_get_qregoiv4si (o, 0); ret.val[1] = (float64x2_t) builtin_aarch64_get_qregoiv4si (o, 1); return ret; }
DEF: poly8x16x2 vld2q_lane_p8 (const poly8_t * ptr, poly8x16x2_t b, const int c) { builtin_aarch64_simd_oi o; poly8x16x2_t ret; o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_ld2_lanev16qi ( (builtin_aarch64_simd_qi *) ptr, o, c); ret.val[0] = (poly8x16_t) builtin_aarch64_get_qregoiv4si (o, 0); ret.val[1] = (poly8x16_t) builtin_aarch64_get_qregoiv4si (o, 1); return ret; }
DEF: poly16x8x2 vld2q_lane_p16 (const poly16_t * ptr, poly16x8x2_t b, const int c) { builtin_aarch64_simd_oi o; poly16x8x2_t ret; o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_ld2_lanev8hi ( (builtin_aarch64_simd_hi *) ptr, o, c); ret.val[0] = (poly16x8_t) builtin_aarch64_get_qregoiv4si (o, 0); ret.val[1] = (poly16x8_t) builtin_aarch64_get_qregoiv4si (o, 1); return ret; }
DEF: int8x16x2 vld2q_lane_s8 (const int8_t * ptr, int8x16x2_t b, const int c) { builtin_aarch64_simd_oi o; int8x16x2_t ret; o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_ld2_lanev16qi ( (builtin_aarch64_simd_qi *) ptr, o, c); ret.val[0] = (int8x16_t) builtin_aarch64_get_qregoiv4si (o, 0); ret.val[1] = (int8x16_t) builtin_aarch64_get_qregoiv4si (o, 1); return ret; }
DEF: int16x8x2 vld2q_lane_s16 (const int16_t * ptr, int16x8x2_t b, const int c) { builtin_aarch64_simd_oi o; int16x8x2_t ret; o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_ld2_lanev8hi ( (builtin_aarch64_simd_hi *) ptr, o, c); ret.val[0] = (int16x8_t) builtin_aarch64_get_qregoiv4si (o, 0); ret.val[1] = (int16x8_t) builtin_aarch64_get_qregoiv4si (o, 1); return ret; }
DEF: int32x4x2 vld2q_lane_s32 (const int32_t * ptr, int32x4x2_t b, const int c) { builtin_aarch64_simd_oi o; int32x4x2_t ret; o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_ld2_lanev4si ( (builtin_aarch64_simd_si *) ptr, o, c); ret.val[0] = (int32x4_t) builtin_aarch64_get_qregoiv4si (o, 0); ret.val[1] = (int32x4_t) builtin_aarch64_get_qregoiv4si (o, 1); return ret; }
DEF: int64x2x2 vld2q_lane_s64 (const int64_t * ptr, int64x2x2_t b, const int c) { builtin_aarch64_simd_oi o; int64x2x2_t ret; o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_ld2_lanev2di ( (builtin_aarch64_simd_di *) ptr, o, c); ret.val[0] = (int64x2_t) builtin_aarch64_get_qregoiv4si (o, 0); ret.val[1] = (int64x2_t) builtin_aarch64_get_qregoiv4si (o, 1); return ret; }
DEF: uint8x16x2 vld2q_lane_u8 (const uint8_t * ptr, uint8x16x2_t b, const int c) { builtin_aarch64_simd_oi o; uint8x16x2_t ret; o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_ld2_lanev16qi ( (builtin_aarch64_simd_qi *) ptr, o, c); ret.val[0] = (uint8x16_t) builtin_aarch64_get_qregoiv4si (o, 0); ret.val[1] = (uint8x16_t) builtin_aarch64_get_qregoiv4si (o, 1); return ret; }
DEF: uint16x8x2 vld2q_lane_u16 (const uint16_t * ptr, uint16x8x2_t b, const int c) { builtin_aarch64_simd_oi o; uint16x8x2_t ret; o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_ld2_lanev8hi ( (builtin_aarch64_simd_hi *) ptr, o, c); ret.val[0] = (uint16x8_t) builtin_aarch64_get_qregoiv4si (o, 0); ret.val[1] = (uint16x8_t) builtin_aarch64_get_qregoiv4si (o, 1); return ret; }
DEF: uint32x4x2 vld2q_lane_u32 (const uint32_t * ptr, uint32x4x2_t b, const int c) { builtin_aarch64_simd_oi o; uint32x4x2_t ret; o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_ld2_lanev4si ( (builtin_aarch64_simd_si *) ptr, o, c); ret.val[0] = (uint32x4_t) builtin_aarch64_get_qregoiv4si (o, 0); ret.val[1] = (uint32x4_t) builtin_aarch64_get_qregoiv4si (o, 1); return ret; }
DEF: uint64x2x2 vld2q_lane_u64 (const uint64_t * ptr, uint64x2x2_t b, const int c) { builtin_aarch64_simd_oi o; uint64x2x2_t ret; o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_ld2_lanev2di ( (builtin_aarch64_simd_di *) ptr, o, c); ret.val[0] = (uint64x2_t) builtin_aarch64_get_qregoiv4si (o, 0); ret.val[1] = (uint64x2_t) builtin_aarch64_get_qregoiv4si (o, 1); return ret; }
# 16933 "arm_neon.h"
DEF: float32x2x3 vld3_lane_f32 (const float32_t * ptr, float32x2x3_t b, const int c) { builtin_aarch64_simd_ci o; float32x4x3_t temp; temp.val[0] = vcombine_f32 (b.val[0], vcreate_f32 (0)); temp.val[1] = vcombine_f32 (b.val[1], vcreate_f32 (0)); temp.val[2] = vcombine_f32 (b.val[2], vcreate_f32 (0)); o = builtin_aarch64_set_qregciv4sf (o, (float32x4_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv4sf (o, (float32x4_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv4sf (o, (float32x4_t) temp.val[2], 2); o = builtin_aarch64_ld3_lanev2sf ( (builtin_aarch64_simd_sf *) ptr, o, c); b.val[0] = (float32x2_t) builtin_aarch64_get_dregcidi (o, 0); b.val[1] = (float32x2_t) builtin_aarch64_get_dregcidi (o, 1); b.val[2] = (float32x2_t) builtin_aarch64_get_dregcidi (o, 2); return b; }

DEF: float64x1x3 vld3_lane_f64 (const float64_t * ptr, float64x1x3_t b, const int c) { builtin_aarch64_simd_ci o; float64x2x3_t temp; temp.val[0] = vcombine_f64 (b.val[0], vcreate_f64 (0)); temp.val[1] = vcombine_f64 (b.val[1], vcreate_f64 (0)); temp.val[2] = vcombine_f64 (b.val[2], vcreate_f64 (0)); o = builtin_aarch64_set_qregciv2df (o, (float64x2_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv2df (o, (float64x2_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv2df (o, (float64x2_t) temp.val[2], 2); o = builtin_aarch64_ld3_lanedf ( (builtin_aarch64_simd_df *) ptr, o, c); b.val[0] = (float64x1_t) builtin_aarch64_get_dregcidi (o, 0); b.val[1] = (float64x1_t) builtin_aarch64_get_dregcidi (o, 1); b.val[2] = (float64x1_t) builtin_aarch64_get_dregcidi (o, 2); return b; }

DEF: poly8x8x3 vld3_lane_p8 (const poly8_t * ptr, poly8x8x3_t b, const int c) { builtin_aarch64_simd_ci o; poly8x16x3_t temp; temp.val[0] = vcombine_p8 (b.val[0], vcreate_p8 (0)); temp.val[1] = vcombine_p8 (b.val[1], vcreate_p8 (0)); temp.val[2] = vcombine_p8 (b.val[2], vcreate_p8 (0)); o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[2], 2); o = builtin_aarch64_ld3_lanev8qi ( (builtin_aarch64_simd_qi *) ptr, o, c); b.val[0] = (poly8x8_t) builtin_aarch64_get_dregcidi (o, 0); b.val[1] = (poly8x8_t) builtin_aarch64_get_dregcidi (o, 1); b.val[2] = (poly8x8_t) builtin_aarch64_get_dregcidi (o, 2); return b; }

DEF: poly16x4x3 vld3_lane_p16 (const poly16_t * ptr, poly16x4x3_t b, const int c) { builtin_aarch64_simd_ci o; poly16x8x3_t temp; temp.val[0] = vcombine_p16 (b.val[0], vcreate_p16 (0)); temp.val[1] = vcombine_p16 (b.val[1], vcreate_p16 (0)); temp.val[2] = vcombine_p16 (b.val[2], vcreate_p16 (0)); o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[2], 2); o = builtin_aarch64_ld3_lanev4hi ( (builtin_aarch64_simd_hi *) ptr, o, c); b.val[0] = (poly16x4_t) builtin_aarch64_get_dregcidi (o, 0); b.val[1] = (poly16x4_t) builtin_aarch64_get_dregcidi (o, 1); b.val[2] = (poly16x4_t) builtin_aarch64_get_dregcidi (o, 2); return b; }

DEF: int8x8x3 vld3_lane_s8 (const int8_t * ptr, int8x8x3_t b, const int c) { builtin_aarch64_simd_ci o; int8x16x3_t temp; temp.val[0] = vcombine_s8 (b.val[0], vcreate_s8 (0)); temp.val[1] = vcombine_s8 (b.val[1], vcreate_s8 (0)); temp.val[2] = vcombine_s8 (b.val[2], vcreate_s8 (0)); o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[2], 2); o = builtin_aarch64_ld3_lanev8qi ( (builtin_aarch64_simd_qi *) ptr, o, c); b.val[0] = (int8x8_t) builtin_aarch64_get_dregcidi (o, 0); b.val[1] = (int8x8_t) builtin_aarch64_get_dregcidi (o, 1); b.val[2] = (int8x8_t) builtin_aarch64_get_dregcidi (o, 2); return b; }

DEF: int16x4x3 vld3_lane_s16 (const int16_t * ptr, int16x4x3_t b, const int c) { builtin_aarch64_simd_ci o; int16x8x3_t temp; temp.val[0] = vcombine_s16 (b.val[0], vcreate_s16 (0)); temp.val[1] = vcombine_s16 (b.val[1], vcreate_s16 (0)); temp.val[2] = vcombine_s16 (b.val[2], vcreate_s16 (0)); o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[2], 2); o = builtin_aarch64_ld3_lanev4hi ( (builtin_aarch64_simd_hi *) ptr, o, c); b.val[0] = (int16x4_t) builtin_aarch64_get_dregcidi (o, 0); b.val[1] = (int16x4_t) builtin_aarch64_get_dregcidi (o, 1); b.val[2] = (int16x4_t) builtin_aarch64_get_dregcidi (o, 2); return b; }

DEF: int32x2x3 vld3_lane_s32 (const int32_t * ptr, int32x2x3_t b, const int c) { builtin_aarch64_simd_ci o; int32x4x3_t temp; temp.val[0] = vcombine_s32 (b.val[0], vcreate_s32 (0)); temp.val[1] = vcombine_s32 (b.val[1], vcreate_s32 (0)); temp.val[2] = vcombine_s32 (b.val[2], vcreate_s32 (0)); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) temp.val[2], 2); o = builtin_aarch64_ld3_lanev2si ( (builtin_aarch64_simd_si *) ptr, o, c); b.val[0] = (int32x2_t) builtin_aarch64_get_dregcidi (o, 0); b.val[1] = (int32x2_t) builtin_aarch64_get_dregcidi (o, 1); b.val[2] = (int32x2_t) builtin_aarch64_get_dregcidi (o, 2); return b; }

DEF: int64x1x3 vld3_lane_s64 (const int64_t * ptr, int64x1x3_t b, const int c) { builtin_aarch64_simd_ci o; int64x2x3_t temp; temp.val[0] = vcombine_s64 (b.val[0], vcreate_s64 (0)); temp.val[1] = vcombine_s64 (b.val[1], vcreate_s64 (0)); temp.val[2] = vcombine_s64 (b.val[2], vcreate_s64 (0)); o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) temp.val[2], 2); o = builtin_aarch64_ld3_lanedi ( (builtin_aarch64_simd_di *) ptr, o, c); b.val[0] = (int64x1_t) builtin_aarch64_get_dregcidi (o, 0); b.val[1] = (int64x1_t) builtin_aarch64_get_dregcidi (o, 1); b.val[2] = (int64x1_t) builtin_aarch64_get_dregcidi (o, 2); return b; }

DEF: uint8x8x3 vld3_lane_u8 (const uint8_t * ptr, uint8x8x3_t b, const int c) { builtin_aarch64_simd_ci o; uint8x16x3_t temp; temp.val[0] = vcombine_u8 (b.val[0], vcreate_u8 (0)); temp.val[1] = vcombine_u8 (b.val[1], vcreate_u8 (0)); temp.val[2] = vcombine_u8 (b.val[2], vcreate_u8 (0)); o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[2], 2); o = builtin_aarch64_ld3_lanev8qi ( (builtin_aarch64_simd_qi *) ptr, o, c); b.val[0] = (uint8x8_t) builtin_aarch64_get_dregcidi (o, 0); b.val[1] = (uint8x8_t) builtin_aarch64_get_dregcidi (o, 1); b.val[2] = (uint8x8_t) builtin_aarch64_get_dregcidi (o, 2); return b; }

DEF: uint16x4x3 vld3_lane_u16 (const uint16_t * ptr, uint16x4x3_t b, const int c) { builtin_aarch64_simd_ci o; uint16x8x3_t temp; temp.val[0] = vcombine_u16 (b.val[0], vcreate_u16 (0)); temp.val[1] = vcombine_u16 (b.val[1], vcreate_u16 (0)); temp.val[2] = vcombine_u16 (b.val[2], vcreate_u16 (0)); o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[2], 2); o = builtin_aarch64_ld3_lanev4hi ( (builtin_aarch64_simd_hi *) ptr, o, c); b.val[0] = (uint16x4_t) builtin_aarch64_get_dregcidi (o, 0); b.val[1] = (uint16x4_t) builtin_aarch64_get_dregcidi (o, 1); b.val[2] = (uint16x4_t) builtin_aarch64_get_dregcidi (o, 2); return b; }

DEF: uint32x2x3 vld3_lane_u32 (const uint32_t * ptr, uint32x2x3_t b, const int c) { builtin_aarch64_simd_ci o; uint32x4x3_t temp; temp.val[0] = vcombine_u32 (b.val[0], vcreate_u32 (0)); temp.val[1] = vcombine_u32 (b.val[1], vcreate_u32 (0)); temp.val[2] = vcombine_u32 (b.val[2], vcreate_u32 (0)); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) temp.val[2], 2); o = builtin_aarch64_ld3_lanev2si ( (builtin_aarch64_simd_si *) ptr, o, c); b.val[0] = (uint32x2_t) builtin_aarch64_get_dregcidi (o, 0); b.val[1] = (uint32x2_t) builtin_aarch64_get_dregcidi (o, 1); b.val[2] = (uint32x2_t) builtin_aarch64_get_dregcidi (o, 2); return b; }

DEF: uint64x1x3 vld3_lane_u64 (const uint64_t * ptr, uint64x1x3_t b, const int c) { builtin_aarch64_simd_ci o; uint64x2x3_t temp; temp.val[0] = vcombine_u64 (b.val[0], vcreate_u64 (0)); temp.val[1] = vcombine_u64 (b.val[1], vcreate_u64 (0)); temp.val[2] = vcombine_u64 (b.val[2], vcreate_u64 (0)); o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) temp.val[0], 0); o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) temp.val[1], 1); o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) temp.val[2], 2); o = builtin_aarch64_ld3_lanedi ( (builtin_aarch64_simd_di *) ptr, o, c); b.val[0] = (uint64x1_t) builtin_aarch64_get_dregcidi (o, 0); b.val[1] = (uint64x1_t) builtin_aarch64_get_dregcidi (o, 1); b.val[2] = (uint64x1_t) builtin_aarch64_get_dregcidi (o, 2); return b; }
# 16979 "arm_neon.h"
DEF: float32x4x3 vld3q_lane_f32 (const float32_t * ptr, float32x4x3_t b, const int c) { builtin_aarch64_simd_ci o; float32x4x3_t ret; o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_ld3_lanev4sf ( (builtin_aarch64_simd_sf *) ptr, o, c); ret.val[0] = (float32x4_t) builtin_aarch64_get_qregciv4si (o, 0); ret.val[1] = (float32x4_t) builtin_aarch64_get_qregciv4si (o, 1); ret.val[2] = (float32x4_t) builtin_aarch64_get_qregciv4si (o, 2); return ret; }
DEF: float64x2x3 vld3q_lane_f64 (const float64_t * ptr, float64x2x3_t b, const int c) { builtin_aarch64_simd_ci o; float64x2x3_t ret; o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_ld3_lanev2df ( (builtin_aarch64_simd_df *) ptr, o, c); ret.val[0] = (float64x2_t) builtin_aarch64_get_qregciv4si (o, 0); ret.val[1] = (float64x2_t) builtin_aarch64_get_qregciv4si (o, 1); ret.val[2] = (float64x2_t) builtin_aarch64_get_qregciv4si (o, 2); return ret; }
DEF: poly8x16x3 vld3q_lane_p8 (const poly8_t * ptr, poly8x16x3_t b, const int c) { builtin_aarch64_simd_ci o; poly8x16x3_t ret; o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_ld3_lanev16qi ( (builtin_aarch64_simd_qi *) ptr, o, c); ret.val[0] = (poly8x16_t) builtin_aarch64_get_qregciv4si (o, 0); ret.val[1] = (poly8x16_t) builtin_aarch64_get_qregciv4si (o, 1); ret.val[2] = (poly8x16_t) builtin_aarch64_get_qregciv4si (o, 2); return ret; }
DEF: poly16x8x3 vld3q_lane_p16 (const poly16_t * ptr, poly16x8x3_t b, const int c) { builtin_aarch64_simd_ci o; poly16x8x3_t ret; o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_ld3_lanev8hi ( (builtin_aarch64_simd_hi *) ptr, o, c); ret.val[0] = (poly16x8_t) builtin_aarch64_get_qregciv4si (o, 0); ret.val[1] = (poly16x8_t) builtin_aarch64_get_qregciv4si (o, 1); ret.val[2] = (poly16x8_t) builtin_aarch64_get_qregciv4si (o, 2); return ret; }
DEF: int8x16x3 vld3q_lane_s8 (const int8_t * ptr, int8x16x3_t b, const int c) { builtin_aarch64_simd_ci o; int8x16x3_t ret; o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_ld3_lanev16qi ( (builtin_aarch64_simd_qi *) ptr, o, c); ret.val[0] = (int8x16_t) builtin_aarch64_get_qregciv4si (o, 0); ret.val[1] = (int8x16_t) builtin_aarch64_get_qregciv4si (o, 1); ret.val[2] = (int8x16_t) builtin_aarch64_get_qregciv4si (o, 2); return ret; }
DEF: int16x8x3 vld3q_lane_s16 (const int16_t * ptr, int16x8x3_t b, const int c) { builtin_aarch64_simd_ci o; int16x8x3_t ret; o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_ld3_lanev8hi ( (builtin_aarch64_simd_hi *) ptr, o, c); ret.val[0] = (int16x8_t) builtin_aarch64_get_qregciv4si (o, 0); ret.val[1] = (int16x8_t) builtin_aarch64_get_qregciv4si (o, 1); ret.val[2] = (int16x8_t) builtin_aarch64_get_qregciv4si (o, 2); return ret; }
DEF: int32x4x3 vld3q_lane_s32 (const int32_t * ptr, int32x4x3_t b, const int c) { builtin_aarch64_simd_ci o; int32x4x3_t ret; o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_ld3_lanev4si ( (builtin_aarch64_simd_si *) ptr, o, c); ret.val[0] = (int32x4_t) builtin_aarch64_get_qregciv4si (o, 0); ret.val[1] = (int32x4_t) builtin_aarch64_get_qregciv4si (o, 1); ret.val[2] = (int32x4_t) builtin_aarch64_get_qregciv4si (o, 2); return ret; }
DEF: int64x2x3 vld3q_lane_s64 (const int64_t * ptr, int64x2x3_t b, const int c) { builtin_aarch64_simd_ci o; int64x2x3_t ret; o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_ld3_lanev2di ( (builtin_aarch64_simd_di *) ptr, o, c); ret.val[0] = (int64x2_t) builtin_aarch64_get_qregciv4si (o, 0); ret.val[1] = (int64x2_t) builtin_aarch64_get_qregciv4si (o, 1); ret.val[2] = (int64x2_t) builtin_aarch64_get_qregciv4si (o, 2); return ret; }
DEF: uint8x16x3 vld3q_lane_u8 (const uint8_t * ptr, uint8x16x3_t b, const int c) { builtin_aarch64_simd_ci o; uint8x16x3_t ret; o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_ld3_lanev16qi ( (builtin_aarch64_simd_qi *) ptr, o, c); ret.val[0] = (uint8x16_t) builtin_aarch64_get_qregciv4si (o, 0); ret.val[1] = (uint8x16_t) builtin_aarch64_get_qregciv4si (o, 1); ret.val[2] = (uint8x16_t) builtin_aarch64_get_qregciv4si (o, 2); return ret; }
DEF: uint16x8x3 vld3q_lane_u16 (const uint16_t * ptr, uint16x8x3_t b, const int c) { builtin_aarch64_simd_ci o; uint16x8x3_t ret; o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_ld3_lanev8hi ( (builtin_aarch64_simd_hi *) ptr, o, c); ret.val[0] = (uint16x8_t) builtin_aarch64_get_qregciv4si (o, 0); ret.val[1] = (uint16x8_t) builtin_aarch64_get_qregciv4si (o, 1); ret.val[2] = (uint16x8_t) builtin_aarch64_get_qregciv4si (o, 2); return ret; }
DEF: uint32x4x3 vld3q_lane_u32 (const uint32_t * ptr, uint32x4x3_t b, const int c) { builtin_aarch64_simd_ci o; uint32x4x3_t ret; o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_ld3_lanev4si ( (builtin_aarch64_simd_si *) ptr, o, c); ret.val[0] = (uint32x4_t) builtin_aarch64_get_qregciv4si (o, 0); ret.val[1] = (uint32x4_t) builtin_aarch64_get_qregciv4si (o, 1); ret.val[2] = (uint32x4_t) builtin_aarch64_get_qregciv4si (o, 2); return ret; }
DEF: uint64x2x3 vld3q_lane_u64 (const uint64_t * ptr, uint64x2x3_t b, const int c) { builtin_aarch64_simd_ci o; uint64x2x3_t ret; o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_ld3_lanev2di ( (builtin_aarch64_simd_di *) ptr, o, c); ret.val[0] = (uint64x2_t) builtin_aarch64_get_qregciv4si (o, 0); ret.val[1] = (uint64x2_t) builtin_aarch64_get_qregciv4si (o, 1); ret.val[2] = (uint64x2_t) builtin_aarch64_get_qregciv4si (o, 2); return ret; }
# 17034 "arm_neon.h"
DEF: float32x2x4 vld4_lane_f32 (const float32_t * ptr, float32x2x4_t b, const int c) { builtin_aarch64_simd_xi o; float32x4x4_t temp; temp.val[0] = vcombine_f32 (b.val[0], vcreate_f32 (0)); temp.val[1] = vcombine_f32 (b.val[1], vcreate_f32 (0)); temp.val[2] = vcombine_f32 (b.val[2], vcreate_f32 (0)); temp.val[3] = vcombine_f32 (b.val[3], vcreate_f32 (0)); o = builtin_aarch64_set_qregxiv4sf (o, (float32x4_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv4sf (o, (float32x4_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv4sf (o, (float32x4_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv4sf (o, (float32x4_t) temp.val[3], 3); o = builtin_aarch64_ld4_lanev2sf ( (builtin_aarch64_simd_sf *) ptr, o, c); b.val[0] = (float32x2_t) builtin_aarch64_get_dregxidi (o, 0); b.val[1] = (float32x2_t) builtin_aarch64_get_dregxidi (o, 1); b.val[2] = (float32x2_t) builtin_aarch64_get_dregxidi (o, 2); b.val[3] = (float32x2_t) builtin_aarch64_get_dregxidi (o, 3); return b; }

DEF: float64x1x4 vld4_lane_f64 (const float64_t * ptr, float64x1x4_t b, const int c) { builtin_aarch64_simd_xi o; float64x2x4_t temp; temp.val[0] = vcombine_f64 (b.val[0], vcreate_f64 (0)); temp.val[1] = vcombine_f64 (b.val[1], vcreate_f64 (0)); temp.val[2] = vcombine_f64 (b.val[2], vcreate_f64 (0)); temp.val[3] = vcombine_f64 (b.val[3], vcreate_f64 (0)); o = builtin_aarch64_set_qregxiv2df (o, (float64x2_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv2df (o, (float64x2_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv2df (o, (float64x2_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv2df (o, (float64x2_t) temp.val[3], 3); o = builtin_aarch64_ld4_lanedf ( (builtin_aarch64_simd_df *) ptr, o, c); b.val[0] = (float64x1_t) builtin_aarch64_get_dregxidi (o, 0); b.val[1] = (float64x1_t) builtin_aarch64_get_dregxidi (o, 1); b.val[2] = (float64x1_t) builtin_aarch64_get_dregxidi (o, 2); b.val[3] = (float64x1_t) builtin_aarch64_get_dregxidi (o, 3); return b; }

DEF: poly8x8x4 vld4_lane_p8 (const poly8_t * ptr, poly8x8x4_t b, const int c) { builtin_aarch64_simd_xi o; poly8x16x4_t temp; temp.val[0] = vcombine_p8 (b.val[0], vcreate_p8 (0)); temp.val[1] = vcombine_p8 (b.val[1], vcreate_p8 (0)); temp.val[2] = vcombine_p8 (b.val[2], vcreate_p8 (0)); temp.val[3] = vcombine_p8 (b.val[3], vcreate_p8 (0)); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[3], 3); o = builtin_aarch64_ld4_lanev8qi ( (builtin_aarch64_simd_qi *) ptr, o, c); b.val[0] = (poly8x8_t) builtin_aarch64_get_dregxidi (o, 0); b.val[1] = (poly8x8_t) builtin_aarch64_get_dregxidi (o, 1); b.val[2] = (poly8x8_t) builtin_aarch64_get_dregxidi (o, 2); b.val[3] = (poly8x8_t) builtin_aarch64_get_dregxidi (o, 3); return b; }

DEF: poly16x4x4 vld4_lane_p16 (const poly16_t * ptr, poly16x4x4_t b, const int c) { builtin_aarch64_simd_xi o; poly16x8x4_t temp; temp.val[0] = vcombine_p16 (b.val[0], vcreate_p16 (0)); temp.val[1] = vcombine_p16 (b.val[1], vcreate_p16 (0)); temp.val[2] = vcombine_p16 (b.val[2], vcreate_p16 (0)); temp.val[3] = vcombine_p16 (b.val[3], vcreate_p16 (0)); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[3], 3); o = builtin_aarch64_ld4_lanev4hi ( (builtin_aarch64_simd_hi *) ptr, o, c); b.val[0] = (poly16x4_t) builtin_aarch64_get_dregxidi (o, 0); b.val[1] = (poly16x4_t) builtin_aarch64_get_dregxidi (o, 1); b.val[2] = (poly16x4_t) builtin_aarch64_get_dregxidi (o, 2); b.val[3] = (poly16x4_t) builtin_aarch64_get_dregxidi (o, 3); return b; }

DEF: int8x8x4 vld4_lane_s8 (const int8_t * ptr, int8x8x4_t b, const int c) { builtin_aarch64_simd_xi o; int8x16x4_t temp; temp.val[0] = vcombine_s8 (b.val[0], vcreate_s8 (0)); temp.val[1] = vcombine_s8 (b.val[1], vcreate_s8 (0)); temp.val[2] = vcombine_s8 (b.val[2], vcreate_s8 (0)); temp.val[3] = vcombine_s8 (b.val[3], vcreate_s8 (0)); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[3], 3); o = builtin_aarch64_ld4_lanev8qi ( (builtin_aarch64_simd_qi *) ptr, o, c); b.val[0] = (int8x8_t) builtin_aarch64_get_dregxidi (o, 0); b.val[1] = (int8x8_t) builtin_aarch64_get_dregxidi (o, 1); b.val[2] = (int8x8_t) builtin_aarch64_get_dregxidi (o, 2); b.val[3] = (int8x8_t) builtin_aarch64_get_dregxidi (o, 3); return b; }

DEF: int16x4x4 vld4_lane_s16 (const int16_t * ptr, int16x4x4_t b, const int c) { builtin_aarch64_simd_xi o; int16x8x4_t temp; temp.val[0] = vcombine_s16 (b.val[0], vcreate_s16 (0)); temp.val[1] = vcombine_s16 (b.val[1], vcreate_s16 (0)); temp.val[2] = vcombine_s16 (b.val[2], vcreate_s16 (0)); temp.val[3] = vcombine_s16 (b.val[3], vcreate_s16 (0)); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[3], 3); o = builtin_aarch64_ld4_lanev4hi ( (builtin_aarch64_simd_hi *) ptr, o, c); b.val[0] = (int16x4_t) builtin_aarch64_get_dregxidi (o, 0); b.val[1] = (int16x4_t) builtin_aarch64_get_dregxidi (o, 1); b.val[2] = (int16x4_t) builtin_aarch64_get_dregxidi (o, 2); b.val[3] = (int16x4_t) builtin_aarch64_get_dregxidi (o, 3); return b; }

DEF: int32x2x4 vld4_lane_s32 (const int32_t * ptr, int32x2x4_t b, const int c) { builtin_aarch64_simd_xi o; int32x4x4_t temp; temp.val[0] = vcombine_s32 (b.val[0], vcreate_s32 (0)); temp.val[1] = vcombine_s32 (b.val[1], vcreate_s32 (0)); temp.val[2] = vcombine_s32 (b.val[2], vcreate_s32 (0)); temp.val[3] = vcombine_s32 (b.val[3], vcreate_s32 (0)); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[3], 3); o = builtin_aarch64_ld4_lanev2si ( (builtin_aarch64_simd_si *) ptr, o, c); b.val[0] = (int32x2_t) builtin_aarch64_get_dregxidi (o, 0); b.val[1] = (int32x2_t) builtin_aarch64_get_dregxidi (o, 1); b.val[2] = (int32x2_t) builtin_aarch64_get_dregxidi (o, 2); b.val[3] = (int32x2_t) builtin_aarch64_get_dregxidi (o, 3); return b; }

DEF: int64x1x4 vld4_lane_s64 (const int64_t * ptr, int64x1x4_t b, const int c) { builtin_aarch64_simd_xi o; int64x2x4_t temp; temp.val[0] = vcombine_s64 (b.val[0], vcreate_s64 (0)); temp.val[1] = vcombine_s64 (b.val[1], vcreate_s64 (0)); temp.val[2] = vcombine_s64 (b.val[2], vcreate_s64 (0)); temp.val[3] = vcombine_s64 (b.val[3], vcreate_s64 (0)); o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[3], 3); o = builtin_aarch64_ld4_lanedi ( (builtin_aarch64_simd_di *) ptr, o, c); b.val[0] = (int64x1_t) builtin_aarch64_get_dregxidi (o, 0); b.val[1] = (int64x1_t) builtin_aarch64_get_dregxidi (o, 1); b.val[2] = (int64x1_t) builtin_aarch64_get_dregxidi (o, 2); b.val[3] = (int64x1_t) builtin_aarch64_get_dregxidi (o, 3); return b; }

DEF: uint8x8x4 vld4_lane_u8 (const uint8_t * ptr, uint8x8x4_t b, const int c) { builtin_aarch64_simd_xi o; uint8x16x4_t temp; temp.val[0] = vcombine_u8 (b.val[0], vcreate_u8 (0)); temp.val[1] = vcombine_u8 (b.val[1], vcreate_u8 (0)); temp.val[2] = vcombine_u8 (b.val[2], vcreate_u8 (0)); temp.val[3] = vcombine_u8 (b.val[3], vcreate_u8 (0)); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[3], 3); o = builtin_aarch64_ld4_lanev8qi ( (builtin_aarch64_simd_qi *) ptr, o, c); b.val[0] = (uint8x8_t) builtin_aarch64_get_dregxidi (o, 0); b.val[1] = (uint8x8_t) builtin_aarch64_get_dregxidi (o, 1); b.val[2] = (uint8x8_t) builtin_aarch64_get_dregxidi (o, 2); b.val[3] = (uint8x8_t) builtin_aarch64_get_dregxidi (o, 3); return b; }

DEF: uint16x4x4 vld4_lane_u16 (const uint16_t * ptr, uint16x4x4_t b, const int c) { builtin_aarch64_simd_xi o; uint16x8x4_t temp; temp.val[0] = vcombine_u16 (b.val[0], vcreate_u16 (0)); temp.val[1] = vcombine_u16 (b.val[1], vcreate_u16 (0)); temp.val[2] = vcombine_u16 (b.val[2], vcreate_u16 (0)); temp.val[3] = vcombine_u16 (b.val[3], vcreate_u16 (0)); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[3], 3); o = builtin_aarch64_ld4_lanev4hi ( (builtin_aarch64_simd_hi *) ptr, o, c); b.val[0] = (uint16x4_t) builtin_aarch64_get_dregxidi (o, 0); b.val[1] = (uint16x4_t) builtin_aarch64_get_dregxidi (o, 1); b.val[2] = (uint16x4_t) builtin_aarch64_get_dregxidi (o, 2); b.val[3] = (uint16x4_t) builtin_aarch64_get_dregxidi (o, 3); return b; }

DEF: uint32x2x4 vld4_lane_u32 (const uint32_t * ptr, uint32x2x4_t b, const int c) { builtin_aarch64_simd_xi o; uint32x4x4_t temp; temp.val[0] = vcombine_u32 (b.val[0], vcreate_u32 (0)); temp.val[1] = vcombine_u32 (b.val[1], vcreate_u32 (0)); temp.val[2] = vcombine_u32 (b.val[2], vcreate_u32 (0)); temp.val[3] = vcombine_u32 (b.val[3], vcreate_u32 (0)); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[3], 3); o = builtin_aarch64_ld4_lanev2si ( (builtin_aarch64_simd_si *) ptr, o, c); b.val[0] = (uint32x2_t) builtin_aarch64_get_dregxidi (o, 0); b.val[1] = (uint32x2_t) builtin_aarch64_get_dregxidi (o, 1); b.val[2] = (uint32x2_t) builtin_aarch64_get_dregxidi (o, 2); b.val[3] = (uint32x2_t) builtin_aarch64_get_dregxidi (o, 3); return b; }

DEF: uint64x1x4 vld4_lane_u64 (const uint64_t * ptr, uint64x1x4_t b, const int c) { builtin_aarch64_simd_xi o; uint64x2x4_t temp; temp.val[0] = vcombine_u64 (b.val[0], vcreate_u64 (0)); temp.val[1] = vcombine_u64 (b.val[1], vcreate_u64 (0)); temp.val[2] = vcombine_u64 (b.val[2], vcreate_u64 (0)); temp.val[3] = vcombine_u64 (b.val[3], vcreate_u64 (0)); o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[0], 0); o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[1], 1); o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[2], 2); o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[3], 3); o = builtin_aarch64_ld4_lanedi ( (builtin_aarch64_simd_di *) ptr, o, c); b.val[0] = (uint64x1_t) builtin_aarch64_get_dregxidi (o, 0); b.val[1] = (uint64x1_t) builtin_aarch64_get_dregxidi (o, 1); b.val[2] = (uint64x1_t) builtin_aarch64_get_dregxidi (o, 2); b.val[3] = (uint64x1_t) builtin_aarch64_get_dregxidi (o, 3); return b; }
# 17082 "arm_neon.h"
DEF: float32x4x4 vld4q_lane_f32 (const float32_t * ptr, float32x4x4_t b, const int c) { builtin_aarch64_simd_xi o; float32x4x4_t ret; o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[3], 3); o = builtin_aarch64_ld4_lanev4sf ( (builtin_aarch64_simd_sf *) ptr, o, c); ret.val[0] = (float32x4_t) builtin_aarch64_get_qregxiv4si (o, 0); ret.val[1] = (float32x4_t) builtin_aarch64_get_qregxiv4si (o, 1); ret.val[2] = (float32x4_t) builtin_aarch64_get_qregxiv4si (o, 2); ret.val[3] = (float32x4_t) builtin_aarch64_get_qregxiv4si (o, 3); return ret; }
DEF: float64x2x4 vld4q_lane_f64 (const float64_t * ptr, float64x2x4_t b, const int c) { builtin_aarch64_simd_xi o; float64x2x4_t ret; o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[3], 3); o = builtin_aarch64_ld4_lanev2df ( (builtin_aarch64_simd_df *) ptr, o, c); ret.val[0] = (float64x2_t) builtin_aarch64_get_qregxiv4si (o, 0); ret.val[1] = (float64x2_t) builtin_aarch64_get_qregxiv4si (o, 1); ret.val[2] = (float64x2_t) builtin_aarch64_get_qregxiv4si (o, 2); ret.val[3] = (float64x2_t) builtin_aarch64_get_qregxiv4si (o, 3); return ret; }
DEF: poly8x16x4 vld4q_lane_p8 (const poly8_t * ptr, poly8x16x4_t b, const int c) { builtin_aarch64_simd_xi o; poly8x16x4_t ret; o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[3], 3); o = builtin_aarch64_ld4_lanev16qi ( (builtin_aarch64_simd_qi *) ptr, o, c); ret.val[0] = (poly8x16_t) builtin_aarch64_get_qregxiv4si (o, 0); ret.val[1] = (poly8x16_t) builtin_aarch64_get_qregxiv4si (o, 1); ret.val[2] = (poly8x16_t) builtin_aarch64_get_qregxiv4si (o, 2); ret.val[3] = (poly8x16_t) builtin_aarch64_get_qregxiv4si (o, 3); return ret; }
DEF: poly16x8x4 vld4q_lane_p16 (const poly16_t * ptr, poly16x8x4_t b, const int c) { builtin_aarch64_simd_xi o; poly16x8x4_t ret; o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[3], 3); o = builtin_aarch64_ld4_lanev8hi ( (builtin_aarch64_simd_hi *) ptr, o, c); ret.val[0] = (poly16x8_t) builtin_aarch64_get_qregxiv4si (o, 0); ret.val[1] = (poly16x8_t) builtin_aarch64_get_qregxiv4si (o, 1); ret.val[2] = (poly16x8_t) builtin_aarch64_get_qregxiv4si (o, 2); ret.val[3] = (poly16x8_t) builtin_aarch64_get_qregxiv4si (o, 3); return ret; }
DEF: int8x16x4 vld4q_lane_s8 (const int8_t * ptr, int8x16x4_t b, const int c) { builtin_aarch64_simd_xi o; int8x16x4_t ret; o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[3], 3); o = builtin_aarch64_ld4_lanev16qi ( (builtin_aarch64_simd_qi *) ptr, o, c); ret.val[0] = (int8x16_t) builtin_aarch64_get_qregxiv4si (o, 0); ret.val[1] = (int8x16_t) builtin_aarch64_get_qregxiv4si (o, 1); ret.val[2] = (int8x16_t) builtin_aarch64_get_qregxiv4si (o, 2); ret.val[3] = (int8x16_t) builtin_aarch64_get_qregxiv4si (o, 3); return ret; }
DEF: int16x8x4 vld4q_lane_s16 (const int16_t * ptr, int16x8x4_t b, const int c) { builtin_aarch64_simd_xi o; int16x8x4_t ret; o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[3], 3); o = builtin_aarch64_ld4_lanev8hi ( (builtin_aarch64_simd_hi *) ptr, o, c); ret.val[0] = (int16x8_t) builtin_aarch64_get_qregxiv4si (o, 0); ret.val[1] = (int16x8_t) builtin_aarch64_get_qregxiv4si (o, 1); ret.val[2] = (int16x8_t) builtin_aarch64_get_qregxiv4si (o, 2); ret.val[3] = (int16x8_t) builtin_aarch64_get_qregxiv4si (o, 3); return ret; }
DEF: int32x4x4 vld4q_lane_s32 (const int32_t * ptr, int32x4x4_t b, const int c) { builtin_aarch64_simd_xi o; int32x4x4_t ret; o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[3], 3); o = builtin_aarch64_ld4_lanev4si ( (builtin_aarch64_simd_si *) ptr, o, c); ret.val[0] = (int32x4_t) builtin_aarch64_get_qregxiv4si (o, 0); ret.val[1] = (int32x4_t) builtin_aarch64_get_qregxiv4si (o, 1); ret.val[2] = (int32x4_t) builtin_aarch64_get_qregxiv4si (o, 2); ret.val[3] = (int32x4_t) builtin_aarch64_get_qregxiv4si (o, 3); return ret; }
DEF: int64x2x4 vld4q_lane_s64 (const int64_t * ptr, int64x2x4_t b, const int c) { builtin_aarch64_simd_xi o; int64x2x4_t ret; o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[3], 3); o = builtin_aarch64_ld4_lanev2di ( (builtin_aarch64_simd_di *) ptr, o, c); ret.val[0] = (int64x2_t) builtin_aarch64_get_qregxiv4si (o, 0); ret.val[1] = (int64x2_t) builtin_aarch64_get_qregxiv4si (o, 1); ret.val[2] = (int64x2_t) builtin_aarch64_get_qregxiv4si (o, 2); ret.val[3] = (int64x2_t) builtin_aarch64_get_qregxiv4si (o, 3); return ret; }
DEF: uint8x16x4 vld4q_lane_u8 (const uint8_t * ptr, uint8x16x4_t b, const int c) { builtin_aarch64_simd_xi o; uint8x16x4_t ret; o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[3], 3); o = builtin_aarch64_ld4_lanev16qi ( (builtin_aarch64_simd_qi *) ptr, o, c); ret.val[0] = (uint8x16_t) builtin_aarch64_get_qregxiv4si (o, 0); ret.val[1] = (uint8x16_t) builtin_aarch64_get_qregxiv4si (o, 1); ret.val[2] = (uint8x16_t) builtin_aarch64_get_qregxiv4si (o, 2); ret.val[3] = (uint8x16_t) builtin_aarch64_get_qregxiv4si (o, 3); return ret; }
DEF: uint16x8x4 vld4q_lane_u16 (const uint16_t * ptr, uint16x8x4_t b, const int c) { builtin_aarch64_simd_xi o; uint16x8x4_t ret; o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[3], 3); o = builtin_aarch64_ld4_lanev8hi ( (builtin_aarch64_simd_hi *) ptr, o, c); ret.val[0] = (uint16x8_t) builtin_aarch64_get_qregxiv4si (o, 0); ret.val[1] = (uint16x8_t) builtin_aarch64_get_qregxiv4si (o, 1); ret.val[2] = (uint16x8_t) builtin_aarch64_get_qregxiv4si (o, 2); ret.val[3] = (uint16x8_t) builtin_aarch64_get_qregxiv4si (o, 3); return ret; }
DEF: uint32x4x4 vld4q_lane_u32 (const uint32_t * ptr, uint32x4x4_t b, const int c) { builtin_aarch64_simd_xi o; uint32x4x4_t ret; o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[3], 3); o = builtin_aarch64_ld4_lanev4si ( (builtin_aarch64_simd_si *) ptr, o, c); ret.val[0] = (uint32x4_t) builtin_aarch64_get_qregxiv4si (o, 0); ret.val[1] = (uint32x4_t) builtin_aarch64_get_qregxiv4si (o, 1); ret.val[2] = (uint32x4_t) builtin_aarch64_get_qregxiv4si (o, 2); ret.val[3] = (uint32x4_t) builtin_aarch64_get_qregxiv4si (o, 3); return ret; }
DEF: uint64x2x4 vld4q_lane_u64 (const uint64_t * ptr, uint64x2x4_t b, const int c) { builtin_aarch64_simd_xi o; uint64x2x4_t ret; o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[0], 0); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[1], 1); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[2], 2); o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) b.val[3], 3); o = builtin_aarch64_ld4_lanev2di ( (builtin_aarch64_simd_di *) ptr, o, c); ret.val[0] = (uint64x2_t) builtin_aarch64_get_qregxiv4si (o, 0); ret.val[1] = (uint64x2_t) builtin_aarch64_get_qregxiv4si (o, 1); ret.val[2] = (uint64x2_t) builtin_aarch64_get_qregxiv4si (o, 2); ret.val[3] = (uint64x2_t) builtin_aarch64_get_qregxiv4si (o, 3); return ret; }

DEF: float32x2
vmax_f32 (float32x2_t a, float32x2_t b)
{
  return builtin_aarch64_smax_nanv2sf (a, b);
}

DEF: int8x8
vmax_s8 (int8x8_t a, int8x8_t b)
{
  return builtin_aarch64_smaxv8qi (a, b);
}

DEF: int16x4
vmax_s16 (int16x4_t a, int16x4_t b)
{
  return builtin_aarch64_smaxv4hi (a, b);
}

DEF: int32x2
vmax_s32 (int32x2_t a, int32x2_t b)
{
  return builtin_aarch64_smaxv2si (a, b);
}

DEF: uint8x8
vmax_u8 (uint8x8_t a, uint8x8_t b)
{
  return (uint8x8_t) builtin_aarch64_umaxv8qi ((int8x8_t) a,
       (int8x8_t) b);
}

DEF: uint16x4
vmax_u16 (uint16x4_t a, uint16x4_t b)
{
  return (uint16x4_t) builtin_aarch64_umaxv4hi ((int16x4_t) a,
        (int16x4_t) b);
}

DEF: uint32x2
vmax_u32 (uint32x2_t a, uint32x2_t b)
{
  return (uint32x2_t) builtin_aarch64_umaxv2si ((int32x2_t) a,
        (int32x2_t) b);
}

DEF: float32x4
vmaxq_f32 (float32x4_t a, float32x4_t b)
{
  return builtin_aarch64_smax_nanv4sf (a, b);
}

DEF: float64x2
vmaxq_f64 (float64x2_t a, float64x2_t b)
{
  return builtin_aarch64_smax_nanv2df (a, b);
}

DEF: int8x16
vmaxq_s8 (int8x16_t a, int8x16_t b)
{
  return builtin_aarch64_smaxv16qi (a, b);
}

DEF: int16x8
vmaxq_s16 (int16x8_t a, int16x8_t b)
{
  return builtin_aarch64_smaxv8hi (a, b);
}

DEF: int32x4
vmaxq_s32 (int32x4_t a, int32x4_t b)
{
  return builtin_aarch64_smaxv4si (a, b);
}

DEF: uint8x16
vmaxq_u8 (uint8x16_t a, uint8x16_t b)
{
  return (uint8x16_t) builtin_aarch64_umaxv16qi ((int8x16_t) a,
         (int8x16_t) b);
}

DEF: uint16x8
vmaxq_u16 (uint16x8_t a, uint16x8_t b)
{
  return (uint16x8_t) builtin_aarch64_umaxv8hi ((int16x8_t) a,
        (int16x8_t) b);
}

DEF: uint32x4
vmaxq_u32 (uint32x4_t a, uint32x4_t b)
{
  return (uint32x4_t) builtin_aarch64_umaxv4si ((int32x4_t) a,
        (int32x4_t) b);
}

DEF: int8x8
vpmax_s8 (int8x8_t a, int8x8_t b)
{
  return builtin_aarch64_smaxpv8qi (a, b);
}

DEF: int16x4
vpmax_s16 (int16x4_t a, int16x4_t b)
{
  return builtin_aarch64_smaxpv4hi (a, b);
}

DEF: int32x2
vpmax_s32 (int32x2_t a, int32x2_t b)
{
  return builtin_aarch64_smaxpv2si (a, b);
}

DEF: uint8x8
vpmax_u8 (uint8x8_t a, uint8x8_t b)
{
  return (uint8x8_t) builtin_aarch64_umaxpv8qi ((int8x8_t) a,
        (int8x8_t) b);
}

DEF: uint16x4
vpmax_u16 (uint16x4_t a, uint16x4_t b)
{
  return (uint16x4_t) builtin_aarch64_umaxpv4hi ((int16x4_t) a,
         (int16x4_t) b);
}

DEF: uint32x2
vpmax_u32 (uint32x2_t a, uint32x2_t b)
{
  return (uint32x2_t) builtin_aarch64_umaxpv2si ((int32x2_t) a,
         (int32x2_t) b);
}

DEF: int8x16
vpmaxq_s8 (int8x16_t a, int8x16_t b)
{
  return builtin_aarch64_smaxpv16qi (a, b);
}

DEF: int16x8
vpmaxq_s16 (int16x8_t a, int16x8_t b)
{
  return builtin_aarch64_smaxpv8hi (a, b);
}

DEF: int32x4
vpmaxq_s32 (int32x4_t a, int32x4_t b)
{
  return builtin_aarch64_smaxpv4si (a, b);
}

DEF: uint8x16
vpmaxq_u8 (uint8x16_t a, uint8x16_t b)
{
  return (uint8x16_t) builtin_aarch64_umaxpv16qi ((int8x16_t) a,
          (int8x16_t) b);
}

DEF: uint16x8
vpmaxq_u16 (uint16x8_t a, uint16x8_t b)
{
  return (uint16x8_t) builtin_aarch64_umaxpv8hi ((int16x8_t) a,
         (int16x8_t) b);
}

DEF: uint32x4
vpmaxq_u32 (uint32x4_t a, uint32x4_t b)
{
  return (uint32x4_t) builtin_aarch64_umaxpv4si ((int32x4_t) a,
         (int32x4_t) b);
}

DEF: float32x2
vpmax_f32 (float32x2_t a, float32x2_t b)
{
  return builtin_aarch64_smax_nanpv2sf (a, b);
}

DEF: float32x4
vpmaxq_f32 (float32x4_t a, float32x4_t b)
{
  return builtin_aarch64_smax_nanpv4sf (a, b);
}

DEF: float64x2
vpmaxq_f64 (float64x2_t a, float64x2_t b)
{
  return builtin_aarch64_smax_nanpv2df (a, b);
}

DEF: float64
vpmaxqd_f64 (float64x2_t a)
{
  return builtin_aarch64_reduc_smax_nan_scal_v2df (a);
}

DEF: float32
vpmaxs_f32 (float32x2_t a)
{
  return builtin_aarch64_reduc_smax_nan_scal_v2sf (a);
}

DEF: float32x2
vpmaxnm_f32 (float32x2_t a, float32x2_t b)
{
  return builtin_aarch64_smaxpv2sf (a, b);
}

DEF: float32x4
vpmaxnmq_f32 (float32x4_t a, float32x4_t b)
{
  return builtin_aarch64_smaxpv4sf (a, b);
}

DEF: float64x2
vpmaxnmq_f64 (float64x2_t a, float64x2_t b)
{
  return builtin_aarch64_smaxpv2df (a, b);
}

DEF: float64
vpmaxnmqd_f64 (float64x2_t a)
{
  return builtin_aarch64_reduc_smax_scal_v2df (a);
}

DEF: float32
vpmaxnms_f32 (float32x2_t a)
{
  return builtin_aarch64_reduc_smax_scal_v2sf (a);
}

DEF: int8x8
vpmin_s8 (int8x8_t a, int8x8_t b)
{
  return builtin_aarch64_sminpv8qi (a, b);
}

DEF: int16x4
vpmin_s16 (int16x4_t a, int16x4_t b)
{
  return builtin_aarch64_sminpv4hi (a, b);
}

DEF: int32x2
vpmin_s32 (int32x2_t a, int32x2_t b)
{
  return builtin_aarch64_sminpv2si (a, b);
}

DEF: uint8x8
vpmin_u8 (uint8x8_t a, uint8x8_t b)
{
  return (uint8x8_t) builtin_aarch64_uminpv8qi ((int8x8_t) a,
        (int8x8_t) b);
}

DEF: uint16x4
vpmin_u16 (uint16x4_t a, uint16x4_t b)
{
  return (uint16x4_t) builtin_aarch64_uminpv4hi ((int16x4_t) a,
         (int16x4_t) b);
}

DEF: uint32x2
vpmin_u32 (uint32x2_t a, uint32x2_t b)
{
  return (uint32x2_t) builtin_aarch64_uminpv2si ((int32x2_t) a,
         (int32x2_t) b);
}

DEF: int8x16
vpminq_s8 (int8x16_t a, int8x16_t b)
{
  return builtin_aarch64_sminpv16qi (a, b);
}

DEF: int16x8
vpminq_s16 (int16x8_t a, int16x8_t b)
{
  return builtin_aarch64_sminpv8hi (a, b);
}

DEF: int32x4
vpminq_s32 (int32x4_t a, int32x4_t b)
{
  return builtin_aarch64_sminpv4si (a, b);
}

DEF: uint8x16
vpminq_u8 (uint8x16_t a, uint8x16_t b)
{
  return (uint8x16_t) builtin_aarch64_uminpv16qi ((int8x16_t) a,
          (int8x16_t) b);
}

DEF: uint16x8
vpminq_u16 (uint16x8_t a, uint16x8_t b)
{
  return (uint16x8_t) builtin_aarch64_uminpv8hi ((int16x8_t) a,
         (int16x8_t) b);
}

DEF: uint32x4
vpminq_u32 (uint32x4_t a, uint32x4_t b)
{
  return (uint32x4_t) builtin_aarch64_uminpv4si ((int32x4_t) a,
         (int32x4_t) b);
}

DEF: float32x2
vpmin_f32 (float32x2_t a, float32x2_t b)
{
  return builtin_aarch64_smin_nanpv2sf (a, b);
}

DEF: float32x4
vpminq_f32 (float32x4_t a, float32x4_t b)
{
  return builtin_aarch64_smin_nanpv4sf (a, b);
}

DEF: float64x2
vpminq_f64 (float64x2_t a, float64x2_t b)
{
  return builtin_aarch64_smin_nanpv2df (a, b);
}

DEF: float64
vpminqd_f64 (float64x2_t a)
{
  return builtin_aarch64_reduc_smin_nan_scal_v2df (a);
}

DEF: float32
vpmins_f32 (float32x2_t a)
{
  return builtin_aarch64_reduc_smin_nan_scal_v2sf (a);
}

DEF: float32x2
vpminnm_f32 (float32x2_t a, float32x2_t b)
{
  return builtin_aarch64_sminpv2sf (a, b);
}

DEF: float32x4
vpminnmq_f32 (float32x4_t a, float32x4_t b)
{
  return builtin_aarch64_sminpv4sf (a, b);
}

DEF: float64x2
vpminnmq_f64 (float64x2_t a, float64x2_t b)
{
  return builtin_aarch64_sminpv2df (a, b);
}

DEF: float64
vpminnmqd_f64 (float64x2_t a)
{
  return builtin_aarch64_reduc_smin_scal_v2df (a);
}

DEF: float32
vpminnms_f32 (float32x2_t a)
{
  return builtin_aarch64_reduc_smin_scal_v2sf (a);
}

DEF: float32x2
vmaxnm_f32 (float32x2_t a, float32x2_t b)
{
  return builtin_aarch64_smaxv2sf (a, b);
}

DEF: float32x4
vmaxnmq_f32 (float32x4_t a, float32x4_t b)
{
  return builtin_aarch64_smaxv4sf (a, b);
}

DEF: float64x2
vmaxnmq_f64 (float64x2_t a, float64x2_t b)
{
  return builtin_aarch64_smaxv2df (a, b);
}

DEF: float32
vmaxv_f32 (float32x2_t a)
{
  return builtin_aarch64_reduc_smax_nan_scal_v2sf (a);
}

DEF: int8
vmaxv_s8 (int8x8_t a)
{
  return builtin_aarch64_reduc_smax_scal_v8qi (a);
}

DEF: int16
vmaxv_s16 (int16x4_t a)
{
  return builtin_aarch64_reduc_smax_scal_v4hi (a);
}

DEF: int32
vmaxv_s32 (int32x2_t a)
{
  return builtin_aarch64_reduc_smax_scal_v2si (a);
}

DEF: uint8
vmaxv_u8 (uint8x8_t a)
{
  return builtin_aarch64_reduc_umax_scal_v8qi_uu (a);
}

DEF: uint16
vmaxv_u16 (uint16x4_t a)
{
  return builtin_aarch64_reduc_umax_scal_v4hi_uu (a);
}

DEF: uint32
vmaxv_u32 (uint32x2_t a)
{
  return builtin_aarch64_reduc_umax_scal_v2si_uu (a);
}

DEF: float32
vmaxvq_f32 (float32x4_t a)
{
  return builtin_aarch64_reduc_smax_nan_scal_v4sf (a);
}

DEF: float64
vmaxvq_f64 (float64x2_t a)
{
  return builtin_aarch64_reduc_smax_nan_scal_v2df (a);
}

DEF: int8
vmaxvq_s8 (int8x16_t a)
{
  return builtin_aarch64_reduc_smax_scal_v16qi (a);
}

DEF: int16
vmaxvq_s16 (int16x8_t a)
{
  return builtin_aarch64_reduc_smax_scal_v8hi (a);
}

DEF: int32
vmaxvq_s32 (int32x4_t a)
{
  return builtin_aarch64_reduc_smax_scal_v4si (a);
}

DEF: uint8
vmaxvq_u8 (uint8x16_t a)
{
  return builtin_aarch64_reduc_umax_scal_v16qi_uu (a);
}

DEF: uint16
vmaxvq_u16 (uint16x8_t a)
{
  return builtin_aarch64_reduc_umax_scal_v8hi_uu (a);
}

DEF: uint32
vmaxvq_u32 (uint32x4_t a)
{
  return builtin_aarch64_reduc_umax_scal_v4si_uu (a);
}

DEF: float32
vmaxnmv_f32 (float32x2_t a)
{
  return builtin_aarch64_reduc_smax_scal_v2sf (a);
}

DEF: float32
vmaxnmvq_f32 (float32x4_t a)
{
  return builtin_aarch64_reduc_smax_scal_v4sf (a);
}

DEF: float64
vmaxnmvq_f64 (float64x2_t a)
{
  return builtin_aarch64_reduc_smax_scal_v2df (a);
}

DEF: float32x2
vmin_f32 (float32x2_t a, float32x2_t b)
{
  return builtin_aarch64_smin_nanv2sf (a, b);
}

DEF: int8x8
vmin_s8 (int8x8_t a, int8x8_t b)
{
  return builtin_aarch64_sminv8qi (a, b);
}

DEF: int16x4
vmin_s16 (int16x4_t a, int16x4_t b)
{
  return builtin_aarch64_sminv4hi (a, b);
}

DEF: int32x2
vmin_s32 (int32x2_t a, int32x2_t b)
{
  return builtin_aarch64_sminv2si (a, b);
}

DEF: uint8x8
vmin_u8 (uint8x8_t a, uint8x8_t b)
{
  return (uint8x8_t) builtin_aarch64_uminv8qi ((int8x8_t) a,
       (int8x8_t) b);
}

DEF: uint16x4
vmin_u16 (uint16x4_t a, uint16x4_t b)
{
  return (uint16x4_t) builtin_aarch64_uminv4hi ((int16x4_t) a,
        (int16x4_t) b);
}

DEF: uint32x2
vmin_u32 (uint32x2_t a, uint32x2_t b)
{
  return (uint32x2_t) builtin_aarch64_uminv2si ((int32x2_t) a,
        (int32x2_t) b);
}

DEF: float32x4
vminq_f32 (float32x4_t a, float32x4_t b)
{
  return builtin_aarch64_smin_nanv4sf (a, b);
}

DEF: float64x2
vminq_f64 (float64x2_t a, float64x2_t b)
{
  return builtin_aarch64_smin_nanv2df (a, b);
}

DEF: int8x16
vminq_s8 (int8x16_t a, int8x16_t b)
{
  return builtin_aarch64_sminv16qi (a, b);
}

DEF: int16x8
vminq_s16 (int16x8_t a, int16x8_t b)
{
  return builtin_aarch64_sminv8hi (a, b);
}

DEF: int32x4
vminq_s32 (int32x4_t a, int32x4_t b)
{
  return builtin_aarch64_sminv4si (a, b);
}

DEF: uint8x16
vminq_u8 (uint8x16_t a, uint8x16_t b)
{
  return (uint8x16_t) builtin_aarch64_uminv16qi ((int8x16_t) a,
         (int8x16_t) b);
}

DEF: uint16x8
vminq_u16 (uint16x8_t a, uint16x8_t b)
{
  return (uint16x8_t) builtin_aarch64_uminv8hi ((int16x8_t) a,
        (int16x8_t) b);
}

DEF: uint32x4
vminq_u32 (uint32x4_t a, uint32x4_t b)
{
  return (uint32x4_t) builtin_aarch64_uminv4si ((int32x4_t) a,
        (int32x4_t) b);
}

DEF: float32x2
vminnm_f32 (float32x2_t a, float32x2_t b)
{
  return builtin_aarch64_sminv2sf (a, b);
}

DEF: float32x4
vminnmq_f32 (float32x4_t a, float32x4_t b)
{
  return builtin_aarch64_sminv4sf (a, b);
}

DEF: float64x2
vminnmq_f64 (float64x2_t a, float64x2_t b)
{
  return builtin_aarch64_sminv2df (a, b);
}

DEF: float32
vminv_f32 (float32x2_t a)
{
  return builtin_aarch64_reduc_smin_nan_scal_v2sf (a);
}

DEF: int8
vminv_s8 (int8x8_t a)
{
  return builtin_aarch64_reduc_smin_scal_v8qi (a);
}

DEF: int16
vminv_s16 (int16x4_t a)
{
  return builtin_aarch64_reduc_smin_scal_v4hi (a);
}

DEF: int32
vminv_s32 (int32x2_t a)
{
  return builtin_aarch64_reduc_smin_scal_v2si (a);
}

DEF: uint8
vminv_u8 (uint8x8_t a)
{
  return builtin_aarch64_reduc_umin_scal_v8qi_uu (a);
}

DEF: uint16
vminv_u16 (uint16x4_t a)
{
  return builtin_aarch64_reduc_umin_scal_v4hi_uu (a);
}

DEF: uint32
vminv_u32 (uint32x2_t a)
{
  return builtin_aarch64_reduc_umin_scal_v2si_uu (a);
}

DEF: float32
vminvq_f32 (float32x4_t a)
{
  return builtin_aarch64_reduc_smin_nan_scal_v4sf (a);
}

DEF: float64
vminvq_f64 (float64x2_t a)
{
  return builtin_aarch64_reduc_smin_nan_scal_v2df (a);
}

DEF: int8
vminvq_s8 (int8x16_t a)
{
  return builtin_aarch64_reduc_smin_scal_v16qi (a);
}

DEF: int16
vminvq_s16 (int16x8_t a)
{
  return builtin_aarch64_reduc_smin_scal_v8hi (a);
}

DEF: int32
vminvq_s32 (int32x4_t a)
{
  return builtin_aarch64_reduc_smin_scal_v4si (a);
}

DEF: uint8
vminvq_u8 (uint8x16_t a)
{
  return builtin_aarch64_reduc_umin_scal_v16qi_uu (a);
}

DEF: uint16
vminvq_u16 (uint16x8_t a)
{
  return builtin_aarch64_reduc_umin_scal_v8hi_uu (a);
}

DEF: uint32
vminvq_u32 (uint32x4_t a)
{
  return builtin_aarch64_reduc_umin_scal_v4si_uu (a);
}

DEF: float32
vminnmv_f32 (float32x2_t a)
{
  return builtin_aarch64_reduc_smin_scal_v2sf (a);
}

DEF: float32
vminnmvq_f32 (float32x4_t a)
{
  return builtin_aarch64_reduc_smin_scal_v4sf (a);
}

DEF: float64
vminnmvq_f64 (float64x2_t a)
{
  return builtin_aarch64_reduc_smin_scal_v2df (a);
}

DEF: float32x2
vmla_f32 (float32x2_t a, float32x2_t b, float32x2_t c)
{
  return a + b * c;
}

DEF: float64x1
vmla_f64 (float64x1_t a, float64x1_t b, float64x1_t c)
{
  return a + b * c;
}

DEF: float32x4
vmlaq_f32 (float32x4_t a, float32x4_t b, float32x4_t c)
{
  return a + b * c;
}

DEF: float64x2
vmlaq_f64 (float64x2_t a, float64x2_t b, float64x2_t c)
{
  return a + b * c;
}

DEF: float32x2
vmla_lane_f32 (float32x2_t a, float32x2_t b,
        float32x2_t c, const int lane)
{
  return (a + (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: int16x4
vmla_lane_s16 (int16x4_t a, int16x4_t b,
  int16x4_t c, const int lane)
{
  return (a + (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: int32x2
vmla_lane_s32 (int32x2_t a, int32x2_t b,
  int32x2_t c, const int lane)
{
  return (a + (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: uint16x4
vmla_lane_u16 (uint16x4_t a, uint16x4_t b,
  uint16x4_t c, const int lane)
{
  return (a + (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: uint32x2
vmla_lane_u32 (uint32x2_t a, uint32x2_t b,
        uint32x2_t c, const int lane)
{
  return (a + (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: float32x2
vmla_laneq_f32 (float32x2_t a, float32x2_t b,
         float32x4_t c, const int lane)
{
  return (a + (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: int16x4
vmla_laneq_s16 (int16x4_t a, int16x4_t b,
  int16x8_t c, const int lane)
{
  return (a + (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: int32x2
vmla_laneq_s32 (int32x2_t a, int32x2_t b,
  int32x4_t c, const int lane)
{
  return (a + (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: uint16x4
vmla_laneq_u16 (uint16x4_t a, uint16x4_t b,
  uint16x8_t c, const int lane)
{
  return (a + (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: uint32x2
vmla_laneq_u32 (uint32x2_t a, uint32x2_t b,
  uint32x4_t c, const int lane)
{
  return (a + (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: float32x4
vmlaq_lane_f32 (float32x4_t a, float32x4_t b,
  float32x2_t c, const int lane)
{
  return (a + (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: int16x8
vmlaq_lane_s16 (int16x8_t a, int16x8_t b,
  int16x4_t c, const int lane)
{
  return (a + (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: int32x4
vmlaq_lane_s32 (int32x4_t a, int32x4_t b,
  int32x2_t c, const int lane)
{
  return (a + (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: uint16x8
vmlaq_lane_u16 (uint16x8_t a, uint16x8_t b,
  uint16x4_t c, const int lane)
{
  return (a + (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: uint32x4
vmlaq_lane_u32 (uint32x4_t a, uint32x4_t b,
  uint32x2_t c, const int lane)
{
  return (a + (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: float32x4
vmlaq_laneq_f32 (float32x4_t a, float32x4_t b,
   float32x4_t c, const int lane)
{
  return (a + (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: int16x8
vmlaq_laneq_s16 (int16x8_t a, int16x8_t b,
  int16x8_t c, const int lane)
{
  return (a + (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: int32x4
vmlaq_laneq_s32 (int32x4_t a, int32x4_t b,
  int32x4_t c, const int lane)
{
  return (a + (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: uint16x8
vmlaq_laneq_u16 (uint16x8_t a, uint16x8_t b,
  uint16x8_t c, const int lane)
{
  return (a + (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: uint32x4
vmlaq_laneq_u32 (uint32x4_t a, uint32x4_t b,
  uint32x4_t c, const int lane)
{
  return (a + (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: float32x2
vmls_f32 (float32x2_t a, float32x2_t b, float32x2_t c)
{
  return a - b * c;
}

DEF: float64x1
vmls_f64 (float64x1_t a, float64x1_t b, float64x1_t c)
{
  return a - b * c;
}

DEF: float32x4
vmlsq_f32 (float32x4_t a, float32x4_t b, float32x4_t c)
{
  return a - b * c;
}

DEF: float64x2
vmlsq_f64 (float64x2_t a, float64x2_t b, float64x2_t c)
{
  return a - b * c;
}

DEF: float32x2
vmls_lane_f32 (float32x2_t a, float32x2_t b,
        float32x2_t c, const int lane)
{
  return (a - (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: int16x4
vmls_lane_s16 (int16x4_t a, int16x4_t b,
  int16x4_t c, const int lane)
{
  return (a - (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: int32x2
vmls_lane_s32 (int32x2_t a, int32x2_t b,
  int32x2_t c, const int lane)
{
  return (a - (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: uint16x4
vmls_lane_u16 (uint16x4_t a, uint16x4_t b,
  uint16x4_t c, const int lane)
{
  return (a - (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: uint32x2
vmls_lane_u32 (uint32x2_t a, uint32x2_t b,
        uint32x2_t c, const int lane)
{
  return (a - (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: float32x2
vmls_laneq_f32 (float32x2_t a, float32x2_t b,
        float32x4_t c, const int lane)
{
  return (a - (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: int16x4
vmls_laneq_s16 (int16x4_t a, int16x4_t b,
  int16x8_t c, const int lane)
{
  return (a - (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: int32x2
vmls_laneq_s32 (int32x2_t a, int32x2_t b,
  int32x4_t c, const int lane)
{
  return (a - (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: uint16x4
vmls_laneq_u16 (uint16x4_t a, uint16x4_t b,
  uint16x8_t c, const int lane)
{
  return (a - (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: uint32x2
vmls_laneq_u32 (uint32x2_t a, uint32x2_t b,
  uint32x4_t c, const int lane)
{
  return (a - (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: float32x4
vmlsq_lane_f32 (float32x4_t a, float32x4_t b,
  float32x2_t c, const int lane)
{
  return (a - (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: int16x8
vmlsq_lane_s16 (int16x8_t a, int16x8_t b,
  int16x4_t c, const int lane)
{
  return (a - (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: int32x4
vmlsq_lane_s32 (int32x4_t a, int32x4_t b,
  int32x2_t c, const int lane)
{
  return (a - (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: uint16x8
vmlsq_lane_u16 (uint16x8_t a, uint16x8_t b,
  uint16x4_t c, const int lane)
{
  return (a - (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: uint32x4
vmlsq_lane_u32 (uint32x4_t a, uint32x4_t b,
  uint32x2_t c, const int lane)
{
  return (a - (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: float32x4
vmlsq_laneq_f32 (float32x4_t a, float32x4_t b,
  float32x4_t c, const int lane)
{
  return (a - (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: int16x8
vmlsq_laneq_s16 (int16x8_t a, int16x8_t b,
  int16x8_t c, const int lane)
{
  return (a - (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: int32x4
vmlsq_laneq_s32 (int32x4_t a, int32x4_t b,
  int32x4_t c, const int lane)
{
  return (a - (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}
DEF: uint16x8
vmlsq_laneq_u16 (uint16x8_t a, uint16x8_t b,
  uint16x8_t c, const int lane)
{
  return (a - (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: uint32x4
vmlsq_laneq_u32 (uint32x4_t a, uint32x4_t b,
  uint32x4_t c, const int lane)
{
  return (a - (b * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(c), sizeof(c[0]), lane); c[lane]; })));
}

DEF: float32x2
vmov_n_f32 (float32_t a)
{
  return vdup_n_f32 (a);
}

DEF: float64x1
vmov_n_f64 (float64_t a)
{
  return (float64x1_t) {a};
}

DEF: poly8x8
vmov_n_p8 (poly8_t a)
{
  return vdup_n_p8 (a);
}

DEF: poly16x4
vmov_n_p16 (poly16_t a)
{
  return vdup_n_p16 (a);
}

DEF: int8x8
vmov_n_s8 (int8_t a)
{
  return vdup_n_s8 (a);
}

DEF: int16x4
vmov_n_s16 (int16_t a)
{
  return vdup_n_s16 (a);
}

DEF: int32x2
vmov_n_s32 (int32_t a)
{
  return vdup_n_s32 (a);
}

DEF: int64x1
vmov_n_s64 (int64_t a)
{
  return (int64x1_t) {a};
}

DEF: uint8x8
vmov_n_u8 (uint8_t a)
{
  return vdup_n_u8 (a);
}

DEF: uint16x4
vmov_n_u16 (uint16_t a)
{
    return vdup_n_u16 (a);
}

DEF: uint32x2
vmov_n_u32 (uint32_t a)
{
   return vdup_n_u32 (a);
}

DEF: uint64x1
vmov_n_u64 (uint64_t a)
{
  return (uint64x1_t) {a};
}

DEF: float32x4
vmovq_n_f32 (float32_t a)
{
  return vdupq_n_f32 (a);
}

DEF: float64x2
vmovq_n_f64 (float64_t a)
{
  return vdupq_n_f64 (a);
}

DEF: poly8x16
vmovq_n_p8 (poly8_t a)
{
  return vdupq_n_p8 (a);
}

DEF: poly16x8
vmovq_n_p16 (poly16_t a)
{
  return vdupq_n_p16 (a);
}

DEF: int8x16
vmovq_n_s8 (int8_t a)
{
  return vdupq_n_s8 (a);
}

DEF: int16x8
vmovq_n_s16 (int16_t a)
{
  return vdupq_n_s16 (a);
}

DEF: int32x4
vmovq_n_s32 (int32_t a)
{
  return vdupq_n_s32 (a);
}

DEF: int64x2
vmovq_n_s64 (int64_t a)
{
  return vdupq_n_s64 (a);
}

DEF: uint8x16
vmovq_n_u8 (uint8_t a)
{
  return vdupq_n_u8 (a);
}

DEF: uint16x8
vmovq_n_u16 (uint16_t a)
{
  return vdupq_n_u16 (a);
}

DEF: uint32x4
vmovq_n_u32 (uint32_t a)
{
  return vdupq_n_u32 (a);
}

DEF: uint64x2
vmovq_n_u64 (uint64_t a)
{
  return vdupq_n_u64 (a);
}

DEF: float32x2
vmul_lane_f32 (float32x2_t a, float32x2_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: float64x1
vmul_lane_f64 (float64x1_t a, float64x1_t b, const int lane)
{
  return a * b;
}

DEF: int16x4
vmul_lane_s16 (int16x4_t a, int16x4_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: int32x2
vmul_lane_s32 (int32x2_t a, int32x2_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: uint16x4
vmul_lane_u16 (uint16x4_t a, uint16x4_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: uint32x2
vmul_lane_u32 (uint32x2_t a, uint32x2_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: float64
vmuld_lane_f64 (float64_t a, float64x1_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: float64
vmuld_laneq_f64 (float64_t a, float64x2_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: float32
vmuls_lane_f32 (float32_t a, float32x2_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: float32
vmuls_laneq_f32 (float32_t a, float32x4_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: float32x2
vmul_laneq_f32 (float32x2_t a, float32x4_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: float64x1
vmul_laneq_f64 (float64x1_t a, float64x2_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: int16x4
vmul_laneq_s16 (int16x4_t a, int16x8_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: int32x2
vmul_laneq_s32 (int32x2_t a, int32x4_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: uint16x4
vmul_laneq_u16 (uint16x4_t a, uint16x8_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: uint32x2
vmul_laneq_u32 (uint32x2_t a, uint32x4_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: float64x1
vmul_n_f64 (float64x1_t a, float64_t b)
{
  return (float64x1_t) { vget_lane_f64 (a, 0) * b };
}

DEF: float32x4
vmulq_lane_f32 (float32x4_t a, float32x2_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: float64x2
vmulq_lane_f64 (float64x2_t a, float64x1_t b, const int lane)
{
  builtin_aarch64_im_lane_boundsi (sizeof(a), sizeof(a[0]), lane);
  return a * b[0];
}

DEF: int16x8
vmulq_lane_s16 (int16x8_t a, int16x4_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: int32x4
vmulq_lane_s32 (int32x4_t a, int32x2_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: uint16x8
vmulq_lane_u16 (uint16x8_t a, uint16x4_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: uint32x4
vmulq_lane_u32 (uint32x4_t a, uint32x2_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: float32x4
vmulq_laneq_f32 (float32x4_t a, float32x4_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: float64x2
vmulq_laneq_f64 (float64x2_t a, float64x2_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: int16x8
vmulq_laneq_s16 (int16x8_t a, int16x8_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: int32x4
vmulq_laneq_s32 (int32x4_t a, int32x4_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: uint16x8
vmulq_laneq_u16 (uint16x8_t a, uint16x8_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: uint32x4
vmulq_laneq_u32 (uint32x4_t a, uint32x4_t b, const int lane)
{
  return a * extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: float32x2
vneg_f32 (float32x2_t a)
{
  return -a;
}

DEF: float64x1
vneg_f64 (float64x1_t a)
{
  return -a;
}

DEF: int8x8
vneg_s8 (int8x8_t a)
{
  return -a;
}

DEF: int16x4
vneg_s16 (int16x4_t a)
{
  return -a;
}

DEF: int32x2
vneg_s32 (int32x2_t a)
{
  return -a;
}

DEF: int64x1
vneg_s64 (int64x1_t a)
{
  return -a;
}

DEF: float32x4
vnegq_f32 (float32x4_t a)
{
  return -a;
}

DEF: float64x2
vnegq_f64 (float64x2_t a)
{
  return -a;
}

DEF: int8x16
vnegq_s8 (int8x16_t a)
{
  return -a;
}

DEF: int16x8
vnegq_s16 (int16x8_t a)
{
  return -a;
}

DEF: int32x4
vnegq_s32 (int32x4_t a)
{
  return -a;
}

DEF: int64x2
vnegq_s64 (int64x2_t a)
{
  return -a;
}

DEF: int8x8
vpadd_s8 (int8x8_t a, int8x8_t b)
{
  return builtin_aarch64_addpv8qi (a, b);
}

DEF: int16x4
vpadd_s16 (int16x4_t a, int16x4_t b)
{
  return builtin_aarch64_addpv4hi (a, b);
}

DEF: int32x2
vpadd_s32 (int32x2_t a, int32x2_t b)
{
  return builtin_aarch64_addpv2si (a, b);
}

DEF: uint8x8
vpadd_u8 (uint8x8_t a, uint8x8_t b)
{
  return (uint8x8_t) builtin_aarch64_addpv8qi ((int8x8_t) a,
       (int8x8_t) b);
}

DEF: uint16x4
vpadd_u16 (uint16x4_t a, uint16x4_t b)
{
  return (uint16x4_t) builtin_aarch64_addpv4hi ((int16x4_t) a,
        (int16x4_t) b);
}

DEF: uint32x2
vpadd_u32 (uint32x2_t a, uint32x2_t b)
{
  return (uint32x2_t) builtin_aarch64_addpv2si ((int32x2_t) a,
        (int32x2_t) b);
}

DEF: float64
vpaddd_f64 (float64x2_t a)
{
  return builtin_aarch64_reduc_plus_scal_v2df (a);
}

DEF: int64
vpaddd_s64 (int64x2_t a)
{
  return builtin_aarch64_addpdi (a);
}

DEF: uint64
vpaddd_u64 (uint64x2_t a)
{
  return builtin_aarch64_addpdi ((int64x2_t) a);
}

DEF: int64x2
vqabsq_s64 (int64x2_t a)
{
  return (int64x2_t) builtin_aarch64_sqabsv2di (a);
}

DEF: int8
vqabsb_s8 (int8_t a)
{
  return (int8_t) builtin_aarch64_sqabsqi (a);
}

DEF: int16
vqabsh_s16 (int16_t a)
{
  return (int16_t) builtin_aarch64_sqabshi (a);
}

DEF: int32
vqabss_s32 (int32_t a)
{
  return (int32_t) builtin_aarch64_sqabssi (a);
}

DEF: int64
vqabsd_s64 (int64_t a)
{
  return builtin_aarch64_sqabsdi (a);
}

DEF: int8
vqaddb_s8 (int8_t a, int8_t b)
{
  return (int8_t) builtin_aarch64_sqaddqi (a, b);
}

DEF: int16
vqaddh_s16 (int16_t a, int16_t b)
{
  return (int16_t) builtin_aarch64_sqaddhi (a, b);
}

DEF: int32
vqadds_s32 (int32_t a, int32_t b)
{
  return (int32_t) builtin_aarch64_sqaddsi (a, b);
}

DEF: int64
vqaddd_s64 (int64_t a, int64_t b)
{
  return builtin_aarch64_sqadddi (a, b);
}

DEF: uint8
vqaddb_u8 (uint8_t a, uint8_t b)
{
  return (uint8_t) builtin_aarch64_uqaddqi_uuu (a, b);
}

DEF: uint16
vqaddh_u16 (uint16_t a, uint16_t b)
{
  return (uint16_t) builtin_aarch64_uqaddhi_uuu (a, b);
}

DEF: uint32
vqadds_u32 (uint32_t a, uint32_t b)
{
  return (uint32_t) builtin_aarch64_uqaddsi_uuu (a, b);
}

DEF: uint64
vqaddd_u64 (uint64_t a, uint64_t b)
{
  return builtin_aarch64_uqadddi_uuu (a, b);
}

DEF: int32x4
vqdmlal_s16 (int32x4_t a, int16x4_t b, int16x4_t c)
{
  return builtin_aarch64_sqdmlalv4hi (a, b, c);
}

DEF: int32x4
vqdmlal_high_s16 (int32x4_t a, int16x8_t b, int16x8_t c)
{
  return builtin_aarch64_sqdmlal2v8hi (a, b, c);
}

DEF: int32x4
vqdmlal_high_lane_s16 (int32x4_t a, int16x8_t b, int16x4_t c,
         int const d)
{
  return builtin_aarch64_sqdmlal2_lanev8hi (a, b, c, d);
}

DEF: int32x4
vqdmlal_high_laneq_s16 (int32x4_t a, int16x8_t b, int16x8_t c,
   int const d)
{
  return builtin_aarch64_sqdmlal2_laneqv8hi (a, b, c, d);
}

DEF: int32x4
vqdmlal_high_n_s16 (int32x4_t a, int16x8_t b, int16_t c)
{
  return builtin_aarch64_sqdmlal2_nv8hi (a, b, c);
}

DEF: int32x4
vqdmlal_lane_s16 (int32x4_t a, int16x4_t b, int16x4_t c, int const d)
{
  return builtin_aarch64_sqdmlal_lanev4hi (a, b, c, d);
}

DEF: int32x4
vqdmlal_laneq_s16 (int32x4_t a, int16x4_t b, int16x8_t c, int const d)
{
  return builtin_aarch64_sqdmlal_laneqv4hi (a, b, c, d);
}

DEF: int32x4
vqdmlal_n_s16 (int32x4_t a, int16x4_t b, int16_t c)
{
  return builtin_aarch64_sqdmlal_nv4hi (a, b, c);
}

DEF: int64x2
vqdmlal_s32 (int64x2_t a, int32x2_t b, int32x2_t c)
{
  return builtin_aarch64_sqdmlalv2si (a, b, c);
}

DEF: int64x2
vqdmlal_high_s32 (int64x2_t a, int32x4_t b, int32x4_t c)
{
  return builtin_aarch64_sqdmlal2v4si (a, b, c);
}

DEF: int64x2
vqdmlal_high_lane_s32 (int64x2_t a, int32x4_t b, int32x2_t c,
         int const d)
{
  return builtin_aarch64_sqdmlal2_lanev4si (a, b, c, d);
}

DEF: int64x2
vqdmlal_high_laneq_s32 (int64x2_t a, int32x4_t b, int32x4_t c,
   int const d)
{
  return builtin_aarch64_sqdmlal2_laneqv4si (a, b, c, d);
}

DEF: int64x2
vqdmlal_high_n_s32 (int64x2_t a, int32x4_t b, int32_t c)
{
  return builtin_aarch64_sqdmlal2_nv4si (a, b, c);
}

DEF: int64x2
vqdmlal_lane_s32 (int64x2_t a, int32x2_t b, int32x2_t c, int const d)
{
  return builtin_aarch64_sqdmlal_lanev2si (a, b, c, d);
}

DEF: int64x2
vqdmlal_laneq_s32 (int64x2_t a, int32x2_t b, int32x4_t c, int const d)
{
  return builtin_aarch64_sqdmlal_laneqv2si (a, b, c, d);
}

DEF: int64x2
vqdmlal_n_s32 (int64x2_t a, int32x2_t b, int32_t c)
{
  return builtin_aarch64_sqdmlal_nv2si (a, b, c);
}

DEF: int32
vqdmlalh_s16 (int32_t a, int16_t b, int16_t c)
{
  return builtin_aarch64_sqdmlalhi (a, b, c);
}

DEF: int32
vqdmlalh_lane_s16 (int32_t a, int16_t b, int16x4_t c, const int d)
{
  return builtin_aarch64_sqdmlal_lanehi (a, b, c, d);
}

DEF: int32
vqdmlalh_laneq_s16 (int32_t a, int16_t b, int16x8_t c, const int d)
{
  return builtin_aarch64_sqdmlal_laneqhi (a, b, c, d);
}

DEF: int64
vqdmlals_s32 (int64_t a, int32_t b, int32_t c)
{
  return builtin_aarch64_sqdmlalsi (a, b, c);
}

DEF: int64
vqdmlals_lane_s32 (int64_t a, int32_t b, int32x2_t c, const int d)
{
  return builtin_aarch64_sqdmlal_lanesi (a, b, c, d);
}

DEF: int64
vqdmlals_laneq_s32 (int64_t a, int32_t b, int32x4_t c, const int d)
{
  return builtin_aarch64_sqdmlal_laneqsi (a, b, c, d);
}

DEF: int32x4
vqdmlsl_s16 (int32x4_t a, int16x4_t b, int16x4_t c)
{
  return builtin_aarch64_sqdmlslv4hi (a, b, c);
}

DEF: int32x4
vqdmlsl_high_s16 (int32x4_t a, int16x8_t b, int16x8_t c)
{
  return builtin_aarch64_sqdmlsl2v8hi (a, b, c);
}

DEF: int32x4
vqdmlsl_high_lane_s16 (int32x4_t a, int16x8_t b, int16x4_t c,
         int const d)
{
  return builtin_aarch64_sqdmlsl2_lanev8hi (a, b, c, d);
}

DEF: int32x4
vqdmlsl_high_laneq_s16 (int32x4_t a, int16x8_t b, int16x8_t c,
   int const d)
{
  return builtin_aarch64_sqdmlsl2_laneqv8hi (a, b, c, d);
}

DEF: int32x4
vqdmlsl_high_n_s16 (int32x4_t a, int16x8_t b, int16_t c)
{
  return builtin_aarch64_sqdmlsl2_nv8hi (a, b, c);
}

DEF: int32x4
vqdmlsl_lane_s16 (int32x4_t a, int16x4_t b, int16x4_t c, int const d)
{
  return builtin_aarch64_sqdmlsl_lanev4hi (a, b, c, d);
}

DEF: int32x4
vqdmlsl_laneq_s16 (int32x4_t a, int16x4_t b, int16x8_t c, int const d)
{
  return builtin_aarch64_sqdmlsl_laneqv4hi (a, b, c, d);
}

DEF: int32x4
vqdmlsl_n_s16 (int32x4_t a, int16x4_t b, int16_t c)
{
  return builtin_aarch64_sqdmlsl_nv4hi (a, b, c);
}

DEF: int64x2
vqdmlsl_s32 (int64x2_t a, int32x2_t b, int32x2_t c)
{
  return builtin_aarch64_sqdmlslv2si (a, b, c);
}

DEF: int64x2
vqdmlsl_high_s32 (int64x2_t a, int32x4_t b, int32x4_t c)
{
  return builtin_aarch64_sqdmlsl2v4si (a, b, c);
}

DEF: int64x2
vqdmlsl_high_lane_s32 (int64x2_t a, int32x4_t b, int32x2_t c,
         int const d)
{
  return builtin_aarch64_sqdmlsl2_lanev4si (a, b, c, d);
}

DEF: int64x2
vqdmlsl_high_laneq_s32 (int64x2_t a, int32x4_t b, int32x4_t c,
   int const d)
{
  return builtin_aarch64_sqdmlsl2_laneqv4si (a, b, c, d);
}

DEF: int64x2
vqdmlsl_high_n_s32 (int64x2_t a, int32x4_t b, int32_t c)
{
  return builtin_aarch64_sqdmlsl2_nv4si (a, b, c);
}

DEF: int64x2
vqdmlsl_lane_s32 (int64x2_t a, int32x2_t b, int32x2_t c, int const d)
{
  return builtin_aarch64_sqdmlsl_lanev2si (a, b, c, d);
}

DEF: int64x2
vqdmlsl_laneq_s32 (int64x2_t a, int32x2_t b, int32x4_t c, int const d)
{
  return builtin_aarch64_sqdmlsl_laneqv2si (a, b, c, d);
}

DEF: int64x2
vqdmlsl_n_s32 (int64x2_t a, int32x2_t b, int32_t c)
{
  return builtin_aarch64_sqdmlsl_nv2si (a, b, c);
}

DEF: int32
vqdmlslh_s16 (int32_t a, int16_t b, int16_t c)
{
  return builtin_aarch64_sqdmlslhi (a, b, c);
}

DEF: int32
vqdmlslh_lane_s16 (int32_t a, int16_t b, int16x4_t c, const int d)
{
  return builtin_aarch64_sqdmlsl_lanehi (a, b, c, d);
}

DEF: int32
vqdmlslh_laneq_s16 (int32_t a, int16_t b, int16x8_t c, const int d)
{
  return builtin_aarch64_sqdmlsl_laneqhi (a, b, c, d);
}

DEF: int64
vqdmlsls_s32 (int64_t a, int32_t b, int32_t c)
{
  return builtin_aarch64_sqdmlslsi (a, b, c);
}

DEF: int64
vqdmlsls_lane_s32 (int64_t a, int32_t b, int32x2_t c, const int d)
{
  return builtin_aarch64_sqdmlsl_lanesi (a, b, c, d);
}

DEF: int64
vqdmlsls_laneq_s32 (int64_t a, int32_t b, int32x4_t c, const int d)
{
  return builtin_aarch64_sqdmlsl_laneqsi (a, b, c, d);
}

DEF: int16x4
vqdmulh_lane_s16 (int16x4_t a, int16x4_t b, const int c)
{
  return builtin_aarch64_sqdmulh_lanev4hi (a, b, c);
}

DEF: int32x2
vqdmulh_lane_s32 (int32x2_t a, int32x2_t b, const int c)
{
  return builtin_aarch64_sqdmulh_lanev2si (a, b, c);
}

DEF: int16x8
vqdmulhq_lane_s16 (int16x8_t a, int16x4_t b, const int c)
{
  return builtin_aarch64_sqdmulh_lanev8hi (a, b, c);
}

DEF: int32x4
vqdmulhq_lane_s32 (int32x4_t a, int32x2_t b, const int c)
{
  return builtin_aarch64_sqdmulh_lanev4si (a, b, c);
}

DEF: int16
vqdmulhh_s16 (int16_t a, int16_t b)
{
  return (int16_t) builtin_aarch64_sqdmulhhi (a, b);
}

DEF: int16
vqdmulhh_lane_s16 (int16_t a, int16x4_t b, const int c)
{
  return builtin_aarch64_sqdmulh_lanehi (a, b, c);
}

DEF: int16
vqdmulhh_laneq_s16 (int16_t a, int16x8_t b, const int c)
{
  return builtin_aarch64_sqdmulh_laneqhi (a, b, c);
}

DEF: int32
vqdmulhs_s32 (int32_t a, int32_t b)
{
  return (int32_t) builtin_aarch64_sqdmulhsi (a, b);
}

DEF: int32
vqdmulhs_lane_s32 (int32_t a, int32x2_t b, const int c)
{
  return builtin_aarch64_sqdmulh_lanesi (a, b, c);
}

DEF: int32
vqdmulhs_laneq_s32 (int32_t a, int32x4_t b, const int c)
{
  return builtin_aarch64_sqdmulh_laneqsi (a, b, c);
}

DEF: int32x4
vqdmull_s16 (int16x4_t a, int16x4_t b)
{
  return builtin_aarch64_sqdmullv4hi (a, b);
}

DEF: int32x4
vqdmull_high_s16 (int16x8_t a, int16x8_t b)
{
  return builtin_aarch64_sqdmull2v8hi (a, b);
}

DEF: int32x4
vqdmull_high_lane_s16 (int16x8_t a, int16x4_t b, int const c)
{
  return builtin_aarch64_sqdmull2_lanev8hi (a, b,c);
}

DEF: int32x4
vqdmull_high_laneq_s16 (int16x8_t a, int16x8_t b, int const c)
{
  return builtin_aarch64_sqdmull2_laneqv8hi (a, b,c);
}

DEF: int32x4
vqdmull_high_n_s16 (int16x8_t a, int16_t b)
{
  return builtin_aarch64_sqdmull2_nv8hi (a, b);
}

DEF: int32x4
vqdmull_lane_s16 (int16x4_t a, int16x4_t b, int const c)
{
  return builtin_aarch64_sqdmull_lanev4hi (a, b, c);
}

DEF: int32x4
vqdmull_laneq_s16 (int16x4_t a, int16x8_t b, int const c)
{
  return builtin_aarch64_sqdmull_laneqv4hi (a, b, c);
}

DEF: int32x4
vqdmull_n_s16 (int16x4_t a, int16_t b)
{
  return builtin_aarch64_sqdmull_nv4hi (a, b);
}

DEF: int64x2
vqdmull_s32 (int32x2_t a, int32x2_t b)
{
  return builtin_aarch64_sqdmullv2si (a, b);
}

DEF: int64x2
vqdmull_high_s32 (int32x4_t a, int32x4_t b)
{
  return builtin_aarch64_sqdmull2v4si (a, b);
}

DEF: int64x2
vqdmull_high_lane_s32 (int32x4_t a, int32x2_t b, int const c)
{
  return builtin_aarch64_sqdmull2_lanev4si (a, b, c);
}

DEF: int64x2
vqdmull_high_laneq_s32 (int32x4_t a, int32x4_t b, int const c)
{
  return builtin_aarch64_sqdmull2_laneqv4si (a, b, c);
}

DEF: int64x2
vqdmull_high_n_s32 (int32x4_t a, int32_t b)
{
  return builtin_aarch64_sqdmull2_nv4si (a, b);
}

DEF: int64x2
vqdmull_lane_s32 (int32x2_t a, int32x2_t b, int const c)
{
  return builtin_aarch64_sqdmull_lanev2si (a, b, c);
}

DEF: int64x2
vqdmull_laneq_s32 (int32x2_t a, int32x4_t b, int const c)
{
  return builtin_aarch64_sqdmull_laneqv2si (a, b, c);
}

DEF: int64x2
vqdmull_n_s32 (int32x2_t a, int32_t b)
{
  return builtin_aarch64_sqdmull_nv2si (a, b);
}

DEF: int32
vqdmullh_s16 (int16_t a, int16_t b)
{
  return (int32_t) builtin_aarch64_sqdmullhi (a, b);
}

DEF: int32
vqdmullh_lane_s16 (int16_t a, int16x4_t b, const int c)
{
  return builtin_aarch64_sqdmull_lanehi (a, b, c);
}

DEF: int32
vqdmullh_laneq_s16 (int16_t a, int16x8_t b, const int c)
{
  return builtin_aarch64_sqdmull_laneqhi (a, b, c);
}

DEF: int64
vqdmulls_s32 (int32_t a, int32_t b)
{
  return builtin_aarch64_sqdmullsi (a, b);
}

DEF: int64
vqdmulls_lane_s32 (int32_t a, int32x2_t b, const int c)
{
  return builtin_aarch64_sqdmull_lanesi (a, b, c);
}

DEF: int64
vqdmulls_laneq_s32 (int32_t a, int32x4_t b, const int c)
{
  return builtin_aarch64_sqdmull_laneqsi (a, b, c);
}

DEF: int8x8
vqmovn_s16 (int16x8_t a)
{
  return (int8x8_t) builtin_aarch64_sqmovnv8hi (a);
}

DEF: int16x4
vqmovn_s32 (int32x4_t a)
{
  return (int16x4_t) builtin_aarch64_sqmovnv4si (a);
}

DEF: int32x2
vqmovn_s64 (int64x2_t a)
{
  return (int32x2_t) builtin_aarch64_sqmovnv2di (a);
}

DEF: uint8x8
vqmovn_u16 (uint16x8_t a)
{
  return (uint8x8_t) builtin_aarch64_uqmovnv8hi ((int16x8_t) a);
}

DEF: uint16x4
vqmovn_u32 (uint32x4_t a)
{
  return (uint16x4_t) builtin_aarch64_uqmovnv4si ((int32x4_t) a);
}

DEF: uint32x2
vqmovn_u64 (uint64x2_t a)
{
  return (uint32x2_t) builtin_aarch64_uqmovnv2di ((int64x2_t) a);
}

DEF: int8
vqmovnh_s16 (int16_t a)
{
  return (int8_t) builtin_aarch64_sqmovnhi (a);
}

DEF: int16
vqmovns_s32 (int32_t a)
{
  return (int16_t) builtin_aarch64_sqmovnsi (a);
}

DEF: int32
vqmovnd_s64 (int64_t a)
{
  return (int32_t) builtin_aarch64_sqmovndi (a);
}

DEF: uint8
vqmovnh_u16 (uint16_t a)
{
  return (uint8_t) builtin_aarch64_uqmovnhi (a);
}

DEF: uint16
vqmovns_u32 (uint32_t a)
{
  return (uint16_t) builtin_aarch64_uqmovnsi (a);
}

DEF: uint32
vqmovnd_u64 (uint64_t a)
{
  return (uint32_t) builtin_aarch64_uqmovndi (a);
}

DEF: uint8x8
vqmovun_s16 (int16x8_t a)
{
  return (uint8x8_t) builtin_aarch64_sqmovunv8hi (a);
}

DEF: uint16x4
vqmovun_s32 (int32x4_t a)
{
  return (uint16x4_t) builtin_aarch64_sqmovunv4si (a);
}

DEF: uint32x2
vqmovun_s64 (int64x2_t a)
{
  return (uint32x2_t) builtin_aarch64_sqmovunv2di (a);
}

DEF: int8
vqmovunh_s16 (int16_t a)
{
  return (int8_t) builtin_aarch64_sqmovunhi (a);
}

DEF: int16
vqmovuns_s32 (int32_t a)
{
  return (int16_t) builtin_aarch64_sqmovunsi (a);
}

DEF: int32
vqmovund_s64 (int64_t a)
{
  return (int32_t) builtin_aarch64_sqmovundi (a);
}

DEF: int64x2
vqnegq_s64 (int64x2_t a)
{
  return (int64x2_t) builtin_aarch64_sqnegv2di (a);
}

DEF: int8
vqnegb_s8 (int8_t a)
{
  return (int8_t) builtin_aarch64_sqnegqi (a);
}

DEF: int16
vqnegh_s16 (int16_t a)
{
  return (int16_t) builtin_aarch64_sqneghi (a);
}

DEF: int32
vqnegs_s32 (int32_t a)
{
  return (int32_t) builtin_aarch64_sqnegsi (a);
}

DEF: int64
vqnegd_s64 (int64_t a)
{
  return builtin_aarch64_sqnegdi (a);
}

DEF: int16x4
vqrdmulh_lane_s16 (int16x4_t a, int16x4_t b, const int c)
{
  return builtin_aarch64_sqrdmulh_lanev4hi (a, b, c);
}

DEF: int32x2
vqrdmulh_lane_s32 (int32x2_t a, int32x2_t b, const int c)
{
  return builtin_aarch64_sqrdmulh_lanev2si (a, b, c);
}

DEF: int16x8
vqrdmulhq_lane_s16 (int16x8_t a, int16x4_t b, const int c)
{
  return builtin_aarch64_sqrdmulh_lanev8hi (a, b, c);
}

DEF: int32x4
vqrdmulhq_lane_s32 (int32x4_t a, int32x2_t b, const int c)
{
  return builtin_aarch64_sqrdmulh_lanev4si (a, b, c);
}

DEF: int16
vqrdmulhh_s16 (int16_t a, int16_t b)
{
  return (int16_t) builtin_aarch64_sqrdmulhhi (a, b);
}

DEF: int16
vqrdmulhh_lane_s16 (int16_t a, int16x4_t b, const int c)
{
  return builtin_aarch64_sqrdmulh_lanehi (a, b, c);
}

DEF: int16
vqrdmulhh_laneq_s16 (int16_t a, int16x8_t b, const int c)
{
  return builtin_aarch64_sqrdmulh_laneqhi (a, b, c);
}

DEF: int32
vqrdmulhs_s32 (int32_t a, int32_t b)
{
  return (int32_t) builtin_aarch64_sqrdmulhsi (a, b);
}

DEF: int32
vqrdmulhs_lane_s32 (int32_t a, int32x2_t b, const int c)
{
  return builtin_aarch64_sqrdmulh_lanesi (a, b, c);
}

DEF: int32
vqrdmulhs_laneq_s32 (int32_t a, int32x4_t b, const int c)
{
  return builtin_aarch64_sqrdmulh_laneqsi (a, b, c);
}

DEF: int8x8
vqrshl_s8 (int8x8_t a, int8x8_t b)
{
  return builtin_aarch64_sqrshlv8qi (a, b);
}

DEF: int16x4
vqrshl_s16 (int16x4_t a, int16x4_t b)
{
  return builtin_aarch64_sqrshlv4hi (a, b);
}

DEF: int32x2
vqrshl_s32 (int32x2_t a, int32x2_t b)
{
  return builtin_aarch64_sqrshlv2si (a, b);
}

DEF: int64x1
vqrshl_s64 (int64x1_t a, int64x1_t b)
{
  return (int64x1_t) {builtin_aarch64_sqrshldi (a[0], b[0])};
}

DEF: uint8x8
vqrshl_u8 (uint8x8_t a, int8x8_t b)
{
  return builtin_aarch64_uqrshlv8qi_uus ( a, b);
}

DEF: uint16x4
vqrshl_u16 (uint16x4_t a, int16x4_t b)
{
  return builtin_aarch64_uqrshlv4hi_uus ( a, b);
}

DEF: uint32x2
vqrshl_u32 (uint32x2_t a, int32x2_t b)
{
  return builtin_aarch64_uqrshlv2si_uus ( a, b);
}

DEF: uint64x1
vqrshl_u64 (uint64x1_t a, int64x1_t b)
{
  return (uint64x1_t) {builtin_aarch64_uqrshldi_uus (a[0], b[0])};
}

DEF: int8x16
vqrshlq_s8 (int8x16_t a, int8x16_t b)
{
  return builtin_aarch64_sqrshlv16qi (a, b);
}

DEF: int16x8
vqrshlq_s16 (int16x8_t a, int16x8_t b)
{
  return builtin_aarch64_sqrshlv8hi (a, b);
}

DEF: int32x4
vqrshlq_s32 (int32x4_t a, int32x4_t b)
{
  return builtin_aarch64_sqrshlv4si (a, b);
}

DEF: int64x2
vqrshlq_s64 (int64x2_t a, int64x2_t b)
{
  return builtin_aarch64_sqrshlv2di (a, b);
}

DEF: uint8x16
vqrshlq_u8 (uint8x16_t a, int8x16_t b)
{
  return builtin_aarch64_uqrshlv16qi_uus ( a, b);
}

DEF: uint16x8
vqrshlq_u16 (uint16x8_t a, int16x8_t b)
{
  return builtin_aarch64_uqrshlv8hi_uus ( a, b);
}

DEF: uint32x4
vqrshlq_u32 (uint32x4_t a, int32x4_t b)
{
  return builtin_aarch64_uqrshlv4si_uus ( a, b);
}

DEF: uint64x2
vqrshlq_u64 (uint64x2_t a, int64x2_t b)
{
  return builtin_aarch64_uqrshlv2di_uus ( a, b);
}

DEF: int8
vqrshlb_s8 (int8_t a, int8_t b)
{
  return builtin_aarch64_sqrshlqi (a, b);
}

DEF: int16
vqrshlh_s16 (int16_t a, int16_t b)
{
  return builtin_aarch64_sqrshlhi (a, b);
}

DEF: int32
vqrshls_s32 (int32_t a, int32_t b)
{
  return builtin_aarch64_sqrshlsi (a, b);
}

DEF: int64
vqrshld_s64 (int64_t a, int64_t b)
{
  return builtin_aarch64_sqrshldi (a, b);
}

DEF: uint8
vqrshlb_u8 (uint8_t a, uint8_t b)
{
  return builtin_aarch64_uqrshlqi_uus (a, b);
}

DEF: uint16
vqrshlh_u16 (uint16_t a, uint16_t b)
{
  return builtin_aarch64_uqrshlhi_uus (a, b);
}

DEF: uint32
vqrshls_u32 (uint32_t a, uint32_t b)
{
  return builtin_aarch64_uqrshlsi_uus (a, b);
}

DEF: uint64
vqrshld_u64 (uint64_t a, uint64_t b)
{
  return builtin_aarch64_uqrshldi_uus (a, b);
}

DEF: int8x8
vqrshrn_n_s16 (int16x8_t a, const int b)
{
  return (int8x8_t) builtin_aarch64_sqrshrn_nv8hi (a, b);
}

DEF: int16x4
vqrshrn_n_s32 (int32x4_t a, const int b)
{
  return (int16x4_t) builtin_aarch64_sqrshrn_nv4si (a, b);
}

DEF: int32x2
vqrshrn_n_s64 (int64x2_t a, const int b)
{
  return (int32x2_t) builtin_aarch64_sqrshrn_nv2di (a, b);
}

DEF: uint8x8
vqrshrn_n_u16 (uint16x8_t a, const int b)
{
  return builtin_aarch64_uqrshrn_nv8hi_uus ( a, b);
}

DEF: uint16x4
vqrshrn_n_u32 (uint32x4_t a, const int b)
{
  return builtin_aarch64_uqrshrn_nv4si_uus ( a, b);
}

DEF: uint32x2
vqrshrn_n_u64 (uint64x2_t a, const int b)
{
  return builtin_aarch64_uqrshrn_nv2di_uus ( a, b);
}

DEF: int8
vqrshrnh_n_s16 (int16_t a, const int b)
{
  return (int8_t) builtin_aarch64_sqrshrn_nhi (a, b);
}

DEF: int16
vqrshrns_n_s32 (int32_t a, const int b)
{
  return (int16_t) builtin_aarch64_sqrshrn_nsi (a, b);
}

DEF: int32
vqrshrnd_n_s64 (int64_t a, const int b)
{
  return (int32_t) builtin_aarch64_sqrshrn_ndi (a, b);
}

DEF: uint8
vqrshrnh_n_u16 (uint16_t a, const int b)
{
  return builtin_aarch64_uqrshrn_nhi_uus (a, b);
}

DEF: uint16
vqrshrns_n_u32 (uint32_t a, const int b)
{
  return builtin_aarch64_uqrshrn_nsi_uus (a, b);
}

DEF: uint32
vqrshrnd_n_u64 (uint64_t a, const int b)
{
  return builtin_aarch64_uqrshrn_ndi_uus (a, b);
}

DEF: uint8x8
vqrshrun_n_s16 (int16x8_t a, const int b)
{
  return (uint8x8_t) builtin_aarch64_sqrshrun_nv8hi (a, b);
}

DEF: uint16x4
vqrshrun_n_s32 (int32x4_t a, const int b)
{
  return (uint16x4_t) builtin_aarch64_sqrshrun_nv4si (a, b);
}

DEF: uint32x2
vqrshrun_n_s64 (int64x2_t a, const int b)
{
  return (uint32x2_t) builtin_aarch64_sqrshrun_nv2di (a, b);
}

DEF: int8
vqrshrunh_n_s16 (int16_t a, const int b)
{
  return (int8_t) builtin_aarch64_sqrshrun_nhi (a, b);
}

DEF: int16
vqrshruns_n_s32 (int32_t a, const int b)
{
  return (int16_t) builtin_aarch64_sqrshrun_nsi (a, b);
}

DEF: int32
vqrshrund_n_s64 (int64_t a, const int b)
{
  return (int32_t) builtin_aarch64_sqrshrun_ndi (a, b);
}

DEF: int8x8
vqshl_s8 (int8x8_t a, int8x8_t b)
{
  return builtin_aarch64_sqshlv8qi (a, b);
}

DEF: int16x4
vqshl_s16 (int16x4_t a, int16x4_t b)
{
  return builtin_aarch64_sqshlv4hi (a, b);
}

DEF: int32x2
vqshl_s32 (int32x2_t a, int32x2_t b)
{
  return builtin_aarch64_sqshlv2si (a, b);
}

DEF: int64x1
vqshl_s64 (int64x1_t a, int64x1_t b)
{
  return (int64x1_t) {builtin_aarch64_sqshldi (a[0], b[0])};
}

DEF: uint8x8
vqshl_u8 (uint8x8_t a, int8x8_t b)
{
  return builtin_aarch64_uqshlv8qi_uus ( a, b);
}

DEF: uint16x4
vqshl_u16 (uint16x4_t a, int16x4_t b)
{
  return builtin_aarch64_uqshlv4hi_uus ( a, b);
}

DEF: uint32x2
vqshl_u32 (uint32x2_t a, int32x2_t b)
{
  return builtin_aarch64_uqshlv2si_uus ( a, b);
}

DEF: uint64x1
vqshl_u64 (uint64x1_t a, int64x1_t b)
{
  return (uint64x1_t) {builtin_aarch64_uqshldi_uus (a[0], b[0])};
}

DEF: int8x16
vqshlq_s8 (int8x16_t a, int8x16_t b)
{
  return builtin_aarch64_sqshlv16qi (a, b);
}

DEF: int16x8
vqshlq_s16 (int16x8_t a, int16x8_t b)
{
  return builtin_aarch64_sqshlv8hi (a, b);
}

DEF: int32x4
vqshlq_s32 (int32x4_t a, int32x4_t b)
{
  return builtin_aarch64_sqshlv4si (a, b);
}

DEF: int64x2
vqshlq_s64 (int64x2_t a, int64x2_t b)
{
  return builtin_aarch64_sqshlv2di (a, b);
}

DEF: uint8x16
vqshlq_u8 (uint8x16_t a, int8x16_t b)
{
  return builtin_aarch64_uqshlv16qi_uus ( a, b);
}

DEF: uint16x8
vqshlq_u16 (uint16x8_t a, int16x8_t b)
{
  return builtin_aarch64_uqshlv8hi_uus ( a, b);
}

DEF: uint32x4
vqshlq_u32 (uint32x4_t a, int32x4_t b)
{
  return builtin_aarch64_uqshlv4si_uus ( a, b);
}

DEF: uint64x2
vqshlq_u64 (uint64x2_t a, int64x2_t b)
{
  return builtin_aarch64_uqshlv2di_uus ( a, b);
}

DEF: int8
vqshlb_s8 (int8_t a, int8_t b)
{
  return builtin_aarch64_sqshlqi (a, b);
}

DEF: int16
vqshlh_s16 (int16_t a, int16_t b)
{
  return builtin_aarch64_sqshlhi (a, b);
}

DEF: int32
vqshls_s32 (int32_t a, int32_t b)
{
  return builtin_aarch64_sqshlsi (a, b);
}

DEF: int64
vqshld_s64 (int64_t a, int64_t b)
{
  return builtin_aarch64_sqshldi (a, b);
}

DEF: uint8
vqshlb_u8 (uint8_t a, uint8_t b)
{
  return builtin_aarch64_uqshlqi_uus (a, b);
}

DEF: uint16
vqshlh_u16 (uint16_t a, uint16_t b)
{
  return builtin_aarch64_uqshlhi_uus (a, b);
}

DEF: uint32
vqshls_u32 (uint32_t a, uint32_t b)
{
  return builtin_aarch64_uqshlsi_uus (a, b);
}

DEF: uint64
vqshld_u64 (uint64_t a, uint64_t b)
{
  return builtin_aarch64_uqshldi_uus (a, b);
}

DEF: int8x8
vqshl_n_s8 (int8x8_t a, const int b)
{
  return (int8x8_t) builtin_aarch64_sqshl_nv8qi (a, b);
}

DEF: int16x4
vqshl_n_s16 (int16x4_t a, const int b)
{
  return (int16x4_t) builtin_aarch64_sqshl_nv4hi (a, b);
}

DEF: int32x2
vqshl_n_s32 (int32x2_t a, const int b)
{
  return (int32x2_t) builtin_aarch64_sqshl_nv2si (a, b);
}

DEF: int64x1
vqshl_n_s64 (int64x1_t a, const int b)
{
  return (int64x1_t) {builtin_aarch64_sqshl_ndi (a[0], b)};
}

DEF: uint8x8
vqshl_n_u8 (uint8x8_t a, const int b)
{
  return builtin_aarch64_uqshl_nv8qi_uus (a, b);
}

DEF: uint16x4
vqshl_n_u16 (uint16x4_t a, const int b)
{
  return builtin_aarch64_uqshl_nv4hi_uus (a, b);
}

DEF: uint32x2
vqshl_n_u32 (uint32x2_t a, const int b)
{
  return builtin_aarch64_uqshl_nv2si_uus (a, b);
}

DEF: uint64x1
vqshl_n_u64 (uint64x1_t a, const int b)
{
  return (uint64x1_t) {builtin_aarch64_uqshl_ndi_uus (a[0], b)};
}

DEF: int8x16
vqshlq_n_s8 (int8x16_t a, const int b)
{
  return (int8x16_t) builtin_aarch64_sqshl_nv16qi (a, b);
}

DEF: int16x8
vqshlq_n_s16 (int16x8_t a, const int b)
{
  return (int16x8_t) builtin_aarch64_sqshl_nv8hi (a, b);
}

DEF: int32x4
vqshlq_n_s32 (int32x4_t a, const int b)
{
  return (int32x4_t) builtin_aarch64_sqshl_nv4si (a, b);
}

DEF: int64x2
vqshlq_n_s64 (int64x2_t a, const int b)
{
  return (int64x2_t) builtin_aarch64_sqshl_nv2di (a, b);
}

DEF: uint8x16
vqshlq_n_u8 (uint8x16_t a, const int b)
{
  return builtin_aarch64_uqshl_nv16qi_uus (a, b);
}

DEF: uint16x8
vqshlq_n_u16 (uint16x8_t a, const int b)
{
  return builtin_aarch64_uqshl_nv8hi_uus (a, b);
}

DEF: uint32x4
vqshlq_n_u32 (uint32x4_t a, const int b)
{
  return builtin_aarch64_uqshl_nv4si_uus (a, b);
}

DEF: uint64x2
vqshlq_n_u64 (uint64x2_t a, const int b)
{
  return builtin_aarch64_uqshl_nv2di_uus (a, b);
}

DEF: int8
vqshlb_n_s8 (int8_t a, const int b)
{
  return (int8_t) builtin_aarch64_sqshl_nqi (a, b);
}

DEF: int16
vqshlh_n_s16 (int16_t a, const int b)
{
  return (int16_t) builtin_aarch64_sqshl_nhi (a, b);
}

DEF: int32
vqshls_n_s32 (int32_t a, const int b)
{
  return (int32_t) builtin_aarch64_sqshl_nsi (a, b);
}

DEF: int64
vqshld_n_s64 (int64_t a, const int b)
{
  return builtin_aarch64_sqshl_ndi (a, b);
}

DEF: uint8
vqshlb_n_u8 (uint8_t a, const int b)
{
  return builtin_aarch64_uqshl_nqi_uus (a, b);
}

DEF: uint16
vqshlh_n_u16 (uint16_t a, const int b)
{
  return builtin_aarch64_uqshl_nhi_uus (a, b);
}

DEF: uint32
vqshls_n_u32 (uint32_t a, const int b)
{
  return builtin_aarch64_uqshl_nsi_uus (a, b);
}

DEF: uint64
vqshld_n_u64 (uint64_t a, const int b)
{
  return builtin_aarch64_uqshl_ndi_uus (a, b);
}

DEF: uint8x8
vqshlu_n_s8 (int8x8_t a, const int b)
{
  return builtin_aarch64_sqshlu_nv8qi_uss (a, b);
}

DEF: uint16x4
vqshlu_n_s16 (int16x4_t a, const int b)
{
  return builtin_aarch64_sqshlu_nv4hi_uss (a, b);
}

DEF: uint32x2
vqshlu_n_s32 (int32x2_t a, const int b)
{
  return builtin_aarch64_sqshlu_nv2si_uss (a, b);
}

DEF: uint64x1
vqshlu_n_s64 (int64x1_t a, const int b)
{
  return (uint64x1_t) {builtin_aarch64_sqshlu_ndi_uss (a[0], b)};
}

DEF: uint8x16
vqshluq_n_s8 (int8x16_t a, const int b)
{
  return builtin_aarch64_sqshlu_nv16qi_uss (a, b);
}

DEF: uint16x8
vqshluq_n_s16 (int16x8_t a, const int b)
{
  return builtin_aarch64_sqshlu_nv8hi_uss (a, b);
}

DEF: uint32x4
vqshluq_n_s32 (int32x4_t a, const int b)
{
  return builtin_aarch64_sqshlu_nv4si_uss (a, b);
}

DEF: uint64x2
vqshluq_n_s64 (int64x2_t a, const int b)
{
  return builtin_aarch64_sqshlu_nv2di_uss (a, b);
}

DEF: int8
vqshlub_n_s8 (int8_t a, const int b)
{
  return (int8_t) builtin_aarch64_sqshlu_nqi_uss (a, b);
}

DEF: int16
vqshluh_n_s16 (int16_t a, const int b)
{
  return (int16_t) builtin_aarch64_sqshlu_nhi_uss (a, b);
}

DEF: int32
vqshlus_n_s32 (int32_t a, const int b)
{
  return (int32_t) builtin_aarch64_sqshlu_nsi_uss (a, b);
}

DEF: uint64
vqshlud_n_s64 (int64_t a, const int b)
{
  return builtin_aarch64_sqshlu_ndi_uss (a, b);
}

DEF: int8x8
vqshrn_n_s16 (int16x8_t a, const int b)
{
  return (int8x8_t) builtin_aarch64_sqshrn_nv8hi (a, b);
}

DEF: int16x4
vqshrn_n_s32 (int32x4_t a, const int b)
{
  return (int16x4_t) builtin_aarch64_sqshrn_nv4si (a, b);
}

DEF: int32x2
vqshrn_n_s64 (int64x2_t a, const int b)
{
  return (int32x2_t) builtin_aarch64_sqshrn_nv2di (a, b);
}

DEF: uint8x8
vqshrn_n_u16 (uint16x8_t a, const int b)
{
  return builtin_aarch64_uqshrn_nv8hi_uus ( a, b);
}

DEF: uint16x4
vqshrn_n_u32 (uint32x4_t a, const int b)
{
  return builtin_aarch64_uqshrn_nv4si_uus ( a, b);
}

DEF: uint32x2
vqshrn_n_u64 (uint64x2_t a, const int b)
{
  return builtin_aarch64_uqshrn_nv2di_uus ( a, b);
}

DEF: int8
vqshrnh_n_s16 (int16_t a, const int b)
{
  return (int8_t) builtin_aarch64_sqshrn_nhi (a, b);
}

DEF: int16
vqshrns_n_s32 (int32_t a, const int b)
{
  return (int16_t) builtin_aarch64_sqshrn_nsi (a, b);
}

DEF: int32
vqshrnd_n_s64 (int64_t a, const int b)
{
  return (int32_t) builtin_aarch64_sqshrn_ndi (a, b);
}

DEF: uint8
vqshrnh_n_u16 (uint16_t a, const int b)
{
  return builtin_aarch64_uqshrn_nhi_uus (a, b);
}

DEF: uint16
vqshrns_n_u32 (uint32_t a, const int b)
{
  return builtin_aarch64_uqshrn_nsi_uus (a, b);
}

DEF: uint32
vqshrnd_n_u64 (uint64_t a, const int b)
{
  return builtin_aarch64_uqshrn_ndi_uus (a, b);
}

DEF: uint8x8
vqshrun_n_s16 (int16x8_t a, const int b)
{
  return (uint8x8_t) builtin_aarch64_sqshrun_nv8hi (a, b);
}

DEF: uint16x4
vqshrun_n_s32 (int32x4_t a, const int b)
{
  return (uint16x4_t) builtin_aarch64_sqshrun_nv4si (a, b);
}

DEF: uint32x2
vqshrun_n_s64 (int64x2_t a, const int b)
{
  return (uint32x2_t) builtin_aarch64_sqshrun_nv2di (a, b);
}

DEF: int8
vqshrunh_n_s16 (int16_t a, const int b)
{
  return (int8_t) builtin_aarch64_sqshrun_nhi (a, b);
}

DEF: int16
vqshruns_n_s32 (int32_t a, const int b)
{
  return (int16_t) builtin_aarch64_sqshrun_nsi (a, b);
}

DEF: int32
vqshrund_n_s64 (int64_t a, const int b)
{
  return (int32_t) builtin_aarch64_sqshrun_ndi (a, b);
}

DEF: int8
vqsubb_s8 (int8_t a, int8_t b)
{
  return (int8_t) builtin_aarch64_sqsubqi (a, b);
}

DEF: int16
vqsubh_s16 (int16_t a, int16_t b)
{
  return (int16_t) builtin_aarch64_sqsubhi (a, b);
}

DEF: int32
vqsubs_s32 (int32_t a, int32_t b)
{
  return (int32_t) builtin_aarch64_sqsubsi (a, b);
}

DEF: int64
vqsubd_s64 (int64_t a, int64_t b)
{
  return builtin_aarch64_sqsubdi (a, b);
}

DEF: uint8
vqsubb_u8 (uint8_t a, uint8_t b)
{
  return (uint8_t) builtin_aarch64_uqsubqi_uuu (a, b);
}

DEF: uint16
vqsubh_u16 (uint16_t a, uint16_t b)
{
  return (uint16_t) builtin_aarch64_uqsubhi_uuu (a, b);
}

DEF: uint32
vqsubs_u32 (uint32_t a, uint32_t b)
{
  return (uint32_t) builtin_aarch64_uqsubsi_uuu (a, b);
}

DEF: uint64
vqsubd_u64 (uint64_t a, uint64_t b)
{
  return builtin_aarch64_uqsubdi_uuu (a, b);
}

DEF: poly8x8
vrbit_p8 (poly8x8_t a)
{
  return (poly8x8_t) builtin_aarch64_rbitv8qi ((int8x8_t) a);
}

DEF: int8x8
vrbit_s8 (int8x8_t a)
{
  return builtin_aarch64_rbitv8qi (a);
}

DEF: uint8x8
vrbit_u8 (uint8x8_t a)
{
  return (uint8x8_t) builtin_aarch64_rbitv8qi ((int8x8_t) a);
}

DEF: poly8x16
vrbitq_p8 (poly8x16_t a)
{
  return (poly8x16_t) builtin_aarch64_rbitv16qi ((int8x16_t)a);
}

DEF: int8x16
vrbitq_s8 (int8x16_t a)
{
  return builtin_aarch64_rbitv16qi (a);
}

DEF: uint8x16
vrbitq_u8 (uint8x16_t a)
{
  return (uint8x16_t) builtin_aarch64_rbitv16qi ((int8x16_t) a);
}

DEF: uint32x2
vrecpe_u32 (uint32x2_t a)
{
  return (uint32x2_t) builtin_aarch64_urecpev2si ((int32x2_t) a);
}

DEF: uint32x4
vrecpeq_u32 (uint32x4_t a)
{
  return (uint32x4_t) builtin_aarch64_urecpev4si ((int32x4_t) a);
}

DEF: float32
vrecpes_f32 (float32_t a)
{
  return builtin_aarch64_frecpesf (a);
}

DEF: float64
vrecped_f64 (float64_t a)
{
  return builtin_aarch64_frecpedf (a);
}

DEF: float32x2
vrecpe_f32 (float32x2_t a)
{
  return builtin_aarch64_frecpev2sf (a);
}

DEF: float32x4
vrecpeq_f32 (float32x4_t a)
{
  return builtin_aarch64_frecpev4sf (a);
}

DEF: float64x2
vrecpeq_f64 (float64x2_t a)
{
  return builtin_aarch64_frecpev2df (a);
}

DEF: float32
vrecpss_f32 (float32_t a, float32_t b)
{
  return builtin_aarch64_frecpssf (a, b);
}

DEF: float64
vrecpsd_f64 (float64_t a, float64_t b)
{
  return builtin_aarch64_frecpsdf (a, b);
}

DEF: float32x2
vrecps_f32 (float32x2_t a, float32x2_t b)
{
  return builtin_aarch64_frecpsv2sf (a, b);
}

DEF: float32x4
vrecpsq_f32 (float32x4_t a, float32x4_t b)
{
  return builtin_aarch64_frecpsv4sf (a, b);
}

DEF: float64x2
vrecpsq_f64 (float64x2_t a, float64x2_t b)
{
  return builtin_aarch64_frecpsv2df (a, b);
}

DEF: float32
vrecpxs_f32 (float32_t a)
{
  return builtin_aarch64_frecpxsf (a);
}

DEF: float64
vrecpxd_f64 (float64_t a)
{
  return builtin_aarch64_frecpxdf (a);
}


DEF: poly8x8
vrev16_p8 (poly8x8_t a)
{
  return builtin_shuffle (a, (uint8x8_t) { 1, 0, 3, 2, 5, 4, 7, 6 });
}

DEF: int8x8
vrev16_s8 (int8x8_t a)
{
  return builtin_shuffle (a, (uint8x8_t) { 1, 0, 3, 2, 5, 4, 7, 6 });
}

DEF: uint8x8
vrev16_u8 (uint8x8_t a)
{
  return builtin_shuffle (a, (uint8x8_t) { 1, 0, 3, 2, 5, 4, 7, 6 });
}

DEF: poly8x16
vrev16q_p8 (poly8x16_t a)
{
  return builtin_shuffle (a,
      (uint8x16_t) { 1, 0, 3, 2, 5, 4, 7, 6, 9, 8, 11, 10, 13, 12, 15, 14 });
}

DEF: int8x16
vrev16q_s8 (int8x16_t a)
{
  return builtin_shuffle (a,
      (uint8x16_t) { 1, 0, 3, 2, 5, 4, 7, 6, 9, 8, 11, 10, 13, 12, 15, 14 });
}

DEF: uint8x16
vrev16q_u8 (uint8x16_t a)
{
  return builtin_shuffle (a,
      (uint8x16_t) { 1, 0, 3, 2, 5, 4, 7, 6, 9, 8, 11, 10, 13, 12, 15, 14 });
}

DEF: poly8x8
vrev32_p8 (poly8x8_t a)
{
  return builtin_shuffle (a, (uint8x8_t) { 3, 2, 1, 0, 7, 6, 5, 4 });
}

DEF: poly16x4
vrev32_p16 (poly16x4_t a)
{
  return builtin_shuffle (a, (uint16x4_t) { 1, 0, 3, 2 });
}

DEF: int8x8
vrev32_s8 (int8x8_t a)
{
  return builtin_shuffle (a, (uint8x8_t) { 3, 2, 1, 0, 7, 6, 5, 4 });
}

DEF: int16x4
vrev32_s16 (int16x4_t a)
{
  return builtin_shuffle (a, (uint16x4_t) { 1, 0, 3, 2 });
}

DEF: uint8x8
vrev32_u8 (uint8x8_t a)
{
  return builtin_shuffle (a, (uint8x8_t) { 3, 2, 1, 0, 7, 6, 5, 4 });
}

DEF: uint16x4
vrev32_u16 (uint16x4_t a)
{
  return builtin_shuffle (a, (uint16x4_t) { 1, 0, 3, 2 });
}

DEF: poly8x16
vrev32q_p8 (poly8x16_t a)
{
  return builtin_shuffle (a,
      (uint8x16_t) { 3, 2, 1, 0, 7, 6, 5, 4, 11, 10, 9, 8, 15, 14, 13, 12 });
}

DEF: poly16x8
vrev32q_p16 (poly16x8_t a)
{
  return builtin_shuffle (a, (uint16x8_t) { 1, 0, 3, 2, 5, 4, 7, 6 });
}

DEF: int8x16
vrev32q_s8 (int8x16_t a)
{
  return builtin_shuffle (a,
      (uint8x16_t) { 3, 2, 1, 0, 7, 6, 5, 4, 11, 10, 9, 8, 15, 14, 13, 12 });
}

DEF: int16x8
vrev32q_s16 (int16x8_t a)
{
  return builtin_shuffle (a, (uint16x8_t) { 1, 0, 3, 2, 5, 4, 7, 6 });
}

DEF: uint8x16
vrev32q_u8 (uint8x16_t a)
{
  return builtin_shuffle (a,
      (uint8x16_t) { 3, 2, 1, 0, 7, 6, 5, 4, 11, 10, 9, 8, 15, 14, 13, 12 });
}

DEF: uint16x8
vrev32q_u16 (uint16x8_t a)
{
  return builtin_shuffle (a, (uint16x8_t) { 1, 0, 3, 2, 5, 4, 7, 6 });
}

DEF: float32x2
vrev64_f32 (float32x2_t a)
{
  return builtin_shuffle (a, (uint32x2_t) { 1, 0 });
}

DEF: poly8x8
vrev64_p8 (poly8x8_t a)
{
  return builtin_shuffle (a, (uint8x8_t) { 7, 6, 5, 4, 3, 2, 1, 0 });
}

DEF: poly16x4
vrev64_p16 (poly16x4_t a)
{
  return builtin_shuffle (a, (uint16x4_t) { 3, 2, 1, 0 });
}

DEF: int8x8
vrev64_s8 (int8x8_t a)
{
  return builtin_shuffle (a, (uint8x8_t) { 7, 6, 5, 4, 3, 2, 1, 0 });
}

DEF: int16x4
vrev64_s16 (int16x4_t a)
{
  return builtin_shuffle (a, (uint16x4_t) { 3, 2, 1, 0 });
}

DEF: int32x2
vrev64_s32 (int32x2_t a)
{
  return builtin_shuffle (a, (uint32x2_t) { 1, 0 });
}

DEF: uint8x8
vrev64_u8 (uint8x8_t a)
{
  return builtin_shuffle (a, (uint8x8_t) { 7, 6, 5, 4, 3, 2, 1, 0 });
}

DEF: uint16x4
vrev64_u16 (uint16x4_t a)
{
  return builtin_shuffle (a, (uint16x4_t) { 3, 2, 1, 0 });
}

DEF: uint32x2
vrev64_u32 (uint32x2_t a)
{
  return builtin_shuffle (a, (uint32x2_t) { 1, 0 });
}

DEF: float32x4
vrev64q_f32 (float32x4_t a)
{
  return builtin_shuffle (a, (uint32x4_t) { 1, 0, 3, 2 });
}

DEF: poly8x16
vrev64q_p8 (poly8x16_t a)
{
  return builtin_shuffle (a,
      (uint8x16_t) { 7, 6, 5, 4, 3, 2, 1, 0, 15, 14, 13, 12, 11, 10, 9, 8 });
}

DEF: poly16x8
vrev64q_p16 (poly16x8_t a)
{
  return builtin_shuffle (a, (uint16x8_t) { 3, 2, 1, 0, 7, 6, 5, 4 });
}

DEF: int8x16
vrev64q_s8 (int8x16_t a)
{
  return builtin_shuffle (a,
      (uint8x16_t) { 7, 6, 5, 4, 3, 2, 1, 0, 15, 14, 13, 12, 11, 10, 9, 8 });
}

DEF: int16x8
vrev64q_s16 (int16x8_t a)
{
  return builtin_shuffle (a, (uint16x8_t) { 3, 2, 1, 0, 7, 6, 5, 4 });
}

DEF: int32x4
vrev64q_s32 (int32x4_t a)
{
  return builtin_shuffle (a, (uint32x4_t) { 1, 0, 3, 2 });
}

DEF: uint8x16
vrev64q_u8 (uint8x16_t a)
{
  return builtin_shuffle (a,
      (uint8x16_t) { 7, 6, 5, 4, 3, 2, 1, 0, 15, 14, 13, 12, 11, 10, 9, 8 });
}

DEF: uint16x8
vrev64q_u16 (uint16x8_t a)
{
  return builtin_shuffle (a, (uint16x8_t) { 3, 2, 1, 0, 7, 6, 5, 4 });
}

DEF: uint32x4
vrev64q_u32 (uint32x4_t a)
{
  return builtin_shuffle (a, (uint32x4_t) { 1, 0, 3, 2 });
}

DEF: float32x2
vrnd_f32 (float32x2_t a)
{
  return builtin_aarch64_btruncv2sf (a);
}

DEF: float64x1
vrnd_f64 (float64x1_t a)
{
  return vset_lane_f64 (builtin_trunc (vget_lane_f64 (a, 0)), a, 0);
}

DEF: float32x4
vrndq_f32 (float32x4_t a)
{
  return builtin_aarch64_btruncv4sf (a);
}

DEF: float64x2
vrndq_f64 (float64x2_t a)
{
  return builtin_aarch64_btruncv2df (a);
}

DEF: float32x2
vrnda_f32 (float32x2_t a)
{
  return builtin_aarch64_roundv2sf (a);
}

DEF: float64x1
vrnda_f64 (float64x1_t a)
{
  return vset_lane_f64 (builtin_round (vget_lane_f64 (a, 0)), a, 0);
}

DEF: float32x4
vrndaq_f32 (float32x4_t a)
{
  return builtin_aarch64_roundv4sf (a);
}

DEF: float64x2
vrndaq_f64 (float64x2_t a)
{
  return builtin_aarch64_roundv2df (a);
}

DEF: float32x2
vrndi_f32 (float32x2_t a)
{
  return builtin_aarch64_nearbyintv2sf (a);
}

DEF: float64x1
vrndi_f64 (float64x1_t a)
{
  return vset_lane_f64 (builtin_nearbyint (vget_lane_f64 (a, 0)), a, 0);
}

DEF: float32x4
vrndiq_f32 (float32x4_t a)
{
  return builtin_aarch64_nearbyintv4sf (a);
}

DEF: float64x2
vrndiq_f64 (float64x2_t a)
{
  return builtin_aarch64_nearbyintv2df (a);
}

DEF: float32x2
vrndm_f32 (float32x2_t a)
{
  return builtin_aarch64_floorv2sf (a);
}

DEF: float64x1
vrndm_f64 (float64x1_t a)
{
  return vset_lane_f64 (builtin_floor (vget_lane_f64 (a, 0)), a, 0);
}

DEF: float32x4
vrndmq_f32 (float32x4_t a)
{
  return builtin_aarch64_floorv4sf (a);
}

DEF: float64x2
vrndmq_f64 (float64x2_t a)
{
  return builtin_aarch64_floorv2df (a);
}

DEF: float32x2
vrndn_f32 (float32x2_t a)
{
  return builtin_aarch64_frintnv2sf (a);
}

DEF: float64x1
vrndn_f64 (float64x1_t a)
{
  return (float64x1_t) {builtin_aarch64_frintndf (a[0])};
}

DEF: float32x4
vrndnq_f32 (float32x4_t a)
{
  return builtin_aarch64_frintnv4sf (a);
}

DEF: float64x2
vrndnq_f64 (float64x2_t a)
{
  return builtin_aarch64_frintnv2df (a);
}

DEF: float32x2
vrndp_f32 (float32x2_t a)
{
  return builtin_aarch64_ceilv2sf (a);
}

DEF: float64x1
vrndp_f64 (float64x1_t a)
{
  return vset_lane_f64 (builtin_ceil (vget_lane_f64 (a, 0)), a, 0);
}

DEF: float32x4
vrndpq_f32 (float32x4_t a)
{
  return builtin_aarch64_ceilv4sf (a);
}

DEF: float64x2
vrndpq_f64 (float64x2_t a)
{
  return builtin_aarch64_ceilv2df (a);
}

DEF: float32x2
vrndx_f32 (float32x2_t a)
{
  return builtin_aarch64_rintv2sf (a);
}

DEF: float64x1
vrndx_f64 (float64x1_t a)
{
  return vset_lane_f64 (builtin_rint (vget_lane_f64 (a, 0)), a, 0);
}

DEF: float32x4
vrndxq_f32 (float32x4_t a)
{
  return builtin_aarch64_rintv4sf (a);
}

DEF: float64x2
vrndxq_f64 (float64x2_t a)
{
  return builtin_aarch64_rintv2df (a);
}

DEF: int8x8
vrshl_s8 (int8x8_t a, int8x8_t b)
{
  return (int8x8_t) builtin_aarch64_srshlv8qi (a, b);
}

DEF: int16x4
vrshl_s16 (int16x4_t a, int16x4_t b)
{
  return (int16x4_t) builtin_aarch64_srshlv4hi (a, b);
}

DEF: int32x2
vrshl_s32 (int32x2_t a, int32x2_t b)
{
  return (int32x2_t) builtin_aarch64_srshlv2si (a, b);
}

DEF: int64x1
vrshl_s64 (int64x1_t a, int64x1_t b)
{
  return (int64x1_t) {builtin_aarch64_srshldi (a[0], b[0])};
}

DEF: uint8x8
vrshl_u8 (uint8x8_t a, int8x8_t b)
{
  return builtin_aarch64_urshlv8qi_uus (a, b);
}

DEF: uint16x4
vrshl_u16 (uint16x4_t a, int16x4_t b)
{
  return builtin_aarch64_urshlv4hi_uus (a, b);
}

DEF: uint32x2
vrshl_u32 (uint32x2_t a, int32x2_t b)
{
  return builtin_aarch64_urshlv2si_uus (a, b);
}

DEF: uint64x1
vrshl_u64 (uint64x1_t a, int64x1_t b)
{
  return (uint64x1_t) {builtin_aarch64_urshldi_uus (a[0], b[0])};
}

DEF: int8x16
vrshlq_s8 (int8x16_t a, int8x16_t b)
{
  return (int8x16_t) builtin_aarch64_srshlv16qi (a, b);
}

DEF: int16x8
vrshlq_s16 (int16x8_t a, int16x8_t b)
{
  return (int16x8_t) builtin_aarch64_srshlv8hi (a, b);
}

DEF: int32x4
vrshlq_s32 (int32x4_t a, int32x4_t b)
{
  return (int32x4_t) builtin_aarch64_srshlv4si (a, b);
}

DEF: int64x2
vrshlq_s64 (int64x2_t a, int64x2_t b)
{
  return (int64x2_t) builtin_aarch64_srshlv2di (a, b);
}

DEF: uint8x16
vrshlq_u8 (uint8x16_t a, int8x16_t b)
{
  return builtin_aarch64_urshlv16qi_uus (a, b);
}

DEF: uint16x8
vrshlq_u16 (uint16x8_t a, int16x8_t b)
{
  return builtin_aarch64_urshlv8hi_uus (a, b);
}

DEF: uint32x4
vrshlq_u32 (uint32x4_t a, int32x4_t b)
{
  return builtin_aarch64_urshlv4si_uus (a, b);
}

DEF: uint64x2
vrshlq_u64 (uint64x2_t a, int64x2_t b)
{
  return builtin_aarch64_urshlv2di_uus (a, b);
}

DEF: int64
vrshld_s64 (int64_t a, int64_t b)
{
  return builtin_aarch64_srshldi (a, b);
}

DEF: uint64
vrshld_u64 (uint64_t a, int64_t b)
{
  return builtin_aarch64_urshldi_uus (a, b);
}

DEF: int8x8
vrshr_n_s8 (int8x8_t a, const int b)
{
  return (int8x8_t) builtin_aarch64_srshr_nv8qi (a, b);
}

DEF: int16x4
vrshr_n_s16 (int16x4_t a, const int b)
{
  return (int16x4_t) builtin_aarch64_srshr_nv4hi (a, b);
}

DEF: int32x2
vrshr_n_s32 (int32x2_t a, const int b)
{
  return (int32x2_t) builtin_aarch64_srshr_nv2si (a, b);
}

DEF: int64x1
vrshr_n_s64 (int64x1_t a, const int b)
{
  return (int64x1_t) {builtin_aarch64_srshr_ndi (a[0], b)};
}

DEF: uint8x8
vrshr_n_u8 (uint8x8_t a, const int b)
{
  return builtin_aarch64_urshr_nv8qi_uus (a, b);
}

DEF: uint16x4
vrshr_n_u16 (uint16x4_t a, const int b)
{
  return builtin_aarch64_urshr_nv4hi_uus (a, b);
}

DEF: uint32x2
vrshr_n_u32 (uint32x2_t a, const int b)
{
  return builtin_aarch64_urshr_nv2si_uus (a, b);
}

DEF: uint64x1
vrshr_n_u64 (uint64x1_t a, const int b)
{
  return (uint64x1_t) {builtin_aarch64_urshr_ndi_uus (a[0], b)};
}

DEF: int8x16
vrshrq_n_s8 (int8x16_t a, const int b)
{
  return (int8x16_t) builtin_aarch64_srshr_nv16qi (a, b);
}

DEF: int16x8
vrshrq_n_s16 (int16x8_t a, const int b)
{
  return (int16x8_t) builtin_aarch64_srshr_nv8hi (a, b);
}

DEF: int32x4
vrshrq_n_s32 (int32x4_t a, const int b)
{
  return (int32x4_t) builtin_aarch64_srshr_nv4si (a, b);
}

DEF: int64x2
vrshrq_n_s64 (int64x2_t a, const int b)
{
  return (int64x2_t) builtin_aarch64_srshr_nv2di (a, b);
}

DEF: uint8x16
vrshrq_n_u8 (uint8x16_t a, const int b)
{
  return builtin_aarch64_urshr_nv16qi_uus (a, b);
}

DEF: uint16x8
vrshrq_n_u16 (uint16x8_t a, const int b)
{
  return builtin_aarch64_urshr_nv8hi_uus (a, b);
}

DEF: uint32x4
vrshrq_n_u32 (uint32x4_t a, const int b)
{
  return builtin_aarch64_urshr_nv4si_uus (a, b);
}

DEF: uint64x2
vrshrq_n_u64 (uint64x2_t a, const int b)
{
  return builtin_aarch64_urshr_nv2di_uus (a, b);
}

DEF: int64
vrshrd_n_s64 (int64_t a, const int b)
{
  return builtin_aarch64_srshr_ndi (a, b);
}

DEF: uint64
vrshrd_n_u64 (uint64_t a, const int b)
{
  return builtin_aarch64_urshr_ndi_uus (a, b);
}

DEF: int8x8
vrsra_n_s8 (int8x8_t a, int8x8_t b, const int c)
{
  return (int8x8_t) builtin_aarch64_srsra_nv8qi (a, b, c);
}

DEF: int16x4
vrsra_n_s16 (int16x4_t a, int16x4_t b, const int c)
{
  return (int16x4_t) builtin_aarch64_srsra_nv4hi (a, b, c);
}

DEF: int32x2
vrsra_n_s32 (int32x2_t a, int32x2_t b, const int c)
{
  return (int32x2_t) builtin_aarch64_srsra_nv2si (a, b, c);
}

DEF: int64x1
vrsra_n_s64 (int64x1_t a, int64x1_t b, const int c)
{
  return (int64x1_t) {builtin_aarch64_srsra_ndi (a[0], b[0], c)};
}

DEF: uint8x8
vrsra_n_u8 (uint8x8_t a, uint8x8_t b, const int c)
{
  return builtin_aarch64_ursra_nv8qi_uuus (a, b, c);
}

DEF: uint16x4
vrsra_n_u16 (uint16x4_t a, uint16x4_t b, const int c)
{
  return builtin_aarch64_ursra_nv4hi_uuus (a, b, c);
}

DEF: uint32x2
vrsra_n_u32 (uint32x2_t a, uint32x2_t b, const int c)
{
  return builtin_aarch64_ursra_nv2si_uuus (a, b, c);
}

DEF: uint64x1
vrsra_n_u64 (uint64x1_t a, uint64x1_t b, const int c)
{
  return (uint64x1_t) {builtin_aarch64_ursra_ndi_uuus (a[0], b[0], c)};
}

DEF: int8x16
vrsraq_n_s8 (int8x16_t a, int8x16_t b, const int c)
{
  return (int8x16_t) builtin_aarch64_srsra_nv16qi (a, b, c);
}

DEF: int16x8
vrsraq_n_s16 (int16x8_t a, int16x8_t b, const int c)
{
  return (int16x8_t) builtin_aarch64_srsra_nv8hi (a, b, c);
}

DEF: int32x4
vrsraq_n_s32 (int32x4_t a, int32x4_t b, const int c)
{
  return (int32x4_t) builtin_aarch64_srsra_nv4si (a, b, c);
}

DEF: int64x2
vrsraq_n_s64 (int64x2_t a, int64x2_t b, const int c)
{
  return (int64x2_t) builtin_aarch64_srsra_nv2di (a, b, c);
}

DEF: uint8x16
vrsraq_n_u8 (uint8x16_t a, uint8x16_t b, const int c)
{
  return builtin_aarch64_ursra_nv16qi_uuus (a, b, c);
}

DEF: uint16x8
vrsraq_n_u16 (uint16x8_t a, uint16x8_t b, const int c)
{
  return builtin_aarch64_ursra_nv8hi_uuus (a, b, c);
}

DEF: uint32x4
vrsraq_n_u32 (uint32x4_t a, uint32x4_t b, const int c)
{
  return builtin_aarch64_ursra_nv4si_uuus (a, b, c);
}

DEF: uint64x2
vrsraq_n_u64 (uint64x2_t a, uint64x2_t b, const int c)
{
  return builtin_aarch64_ursra_nv2di_uuus (a, b, c);
}

DEF: int64
vrsrad_n_s64 (int64_t a, int64_t b, const int c)
{
  return builtin_aarch64_srsra_ndi (a, b, c);
}

DEF: uint64
vrsrad_n_u64 (uint64_t a, uint64_t b, const int c)
{
  return builtin_aarch64_ursra_ndi_uuus (a, b, c);
}

#pragma GCC push_options
#pragma GCC target ("+nothing+crypto")

DEF: uint32x4
vsha1cq_u32 (uint32x4_t hash_abcd, uint32_t hash_e, uint32x4_t wk)
{
  return builtin_aarch64_crypto_sha1cv4si_uuuu (hash_abcd, hash_e, wk);
}

DEF: uint32x4
vsha1mq_u32 (uint32x4_t hash_abcd, uint32_t hash_e, uint32x4_t wk)
{
  return builtin_aarch64_crypto_sha1mv4si_uuuu (hash_abcd, hash_e, wk);
}

DEF: uint32x4
vsha1pq_u32 (uint32x4_t hash_abcd, uint32_t hash_e, uint32x4_t wk)
{
  return builtin_aarch64_crypto_sha1pv4si_uuuu (hash_abcd, hash_e, wk);
}

DEF: uint32
vsha1h_u32 (uint32_t hash_e)
{
  return builtin_aarch64_crypto_sha1hsi_uu (hash_e);
}

DEF: uint32x4
vsha1su0q_u32 (uint32x4_t w0_3, uint32x4_t w4_7, uint32x4_t w8_11)
{
  return builtin_aarch64_crypto_sha1su0v4si_uuuu (w0_3, w4_7, w8_11);
}

DEF: uint32x4
vsha1su1q_u32 (uint32x4_t tw0_3, uint32x4_t w12_15)
{
  return builtin_aarch64_crypto_sha1su1v4si_uuu (tw0_3, w12_15);
}

DEF: uint32x4
vsha256hq_u32 (uint32x4_t hash_abcd, uint32x4_t hash_efgh, uint32x4_t wk)
{
  return builtin_aarch64_crypto_sha256hv4si_uuuu (hash_abcd, hash_efgh, wk);
}

DEF: uint32x4
vsha256h2q_u32 (uint32x4_t hash_efgh, uint32x4_t hash_abcd, uint32x4_t wk)
{
  return builtin_aarch64_crypto_sha256h2v4si_uuuu (hash_efgh, hash_abcd, wk);
}

DEF: uint32x4
vsha256su0q_u32 (uint32x4_t w0_3, uint32x4_t w4_7)
{
  return builtin_aarch64_crypto_sha256su0v4si_uuu (w0_3, w4_7);
}

DEF: uint32x4
vsha256su1q_u32 (uint32x4_t tw0_3, uint32x4_t w8_11, uint32x4_t w12_15)
{
  return builtin_aarch64_crypto_sha256su1v4si_uuuu (tw0_3, w8_11, w12_15);
}

DEF: poly128
vmull_p64 (poly64_t a, poly64_t b)
{
  return
    builtin_aarch64_crypto_pmulldi_ppp (a, b);
}

DEF: poly128
vmull_high_p64 (poly64x2_t a, poly64x2_t b)
{
  return builtin_aarch64_crypto_pmullv2di_ppp (a, b);
}

#pragma GCC pop_options

DEF: int8x8
vshl_n_s8 (int8x8_t a, const int b)
{
  return (int8x8_t) builtin_aarch64_ashlv8qi (a, b);
}

DEF: int16x4
vshl_n_s16 (int16x4_t a, const int b)
{
  return (int16x4_t) builtin_aarch64_ashlv4hi (a, b);
}

DEF: int32x2
vshl_n_s32 (int32x2_t a, const int b)
{
  return (int32x2_t) builtin_aarch64_ashlv2si (a, b);
}

DEF: int64x1
vshl_n_s64 (int64x1_t a, const int b)
{
  return (int64x1_t) {builtin_aarch64_ashldi (a[0], b)};
}

DEF: uint8x8
vshl_n_u8 (uint8x8_t a, const int b)
{
  return (uint8x8_t) builtin_aarch64_ashlv8qi ((int8x8_t) a, b);
}

DEF: uint16x4
vshl_n_u16 (uint16x4_t a, const int b)
{
  return (uint16x4_t) builtin_aarch64_ashlv4hi ((int16x4_t) a, b);
}

DEF: uint32x2
vshl_n_u32 (uint32x2_t a, const int b)
{
  return (uint32x2_t) builtin_aarch64_ashlv2si ((int32x2_t) a, b);
}

DEF: uint64x1
vshl_n_u64 (uint64x1_t a, const int b)
{
  return (uint64x1_t) {builtin_aarch64_ashldi ((int64_t) a[0], b)};
}

DEF: int8x16
vshlq_n_s8 (int8x16_t a, const int b)
{
  return (int8x16_t) builtin_aarch64_ashlv16qi (a, b);
}

DEF: int16x8
vshlq_n_s16 (int16x8_t a, const int b)
{
  return (int16x8_t) builtin_aarch64_ashlv8hi (a, b);
}

DEF: int32x4
vshlq_n_s32 (int32x4_t a, const int b)
{
  return (int32x4_t) builtin_aarch64_ashlv4si (a, b);
}

DEF: int64x2
vshlq_n_s64 (int64x2_t a, const int b)
{
  return (int64x2_t) builtin_aarch64_ashlv2di (a, b);
}

DEF: uint8x16
vshlq_n_u8 (uint8x16_t a, const int b)
{
  return (uint8x16_t) builtin_aarch64_ashlv16qi ((int8x16_t) a, b);
}

DEF: uint16x8
vshlq_n_u16 (uint16x8_t a, const int b)
{
  return (uint16x8_t) builtin_aarch64_ashlv8hi ((int16x8_t) a, b);
}

DEF: uint32x4
vshlq_n_u32 (uint32x4_t a, const int b)
{
  return (uint32x4_t) builtin_aarch64_ashlv4si ((int32x4_t) a, b);
}

DEF: uint64x2
vshlq_n_u64 (uint64x2_t a, const int b)
{
  return (uint64x2_t) builtin_aarch64_ashlv2di ((int64x2_t) a, b);
}

DEF: int64
vshld_n_s64 (int64_t a, const int b)
{
  return builtin_aarch64_ashldi (a, b);
}

DEF: uint64
vshld_n_u64 (uint64_t a, const int b)
{
  return (uint64_t) builtin_aarch64_ashldi (a, b);
}

DEF: int8x8
vshl_s8 (int8x8_t a, int8x8_t b)
{
  return builtin_aarch64_sshlv8qi (a, b);
}

DEF: int16x4
vshl_s16 (int16x4_t a, int16x4_t b)
{
  return builtin_aarch64_sshlv4hi (a, b);
}

DEF: int32x2
vshl_s32 (int32x2_t a, int32x2_t b)
{
  return builtin_aarch64_sshlv2si (a, b);
}

DEF: int64x1
vshl_s64 (int64x1_t a, int64x1_t b)
{
  return (int64x1_t) {builtin_aarch64_sshldi (a[0], b[0])};
}

DEF: uint8x8
vshl_u8 (uint8x8_t a, int8x8_t b)
{
  return builtin_aarch64_ushlv8qi_uus (a, b);
}

DEF: uint16x4
vshl_u16 (uint16x4_t a, int16x4_t b)
{
  return builtin_aarch64_ushlv4hi_uus (a, b);
}

DEF: uint32x2
vshl_u32 (uint32x2_t a, int32x2_t b)
{
  return builtin_aarch64_ushlv2si_uus (a, b);
}

DEF: uint64x1
vshl_u64 (uint64x1_t a, int64x1_t b)
{
  return (uint64x1_t) {builtin_aarch64_ushldi_uus (a[0], b[0])};
}

DEF: int8x16
vshlq_s8 (int8x16_t a, int8x16_t b)
{
  return builtin_aarch64_sshlv16qi (a, b);
}

DEF: int16x8
vshlq_s16 (int16x8_t a, int16x8_t b)
{
  return builtin_aarch64_sshlv8hi (a, b);
}

DEF: int32x4
vshlq_s32 (int32x4_t a, int32x4_t b)
{
  return builtin_aarch64_sshlv4si (a, b);
}

DEF: int64x2
vshlq_s64 (int64x2_t a, int64x2_t b)
{
  return builtin_aarch64_sshlv2di (a, b);
}

DEF: uint8x16
vshlq_u8 (uint8x16_t a, int8x16_t b)
{
  return builtin_aarch64_ushlv16qi_uus (a, b);
}

DEF: uint16x8
vshlq_u16 (uint16x8_t a, int16x8_t b)
{
  return builtin_aarch64_ushlv8hi_uus (a, b);
}

DEF: uint32x4
vshlq_u32 (uint32x4_t a, int32x4_t b)
{
  return builtin_aarch64_ushlv4si_uus (a, b);
}

DEF: uint64x2
vshlq_u64 (uint64x2_t a, int64x2_t b)
{
  return builtin_aarch64_ushlv2di_uus (a, b);
}

DEF: int64
vshld_s64 (int64_t a, int64_t b)
{
  return builtin_aarch64_sshldi (a, b);
}

DEF: uint64
vshld_u64 (uint64_t a, uint64_t b)
{
  return builtin_aarch64_ushldi_uus (a, b);
}

DEF: int16x8
vshll_high_n_s8 (int8x16_t a, const int b)
{
  return builtin_aarch64_sshll2_nv16qi (a, b);
}

DEF: int32x4
vshll_high_n_s16 (int16x8_t a, const int b)
{
  return builtin_aarch64_sshll2_nv8hi (a, b);
}

DEF: int64x2
vshll_high_n_s32 (int32x4_t a, const int b)
{
  return builtin_aarch64_sshll2_nv4si (a, b);
}

DEF: uint16x8
vshll_high_n_u8 (uint8x16_t a, const int b)
{
  return (uint16x8_t) builtin_aarch64_ushll2_nv16qi ((int8x16_t) a, b);
}

DEF: uint32x4
vshll_high_n_u16 (uint16x8_t a, const int b)
{
  return (uint32x4_t) builtin_aarch64_ushll2_nv8hi ((int16x8_t) a, b);
}

DEF: uint64x2
vshll_high_n_u32 (uint32x4_t a, const int b)
{
  return (uint64x2_t) builtin_aarch64_ushll2_nv4si ((int32x4_t) a, b);
}

DEF: int16x8
vshll_n_s8 (int8x8_t a, const int b)
{
  return builtin_aarch64_sshll_nv8qi (a, b);
}

DEF: int32x4
vshll_n_s16 (int16x4_t a, const int b)
{
  return builtin_aarch64_sshll_nv4hi (a, b);
}

DEF: int64x2
vshll_n_s32 (int32x2_t a, const int b)
{
  return builtin_aarch64_sshll_nv2si (a, b);
}

DEF: uint16x8
vshll_n_u8 (uint8x8_t a, const int b)
{
  return builtin_aarch64_ushll_nv8qi_uus (a, b);
}

DEF: uint32x4
vshll_n_u16 (uint16x4_t a, const int b)
{
  return builtin_aarch64_ushll_nv4hi_uus (a, b);
}

DEF: uint64x2
vshll_n_u32 (uint32x2_t a, const int b)
{
  return builtin_aarch64_ushll_nv2si_uus (a, b);
}

DEF: int8x8
vshr_n_s8 (int8x8_t a, const int b)
{
  return (int8x8_t) builtin_aarch64_ashrv8qi (a, b);
}

DEF: int16x4
vshr_n_s16 (int16x4_t a, const int b)
{
  return (int16x4_t) builtin_aarch64_ashrv4hi (a, b);
}

DEF: int32x2
vshr_n_s32 (int32x2_t a, const int b)
{
  return (int32x2_t) builtin_aarch64_ashrv2si (a, b);
}

DEF: int64x1
vshr_n_s64 (int64x1_t a, const int b)
{
  return (int64x1_t) {builtin_aarch64_ashr_simddi (a[0], b)};
}

DEF: uint8x8
vshr_n_u8 (uint8x8_t a, const int b)
{
  return (uint8x8_t) builtin_aarch64_lshrv8qi ((int8x8_t) a, b);
}

DEF: uint16x4
vshr_n_u16 (uint16x4_t a, const int b)
{
  return (uint16x4_t) builtin_aarch64_lshrv4hi ((int16x4_t) a, b);
}

DEF: uint32x2
vshr_n_u32 (uint32x2_t a, const int b)
{
  return (uint32x2_t) builtin_aarch64_lshrv2si ((int32x2_t) a, b);
}

DEF: uint64x1
vshr_n_u64 (uint64x1_t a, const int b)
{
  return (uint64x1_t) {builtin_aarch64_lshr_simddi_uus ( a[0], b)};
}

DEF: int8x16
vshrq_n_s8 (int8x16_t a, const int b)
{
  return (int8x16_t) builtin_aarch64_ashrv16qi (a, b);
}

DEF: int16x8
vshrq_n_s16 (int16x8_t a, const int b)
{
  return (int16x8_t) builtin_aarch64_ashrv8hi (a, b);
}

DEF: int32x4
vshrq_n_s32 (int32x4_t a, const int b)
{
  return (int32x4_t) builtin_aarch64_ashrv4si (a, b);
}

DEF: int64x2
vshrq_n_s64 (int64x2_t a, const int b)
{
  return (int64x2_t) builtin_aarch64_ashrv2di (a, b);
}

DEF: uint8x16
vshrq_n_u8 (uint8x16_t a, const int b)
{
  return (uint8x16_t) builtin_aarch64_lshrv16qi ((int8x16_t) a, b);
}

DEF: uint16x8
vshrq_n_u16 (uint16x8_t a, const int b)
{
  return (uint16x8_t) builtin_aarch64_lshrv8hi ((int16x8_t) a, b);
}

DEF: uint32x4
vshrq_n_u32 (uint32x4_t a, const int b)
{
  return (uint32x4_t) builtin_aarch64_lshrv4si ((int32x4_t) a, b);
}

DEF: uint64x2
vshrq_n_u64 (uint64x2_t a, const int b)
{
  return (uint64x2_t) builtin_aarch64_lshrv2di ((int64x2_t) a, b);
}

DEF: int64
vshrd_n_s64 (int64_t a, const int b)
{
  return builtin_aarch64_ashr_simddi (a, b);
}

DEF: uint64
vshrd_n_u64 (uint64_t a, const int b)
{
  return builtin_aarch64_lshr_simddi_uus (a, b);
}

DEF: int8x8
vsli_n_s8 (int8x8_t a, int8x8_t b, const int c)
{
  return (int8x8_t) builtin_aarch64_ssli_nv8qi (a, b, c);
}

DEF: int16x4
vsli_n_s16 (int16x4_t a, int16x4_t b, const int c)
{
  return (int16x4_t) builtin_aarch64_ssli_nv4hi (a, b, c);
}

DEF: int32x2
vsli_n_s32 (int32x2_t a, int32x2_t b, const int c)
{
  return (int32x2_t) builtin_aarch64_ssli_nv2si (a, b, c);
}

DEF: int64x1
vsli_n_s64 (int64x1_t a, int64x1_t b, const int c)
{
  return (int64x1_t) {builtin_aarch64_ssli_ndi (a[0], b[0], c)};
}

DEF: uint8x8
vsli_n_u8 (uint8x8_t a, uint8x8_t b, const int c)
{
  return builtin_aarch64_usli_nv8qi_uuus (a, b, c);
}

DEF: uint16x4
vsli_n_u16 (uint16x4_t a, uint16x4_t b, const int c)
{
  return builtin_aarch64_usli_nv4hi_uuus (a, b, c);
}

DEF: uint32x2
vsli_n_u32 (uint32x2_t a, uint32x2_t b, const int c)
{
  return builtin_aarch64_usli_nv2si_uuus (a, b, c);
}

DEF: uint64x1
vsli_n_u64 (uint64x1_t a, uint64x1_t b, const int c)
{
  return (uint64x1_t) {builtin_aarch64_usli_ndi_uuus (a[0], b[0], c)};
}

DEF: int8x16
vsliq_n_s8 (int8x16_t a, int8x16_t b, const int c)
{
  return (int8x16_t) builtin_aarch64_ssli_nv16qi (a, b, c);
}

DEF: int16x8
vsliq_n_s16 (int16x8_t a, int16x8_t b, const int c)
{
  return (int16x8_t) builtin_aarch64_ssli_nv8hi (a, b, c);
}

DEF: int32x4
vsliq_n_s32 (int32x4_t a, int32x4_t b, const int c)
{
  return (int32x4_t) builtin_aarch64_ssli_nv4si (a, b, c);
}

DEF: int64x2
vsliq_n_s64 (int64x2_t a, int64x2_t b, const int c)
{
  return (int64x2_t) builtin_aarch64_ssli_nv2di (a, b, c);
}

DEF: uint8x16
vsliq_n_u8 (uint8x16_t a, uint8x16_t b, const int c)
{
  return builtin_aarch64_usli_nv16qi_uuus (a, b, c);
}

DEF: uint16x8
vsliq_n_u16 (uint16x8_t a, uint16x8_t b, const int c)
{
  return builtin_aarch64_usli_nv8hi_uuus (a, b, c);
}

DEF: uint32x4
vsliq_n_u32 (uint32x4_t a, uint32x4_t b, const int c)
{
  return builtin_aarch64_usli_nv4si_uuus (a, b, c);
}

DEF: uint64x2
vsliq_n_u64 (uint64x2_t a, uint64x2_t b, const int c)
{
  return builtin_aarch64_usli_nv2di_uuus (a, b, c);
}

DEF: int64
vslid_n_s64 (int64_t a, int64_t b, const int c)
{
  return builtin_aarch64_ssli_ndi (a, b, c);
}

DEF: uint64
vslid_n_u64 (uint64_t a, uint64_t b, const int c)
{
  return builtin_aarch64_usli_ndi_uuus (a, b, c);
}

DEF: uint8x8
vsqadd_u8 (uint8x8_t a, int8x8_t b)
{
  return builtin_aarch64_usqaddv8qi_uus (a, b);
}

DEF: uint16x4
vsqadd_u16 (uint16x4_t a, int16x4_t b)
{
  return builtin_aarch64_usqaddv4hi_uus (a, b);
}

DEF: uint32x2
vsqadd_u32 (uint32x2_t a, int32x2_t b)
{
  return builtin_aarch64_usqaddv2si_uus (a, b);
}

DEF: uint64x1
vsqadd_u64 (uint64x1_t a, int64x1_t b)
{
  return (uint64x1_t) {builtin_aarch64_usqadddi_uus (a[0], b[0])};
}

DEF: uint8x16
vsqaddq_u8 (uint8x16_t a, int8x16_t b)
{
  return builtin_aarch64_usqaddv16qi_uus (a, b);
}

DEF: uint16x8
vsqaddq_u16 (uint16x8_t a, int16x8_t b)
{
  return builtin_aarch64_usqaddv8hi_uus (a, b);
}

DEF: uint32x4
vsqaddq_u32 (uint32x4_t a, int32x4_t b)
{
  return builtin_aarch64_usqaddv4si_uus (a, b);
}

DEF: uint64x2
vsqaddq_u64 (uint64x2_t a, int64x2_t b)
{
  return builtin_aarch64_usqaddv2di_uus (a, b);
}

DEF: uint8
vsqaddb_u8 (uint8_t a, int8_t b)
{
  return builtin_aarch64_usqaddqi_uus (a, b);
}

DEF: uint16
vsqaddh_u16 (uint16_t a, int16_t b)
{
  return builtin_aarch64_usqaddhi_uus (a, b);
}

DEF: uint32
vsqadds_u32 (uint32_t a, int32_t b)
{
  return builtin_aarch64_usqaddsi_uus (a, b);
}

DEF: uint64
vsqaddd_u64 (uint64_t a, int64_t b)
{
  return builtin_aarch64_usqadddi_uus (a, b);
}
DEF: float32x2
vsqrt_f32 (float32x2_t a)
{
  return builtin_aarch64_sqrtv2sf (a);
}

DEF: float32x4
vsqrtq_f32 (float32x4_t a)
{
  return builtin_aarch64_sqrtv4sf (a);
}

DEF: float64x1
vsqrt_f64 (float64x1_t a)
{
  return (float64x1_t) { builtin_aarch64_sqrtdf (a[0]) };
}

DEF: float64x2
vsqrtq_f64 (float64x2_t a)
{
  return builtin_aarch64_sqrtv2df (a);
}

DEF: int8x8
vsra_n_s8 (int8x8_t a, int8x8_t b, const int c)
{
  return (int8x8_t) builtin_aarch64_ssra_nv8qi (a, b, c);
}

DEF: int16x4
vsra_n_s16 (int16x4_t a, int16x4_t b, const int c)
{
  return (int16x4_t) builtin_aarch64_ssra_nv4hi (a, b, c);
}

DEF: int32x2
vsra_n_s32 (int32x2_t a, int32x2_t b, const int c)
{
  return (int32x2_t) builtin_aarch64_ssra_nv2si (a, b, c);
}

DEF: int64x1
vsra_n_s64 (int64x1_t a, int64x1_t b, const int c)
{
  return (int64x1_t) {builtin_aarch64_ssra_ndi (a[0], b[0], c)};
}

DEF: uint8x8
vsra_n_u8 (uint8x8_t a, uint8x8_t b, const int c)
{
  return builtin_aarch64_usra_nv8qi_uuus (a, b, c);
}

DEF: uint16x4
vsra_n_u16 (uint16x4_t a, uint16x4_t b, const int c)
{
  return builtin_aarch64_usra_nv4hi_uuus (a, b, c);
}

DEF: uint32x2
vsra_n_u32 (uint32x2_t a, uint32x2_t b, const int c)
{
  return builtin_aarch64_usra_nv2si_uuus (a, b, c);
}

DEF: uint64x1
vsra_n_u64 (uint64x1_t a, uint64x1_t b, const int c)
{
  return (uint64x1_t) {builtin_aarch64_usra_ndi_uuus (a[0], b[0], c)};
}

DEF: int8x16
vsraq_n_s8 (int8x16_t a, int8x16_t b, const int c)
{
  return (int8x16_t) builtin_aarch64_ssra_nv16qi (a, b, c);
}

DEF: int16x8
vsraq_n_s16 (int16x8_t a, int16x8_t b, const int c)
{
  return (int16x8_t) builtin_aarch64_ssra_nv8hi (a, b, c);
}

DEF: int32x4
vsraq_n_s32 (int32x4_t a, int32x4_t b, const int c)
{
  return (int32x4_t) builtin_aarch64_ssra_nv4si (a, b, c);
}

DEF: int64x2
vsraq_n_s64 (int64x2_t a, int64x2_t b, const int c)
{
  return (int64x2_t) builtin_aarch64_ssra_nv2di (a, b, c);
}

DEF: uint8x16
vsraq_n_u8 (uint8x16_t a, uint8x16_t b, const int c)
{
  return builtin_aarch64_usra_nv16qi_uuus (a, b, c);
}

DEF: uint16x8
vsraq_n_u16 (uint16x8_t a, uint16x8_t b, const int c)
{
  return builtin_aarch64_usra_nv8hi_uuus (a, b, c);
}

DEF: uint32x4
vsraq_n_u32 (uint32x4_t a, uint32x4_t b, const int c)
{
  return builtin_aarch64_usra_nv4si_uuus (a, b, c);
}

DEF: uint64x2
vsraq_n_u64 (uint64x2_t a, uint64x2_t b, const int c)
{
  return builtin_aarch64_usra_nv2di_uuus (a, b, c);
}

DEF: int64
vsrad_n_s64 (int64_t a, int64_t b, const int c)
{
  return builtin_aarch64_ssra_ndi (a, b, c);
}

DEF: uint64
vsrad_n_u64 (uint64_t a, uint64_t b, const int c)
{
  return builtin_aarch64_usra_ndi_uuus (a, b, c);
}

DEF: int8x8
vsri_n_s8 (int8x8_t a, int8x8_t b, const int c)
{
  return (int8x8_t) builtin_aarch64_ssri_nv8qi (a, b, c);
}

DEF: int16x4
vsri_n_s16 (int16x4_t a, int16x4_t b, const int c)
{
  return (int16x4_t) builtin_aarch64_ssri_nv4hi (a, b, c);
}

DEF: int32x2
vsri_n_s32 (int32x2_t a, int32x2_t b, const int c)
{
  return (int32x2_t) builtin_aarch64_ssri_nv2si (a, b, c);
}

DEF: int64x1
vsri_n_s64 (int64x1_t a, int64x1_t b, const int c)
{
  return (int64x1_t) {builtin_aarch64_ssri_ndi (a[0], b[0], c)};
}

DEF: uint8x8
vsri_n_u8 (uint8x8_t a, uint8x8_t b, const int c)
{
  return builtin_aarch64_usri_nv8qi_uuus (a, b, c);
}

DEF: uint16x4
vsri_n_u16 (uint16x4_t a, uint16x4_t b, const int c)
{
  return builtin_aarch64_usri_nv4hi_uuus (a, b, c);
}

DEF: uint32x2
vsri_n_u32 (uint32x2_t a, uint32x2_t b, const int c)
{
  return builtin_aarch64_usri_nv2si_uuus (a, b, c);
}

DEF: uint64x1
vsri_n_u64 (uint64x1_t a, uint64x1_t b, const int c)
{
  return (uint64x1_t) {builtin_aarch64_usri_ndi_uuus (a[0], b[0], c)};
}

DEF: int8x16
vsriq_n_s8 (int8x16_t a, int8x16_t b, const int c)
{
  return (int8x16_t) builtin_aarch64_ssri_nv16qi (a, b, c);
}

DEF: int16x8
vsriq_n_s16 (int16x8_t a, int16x8_t b, const int c)
{
  return (int16x8_t) builtin_aarch64_ssri_nv8hi (a, b, c);
}

DEF: int32x4
vsriq_n_s32 (int32x4_t a, int32x4_t b, const int c)
{
  return (int32x4_t) builtin_aarch64_ssri_nv4si (a, b, c);
}

DEF: int64x2
vsriq_n_s64 (int64x2_t a, int64x2_t b, const int c)
{
  return (int64x2_t) builtin_aarch64_ssri_nv2di (a, b, c);
}

DEF: uint8x16
vsriq_n_u8 (uint8x16_t a, uint8x16_t b, const int c)
{
  return builtin_aarch64_usri_nv16qi_uuus (a, b, c);
}

DEF: uint16x8
vsriq_n_u16 (uint16x8_t a, uint16x8_t b, const int c)
{
  return builtin_aarch64_usri_nv8hi_uuus (a, b, c);
}

DEF: uint32x4
vsriq_n_u32 (uint32x4_t a, uint32x4_t b, const int c)
{
  return builtin_aarch64_usri_nv4si_uuus (a, b, c);
}

DEF: uint64x2
vsriq_n_u64 (uint64x2_t a, uint64x2_t b, const int c)
{
  return builtin_aarch64_usri_nv2di_uuus (a, b, c);
}

DEF: int64
vsrid_n_s64 (int64_t a, int64_t b, const int c)
{
  return builtin_aarch64_ssri_ndi (a, b, c);
}

DEF: uint64
vsrid_n_u64 (uint64_t a, uint64_t b, const int c)
{
  return builtin_aarch64_usri_ndi_uuus (a, b, c);
}

DEF: void
vst1_f32 (float32_t *a, float32x2_t b)
{
  builtin_aarch64_st1v2sf ((builtin_aarch64_simd_sf *) a, b);
}

DEF: void
vst1_f64 (float64_t *a, float64x1_t b)
{
  *a = b[0];
}

DEF: void
vst1_p8 (poly8_t *a, poly8x8_t b)
{
  builtin_aarch64_st1v8qi ((builtin_aarch64_simd_qi *) a,
        (int8x8_t) b);
}

DEF: void
vst1_p16 (poly16_t *a, poly16x4_t b)
{
  builtin_aarch64_st1v4hi ((builtin_aarch64_simd_hi *) a,
        (int16x4_t) b);
}

DEF: void
vst1_s8 (int8_t *a, int8x8_t b)
{
  builtin_aarch64_st1v8qi ((builtin_aarch64_simd_qi *) a, b);
}

DEF: void
vst1_s16 (int16_t *a, int16x4_t b)
{
  builtin_aarch64_st1v4hi ((builtin_aarch64_simd_hi *) a, b);
}

DEF: void
vst1_s32 (int32_t *a, int32x2_t b)
{
  builtin_aarch64_st1v2si ((builtin_aarch64_simd_si *) a, b);
}

DEF: void
vst1_s64 (int64_t *a, int64x1_t b)
{
  *a = b[0];
}

DEF: void
vst1_u8 (uint8_t *a, uint8x8_t b)
{
  builtin_aarch64_st1v8qi ((builtin_aarch64_simd_qi *) a,
        (int8x8_t) b);
}

DEF: void
vst1_u16 (uint16_t *a, uint16x4_t b)
{
  builtin_aarch64_st1v4hi ((builtin_aarch64_simd_hi *) a,
        (int16x4_t) b);
}

DEF: void
vst1_u32 (uint32_t *a, uint32x2_t b)
{
  builtin_aarch64_st1v2si ((builtin_aarch64_simd_si *) a,
        (int32x2_t) b);
}

DEF: void
vst1_u64 (uint64_t *a, uint64x1_t b)
{
  *a = b[0];
}

DEF: void
vst1q_f32 (float32_t *a, float32x4_t b)
{
  builtin_aarch64_st1v4sf ((builtin_aarch64_simd_sf *) a, b);
}

DEF: void
vst1q_f64 (float64_t *a, float64x2_t b)
{
  builtin_aarch64_st1v2df ((builtin_aarch64_simd_df *) a, b);
}

DEF: void
vst1q_p8 (poly8_t *a, poly8x16_t b)
{
  builtin_aarch64_st1v16qi ((builtin_aarch64_simd_qi *) a,
         (int8x16_t) b);
}

DEF: void
vst1q_p16 (poly16_t *a, poly16x8_t b)
{
  builtin_aarch64_st1v8hi ((builtin_aarch64_simd_hi *) a,
        (int16x8_t) b);
}

DEF: void
vst1q_s8 (int8_t *a, int8x16_t b)
{
  builtin_aarch64_st1v16qi ((builtin_aarch64_simd_qi *) a, b);
}

DEF: void
vst1q_s16 (int16_t *a, int16x8_t b)
{
  builtin_aarch64_st1v8hi ((builtin_aarch64_simd_hi *) a, b);
}

DEF: void
vst1q_s32 (int32_t *a, int32x4_t b)
{
  builtin_aarch64_st1v4si ((builtin_aarch64_simd_si *) a, b);
}

DEF: void
vst1q_s64 (int64_t *a, int64x2_t b)
{
  builtin_aarch64_st1v2di ((builtin_aarch64_simd_di *) a, b);
}

DEF: void
vst1q_u8 (uint8_t *a, uint8x16_t b)
{
  builtin_aarch64_st1v16qi ((builtin_aarch64_simd_qi *) a,
         (int8x16_t) b);
}

DEF: void
vst1q_u16 (uint16_t *a, uint16x8_t b)
{
  builtin_aarch64_st1v8hi ((builtin_aarch64_simd_hi *) a,
        (int16x8_t) b);
}

DEF: void
vst1q_u32 (uint32_t *a, uint32x4_t b)
{
  builtin_aarch64_st1v4si ((builtin_aarch64_simd_si *) a,
        (int32x4_t) b);
}

DEF: void
vst1q_u64 (uint64_t *a, uint64x2_t b)
{
  builtin_aarch64_st1v2di ((builtin_aarch64_simd_di *) a,
        (int64x2_t) b);
}

DEF: void
vst1_lane_f32 (float32_t *a, float32x2_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1_lane_f64 (float64_t *a, float64x1_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1_lane_p8 (poly8_t *a, poly8x8_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1_lane_p16 (poly16_t *a, poly16x4_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1_lane_s8 (int8_t *a, int8x8_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1_lane_s16 (int16_t *a, int16x4_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1_lane_s32 (int32_t *a, int32x2_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1_lane_s64 (int64_t *a, int64x1_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1_lane_u8 (uint8_t *a, uint8x8_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1_lane_u16 (uint16_t *a, uint16x4_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1_lane_u32 (uint32_t *a, uint32x2_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1_lane_u64 (uint64_t *a, uint64x1_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1q_lane_f32 (float32_t *a, float32x4_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1q_lane_f64 (float64_t *a, float64x2_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1q_lane_p8 (poly8_t *a, poly8x16_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1q_lane_p16 (poly16_t *a, poly16x8_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1q_lane_s8 (int8_t *a, int8x16_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1q_lane_s16 (int16_t *a, int16x8_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1q_lane_s32 (int32_t *a, int32x4_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1q_lane_s64 (int64_t *a, int64x2_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1q_lane_u8 (uint8_t *a, uint8x16_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1q_lane_u16 (uint16_t *a, uint16x8_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1q_lane_u32 (uint32_t *a, uint32x4_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst1q_lane_u64 (uint64_t *a, uint64x2_t b, const int lane)
{
  *a = extension ({ builtin_aarch64_im_lane_boundsi (sizeof(b), sizeof(b[0]), lane); b[lane]; });
}

DEF: void
vst2_s64 (int64_t * a, int64x1x2_t val)
{
  builtin_aarch64_simd_oi o;
  int64x2x2_t temp;
  temp.val[0] = vcombine_s64 (val.val[0], vcreate_s64 (((int64_t) 0)));
  temp.val[1] = vcombine_s64 (val.val[1], vcreate_s64 (((int64_t) 0)));
  o = builtin_aarch64_set_qregoiv2di (o, (int64x2_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregoiv2di (o, (int64x2_t) temp.val[1], 1);
  builtin_aarch64_st2di ((builtin_aarch64_simd_di *) a, o);
}

DEF: void
vst2_u64 (uint64_t * a, uint64x1x2_t val)
{
  builtin_aarch64_simd_oi o;
  uint64x2x2_t temp;
  temp.val[0] = vcombine_u64 (val.val[0], vcreate_u64 (((uint64_t) 0)));
  temp.val[1] = vcombine_u64 (val.val[1], vcreate_u64 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregoiv2di (o, (int64x2_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregoiv2di (o, (int64x2_t) temp.val[1], 1);
  builtin_aarch64_st2di ((builtin_aarch64_simd_di *) a, o);
}

DEF: void
vst2_f64 (float64_t * a, float64x1x2_t val)
{
  builtin_aarch64_simd_oi o;
  float64x2x2_t temp;
  temp.val[0] = vcombine_f64 (val.val[0], vcreate_f64 (((uint64_t) 0)));
  temp.val[1] = vcombine_f64 (val.val[1], vcreate_f64 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregoiv2df (o, (float64x2_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregoiv2df (o, (float64x2_t) temp.val[1], 1);
  builtin_aarch64_st2df ((builtin_aarch64_simd_df *) a, o);
}

DEF: void
vst2_s8 (int8_t * a, int8x8x2_t val)
{
  builtin_aarch64_simd_oi o;
  int8x16x2_t temp;
  temp.val[0] = vcombine_s8 (val.val[0], vcreate_s8 (((int64_t) 0)));
  temp.val[1] = vcombine_s8 (val.val[1], vcreate_s8 (((int64_t) 0)));
  o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) temp.val[1], 1);
  builtin_aarch64_st2v8qi ((builtin_aarch64_simd_qi *) a, o);
}

DEF: void
vst2_p8 (poly8_t * a, poly8x8x2_t val)
{
  builtin_aarch64_simd_oi o;
  poly8x16x2_t temp;
  temp.val[0] = vcombine_p8 (val.val[0], vcreate_p8 (((uint64_t) 0)));
  temp.val[1] = vcombine_p8 (val.val[1], vcreate_p8 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) temp.val[1], 1);
  builtin_aarch64_st2v8qi ((builtin_aarch64_simd_qi *) a, o);
}

DEF: void
vst2_s16 (int16_t * a, int16x4x2_t val)
{
  builtin_aarch64_simd_oi o;
  int16x8x2_t temp;
  temp.val[0] = vcombine_s16 (val.val[0], vcreate_s16 (((int64_t) 0)));
  temp.val[1] = vcombine_s16 (val.val[1], vcreate_s16 (((int64_t) 0)));
  o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) temp.val[1], 1);
  builtin_aarch64_st2v4hi ((builtin_aarch64_simd_hi *) a, o);
}

DEF: void
vst2_p16 (poly16_t * a, poly16x4x2_t val)
{
  builtin_aarch64_simd_oi o;
  poly16x8x2_t temp;
  temp.val[0] = vcombine_p16 (val.val[0], vcreate_p16 (((uint64_t) 0)));
  temp.val[1] = vcombine_p16 (val.val[1], vcreate_p16 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) temp.val[1], 1);
  builtin_aarch64_st2v4hi ((builtin_aarch64_simd_hi *) a, o);
}

DEF: void
vst2_s32 (int32_t * a, int32x2x2_t val)
{
  builtin_aarch64_simd_oi o;
  int32x4x2_t temp;
  temp.val[0] = vcombine_s32 (val.val[0], vcreate_s32 (((int64_t) 0)));
  temp.val[1] = vcombine_s32 (val.val[1], vcreate_s32 (((int64_t) 0)));
  o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) temp.val[1], 1);
  builtin_aarch64_st2v2si ((builtin_aarch64_simd_si *) a, o);
}

DEF: void
vst2_u8 (uint8_t * a, uint8x8x2_t val)
{
  builtin_aarch64_simd_oi o;
  uint8x16x2_t temp;
  temp.val[0] = vcombine_u8 (val.val[0], vcreate_u8 (((uint64_t) 0)));
  temp.val[1] = vcombine_u8 (val.val[1], vcreate_u8 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) temp.val[1], 1);
  builtin_aarch64_st2v8qi ((builtin_aarch64_simd_qi *) a, o);
}

DEF: void
vst2_u16 (uint16_t * a, uint16x4x2_t val)
{
  builtin_aarch64_simd_oi o;
  uint16x8x2_t temp;
  temp.val[0] = vcombine_u16 (val.val[0], vcreate_u16 (((uint64_t) 0)));
  temp.val[1] = vcombine_u16 (val.val[1], vcreate_u16 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) temp.val[1], 1);
  builtin_aarch64_st2v4hi ((builtin_aarch64_simd_hi *) a, o);
}

DEF: void
vst2_u32 (uint32_t * a, uint32x2x2_t val)
{
  builtin_aarch64_simd_oi o;
  uint32x4x2_t temp;
  temp.val[0] = vcombine_u32 (val.val[0], vcreate_u32 (((uint64_t) 0)));
  temp.val[1] = vcombine_u32 (val.val[1], vcreate_u32 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) temp.val[1], 1);
  builtin_aarch64_st2v2si ((builtin_aarch64_simd_si *) a, o);
}

DEF: void
vst2_f32 (float32_t * a, float32x2x2_t val)
{
  builtin_aarch64_simd_oi o;
  float32x4x2_t temp;
  temp.val[0] = vcombine_f32 (val.val[0], vcreate_f32 (((uint64_t) 0)));
  temp.val[1] = vcombine_f32 (val.val[1], vcreate_f32 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregoiv4sf (o, (float32x4_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregoiv4sf (o, (float32x4_t) temp.val[1], 1);
  builtin_aarch64_st2v2sf ((builtin_aarch64_simd_sf *) a, o);
}

DEF: void
vst2q_s8 (int8_t * a, int8x16x2_t val)
{
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) val.val[0], 0);
  o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) val.val[1], 1);
  builtin_aarch64_st2v16qi ((builtin_aarch64_simd_qi *) a, o);
}

DEF: void
vst2q_p8 (poly8_t * a, poly8x16x2_t val)
{
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) val.val[0], 0);
  o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) val.val[1], 1);
  builtin_aarch64_st2v16qi ((builtin_aarch64_simd_qi *) a, o);
}

DEF: void
vst2q_s16 (int16_t * a, int16x8x2_t val)
{
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) val.val[0], 0);
  o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) val.val[1], 1);
  builtin_aarch64_st2v8hi ((builtin_aarch64_simd_hi *) a, o);
}

DEF: void
vst2q_p16 (poly16_t * a, poly16x8x2_t val)
{
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) val.val[0], 0);
  o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) val.val[1], 1);
  builtin_aarch64_st2v8hi ((builtin_aarch64_simd_hi *) a, o);
}

DEF: void
vst2q_s32 (int32_t * a, int32x4x2_t val)
{
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) val.val[0], 0);
  o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) val.val[1], 1);
  builtin_aarch64_st2v4si ((builtin_aarch64_simd_si *) a, o);
}

DEF: void
vst2q_s64 (int64_t * a, int64x2x2_t val)
{
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_set_qregoiv2di (o, (int64x2_t) val.val[0], 0);
  o = builtin_aarch64_set_qregoiv2di (o, (int64x2_t) val.val[1], 1);
  builtin_aarch64_st2v2di ((builtin_aarch64_simd_di *) a, o);
}

DEF: void
vst2q_u8 (uint8_t * a, uint8x16x2_t val)
{
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) val.val[0], 0);
  o = builtin_aarch64_set_qregoiv16qi (o, (int8x16_t) val.val[1], 1);
  builtin_aarch64_st2v16qi ((builtin_aarch64_simd_qi *) a, o);
}

DEF: void
vst2q_u16 (uint16_t * a, uint16x8x2_t val)
{
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) val.val[0], 0);
  o = builtin_aarch64_set_qregoiv8hi (o, (int16x8_t) val.val[1], 1);
  builtin_aarch64_st2v8hi ((builtin_aarch64_simd_hi *) a, o);
}

DEF: void
vst2q_u32 (uint32_t * a, uint32x4x2_t val)
{
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) val.val[0], 0);
  o = builtin_aarch64_set_qregoiv4si (o, (int32x4_t) val.val[1], 1);
  builtin_aarch64_st2v4si ((builtin_aarch64_simd_si *) a, o);
}

DEF: void
vst2q_u64 (uint64_t * a, uint64x2x2_t val)
{
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_set_qregoiv2di (o, (int64x2_t) val.val[0], 0);
  o = builtin_aarch64_set_qregoiv2di (o, (int64x2_t) val.val[1], 1);
  builtin_aarch64_st2v2di ((builtin_aarch64_simd_di *) a, o);
}

DEF: void
vst2q_f32 (float32_t * a, float32x4x2_t val)
{
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_set_qregoiv4sf (o, (float32x4_t) val.val[0], 0);
  o = builtin_aarch64_set_qregoiv4sf (o, (float32x4_t) val.val[1], 1);
  builtin_aarch64_st2v4sf ((builtin_aarch64_simd_sf *) a, o);
}

DEF: void
vst2q_f64 (float64_t * a, float64x2x2_t val)
{
  builtin_aarch64_simd_oi o;
  o = builtin_aarch64_set_qregoiv2df (o, (float64x2_t) val.val[0], 0);
  o = builtin_aarch64_set_qregoiv2df (o, (float64x2_t) val.val[1], 1);
  builtin_aarch64_st2v2df ((builtin_aarch64_simd_df *) a, o);
}

DEF: void
vst3_s64 (int64_t * a, int64x1x3_t val)
{
  builtin_aarch64_simd_ci o;
  int64x2x3_t temp;
  temp.val[0] = vcombine_s64 (val.val[0], vcreate_s64 (((int64_t) 0)));
  temp.val[1] = vcombine_s64 (val.val[1], vcreate_s64 (((int64_t) 0)));
  temp.val[2] = vcombine_s64 (val.val[2], vcreate_s64 (((int64_t) 0)));
  o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) temp.val[2], 2);
  builtin_aarch64_st3di ((builtin_aarch64_simd_di *) a, o);
}

DEF: void
vst3_u64 (uint64_t * a, uint64x1x3_t val)
{
  builtin_aarch64_simd_ci o;
  uint64x2x3_t temp;
  temp.val[0] = vcombine_u64 (val.val[0], vcreate_u64 (((uint64_t) 0)));
  temp.val[1] = vcombine_u64 (val.val[1], vcreate_u64 (((uint64_t) 0)));
  temp.val[2] = vcombine_u64 (val.val[2], vcreate_u64 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) temp.val[2], 2);
  builtin_aarch64_st3di ((builtin_aarch64_simd_di *) a, o);
}

DEF: void
vst3_f64 (float64_t * a, float64x1x3_t val)
{
  builtin_aarch64_simd_ci o;
  float64x2x3_t temp;
  temp.val[0] = vcombine_f64 (val.val[0], vcreate_f64 (((uint64_t) 0)));
  temp.val[1] = vcombine_f64 (val.val[1], vcreate_f64 (((uint64_t) 0)));
  temp.val[2] = vcombine_f64 (val.val[2], vcreate_f64 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregciv2df (o, (float64x2_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregciv2df (o, (float64x2_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregciv2df (o, (float64x2_t) temp.val[2], 2);
  builtin_aarch64_st3df ((builtin_aarch64_simd_df *) a, o);
}

DEF: void
vst3_s8 (int8_t * a, int8x8x3_t val)
{
  builtin_aarch64_simd_ci o;
  int8x16x3_t temp;
  temp.val[0] = vcombine_s8 (val.val[0], vcreate_s8 (((int64_t) 0)));
  temp.val[1] = vcombine_s8 (val.val[1], vcreate_s8 (((int64_t) 0)));
  temp.val[2] = vcombine_s8 (val.val[2], vcreate_s8 (((int64_t) 0)));
  o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[2], 2);
  builtin_aarch64_st3v8qi ((builtin_aarch64_simd_qi *) a, o);
}

DEF: void
vst3_p8 (poly8_t * a, poly8x8x3_t val)
{
  builtin_aarch64_simd_ci o;
  poly8x16x3_t temp;
  temp.val[0] = vcombine_p8 (val.val[0], vcreate_p8 (((uint64_t) 0)));
  temp.val[1] = vcombine_p8 (val.val[1], vcreate_p8 (((uint64_t) 0)));
  temp.val[2] = vcombine_p8 (val.val[2], vcreate_p8 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[2], 2);
  builtin_aarch64_st3v8qi ((builtin_aarch64_simd_qi *) a, o);
}

DEF: void
vst3_s16 (int16_t * a, int16x4x3_t val)
{
  builtin_aarch64_simd_ci o;
  int16x8x3_t temp;
  temp.val[0] = vcombine_s16 (val.val[0], vcreate_s16 (((int64_t) 0)));
  temp.val[1] = vcombine_s16 (val.val[1], vcreate_s16 (((int64_t) 0)));
  temp.val[2] = vcombine_s16 (val.val[2], vcreate_s16 (((int64_t) 0)));
  o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[2], 2);
  builtin_aarch64_st3v4hi ((builtin_aarch64_simd_hi *) a, o);
}

DEF: void
vst3_p16 (poly16_t * a, poly16x4x3_t val)
{
  builtin_aarch64_simd_ci o;
  poly16x8x3_t temp;
  temp.val[0] = vcombine_p16 (val.val[0], vcreate_p16 (((uint64_t) 0)));
  temp.val[1] = vcombine_p16 (val.val[1], vcreate_p16 (((uint64_t) 0)));
  temp.val[2] = vcombine_p16 (val.val[2], vcreate_p16 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[2], 2);
  builtin_aarch64_st3v4hi ((builtin_aarch64_simd_hi *) a, o);
}

DEF: void
vst3_s32 (int32_t * a, int32x2x3_t val)
{
  builtin_aarch64_simd_ci o;
  int32x4x3_t temp;
  temp.val[0] = vcombine_s32 (val.val[0], vcreate_s32 (((int64_t) 0)));
  temp.val[1] = vcombine_s32 (val.val[1], vcreate_s32 (((int64_t) 0)));
  temp.val[2] = vcombine_s32 (val.val[2], vcreate_s32 (((int64_t) 0)));
  o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) temp.val[2], 2);
  builtin_aarch64_st3v2si ((builtin_aarch64_simd_si *) a, o);
}

DEF: void
vst3_u8 (uint8_t * a, uint8x8x3_t val)
{
  builtin_aarch64_simd_ci o;
  uint8x16x3_t temp;
  temp.val[0] = vcombine_u8 (val.val[0], vcreate_u8 (((uint64_t) 0)));
  temp.val[1] = vcombine_u8 (val.val[1], vcreate_u8 (((uint64_t) 0)));
  temp.val[2] = vcombine_u8 (val.val[2], vcreate_u8 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) temp.val[2], 2);
  builtin_aarch64_st3v8qi ((builtin_aarch64_simd_qi *) a, o);
}

DEF: void
vst3_u16 (uint16_t * a, uint16x4x3_t val)
{
  builtin_aarch64_simd_ci o;
  uint16x8x3_t temp;
  temp.val[0] = vcombine_u16 (val.val[0], vcreate_u16 (((uint64_t) 0)));
  temp.val[1] = vcombine_u16 (val.val[1], vcreate_u16 (((uint64_t) 0)));
  temp.val[2] = vcombine_u16 (val.val[2], vcreate_u16 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) temp.val[2], 2);
  builtin_aarch64_st3v4hi ((builtin_aarch64_simd_hi *) a, o);
}

DEF: void
vst3_u32 (uint32_t * a, uint32x2x3_t val)
{
  builtin_aarch64_simd_ci o;
  uint32x4x3_t temp;
  temp.val[0] = vcombine_u32 (val.val[0], vcreate_u32 (((uint64_t) 0)));
  temp.val[1] = vcombine_u32 (val.val[1], vcreate_u32 (((uint64_t) 0)));
  temp.val[2] = vcombine_u32 (val.val[2], vcreate_u32 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) temp.val[2], 2);
  builtin_aarch64_st3v2si ((builtin_aarch64_simd_si *) a, o);
}

DEF: void
vst3_f32 (float32_t * a, float32x2x3_t val)
{
  builtin_aarch64_simd_ci o;
  float32x4x3_t temp;
  temp.val[0] = vcombine_f32 (val.val[0], vcreate_f32 (((uint64_t) 0)));
  temp.val[1] = vcombine_f32 (val.val[1], vcreate_f32 (((uint64_t) 0)));
  temp.val[2] = vcombine_f32 (val.val[2], vcreate_f32 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregciv4sf (o, (float32x4_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregciv4sf (o, (float32x4_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregciv4sf (o, (float32x4_t) temp.val[2], 2);
  builtin_aarch64_st3v2sf ((builtin_aarch64_simd_sf *) a, o);
}

DEF: void
vst3q_s8 (int8_t * a, int8x16x3_t val)
{
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) val.val[0], 0);
  o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) val.val[1], 1);
  o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) val.val[2], 2);
  builtin_aarch64_st3v16qi ((builtin_aarch64_simd_qi *) a, o);
}

DEF: void
vst3q_p8 (poly8_t * a, poly8x16x3_t val)
{
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) val.val[0], 0);
  o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) val.val[1], 1);
  o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) val.val[2], 2);
  builtin_aarch64_st3v16qi ((builtin_aarch64_simd_qi *) a, o);
}

DEF: void
vst3q_s16 (int16_t * a, int16x8x3_t val)
{
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) val.val[0], 0);
  o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) val.val[1], 1);
  o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) val.val[2], 2);
  builtin_aarch64_st3v8hi ((builtin_aarch64_simd_hi *) a, o);
}

DEF: void
vst3q_p16 (poly16_t * a, poly16x8x3_t val)
{
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) val.val[0], 0);
  o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) val.val[1], 1);
  o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) val.val[2], 2);
  builtin_aarch64_st3v8hi ((builtin_aarch64_simd_hi *) a, o);
}

DEF: void
vst3q_s32 (int32_t * a, int32x4x3_t val)
{
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) val.val[0], 0);
  o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) val.val[1], 1);
  o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) val.val[2], 2);
  builtin_aarch64_st3v4si ((builtin_aarch64_simd_si *) a, o);
}

DEF: void
vst3q_s64 (int64_t * a, int64x2x3_t val)
{
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) val.val[0], 0);
  o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) val.val[1], 1);
  o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) val.val[2], 2);
  builtin_aarch64_st3v2di ((builtin_aarch64_simd_di *) a, o);
}

DEF: void
vst3q_u8 (uint8_t * a, uint8x16x3_t val)
{
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) val.val[0], 0);
  o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) val.val[1], 1);
  o = builtin_aarch64_set_qregciv16qi (o, (int8x16_t) val.val[2], 2);
  builtin_aarch64_st3v16qi ((builtin_aarch64_simd_qi *) a, o);
}

DEF: void
vst3q_u16 (uint16_t * a, uint16x8x3_t val)
{
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) val.val[0], 0);
  o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) val.val[1], 1);
  o = builtin_aarch64_set_qregciv8hi (o, (int16x8_t) val.val[2], 2);
  builtin_aarch64_st3v8hi ((builtin_aarch64_simd_hi *) a, o);
}

DEF: void
vst3q_u32 (uint32_t * a, uint32x4x3_t val)
{
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) val.val[0], 0);
  o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) val.val[1], 1);
  o = builtin_aarch64_set_qregciv4si (o, (int32x4_t) val.val[2], 2);
  builtin_aarch64_st3v4si ((builtin_aarch64_simd_si *) a, o);
}

DEF: void
vst3q_u64 (uint64_t * a, uint64x2x3_t val)
{
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) val.val[0], 0);
  o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) val.val[1], 1);
  o = builtin_aarch64_set_qregciv2di (o, (int64x2_t) val.val[2], 2);
  builtin_aarch64_st3v2di ((builtin_aarch64_simd_di *) a, o);
}

DEF: void
vst3q_f32 (float32_t * a, float32x4x3_t val)
{
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_set_qregciv4sf (o, (float32x4_t) val.val[0], 0);
  o = builtin_aarch64_set_qregciv4sf (o, (float32x4_t) val.val[1], 1);
  o = builtin_aarch64_set_qregciv4sf (o, (float32x4_t) val.val[2], 2);
  builtin_aarch64_st3v4sf ((builtin_aarch64_simd_sf *) a, o);
}

DEF: void
vst3q_f64 (float64_t * a, float64x2x3_t val)
{
  builtin_aarch64_simd_ci o;
  o = builtin_aarch64_set_qregciv2df (o, (float64x2_t) val.val[0], 0);
  o = builtin_aarch64_set_qregciv2df (o, (float64x2_t) val.val[1], 1);
  o = builtin_aarch64_set_qregciv2df (o, (float64x2_t) val.val[2], 2);
  builtin_aarch64_st3v2df ((builtin_aarch64_simd_df *) a, o);
}

DEF: void
vst4_s64 (int64_t * a, int64x1x4_t val)
{
  builtin_aarch64_simd_xi o;
  int64x2x4_t temp;
  temp.val[0] = vcombine_s64 (val.val[0], vcreate_s64 (((int64_t) 0)));
  temp.val[1] = vcombine_s64 (val.val[1], vcreate_s64 (((int64_t) 0)));
  temp.val[2] = vcombine_s64 (val.val[2], vcreate_s64 (((int64_t) 0)));
  temp.val[3] = vcombine_s64 (val.val[3], vcreate_s64 (((int64_t) 0)));
  o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[2], 2);
  o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[3], 3);
  builtin_aarch64_st4di ((builtin_aarch64_simd_di *) a, o);
}

DEF: void
vst4_u64 (uint64_t * a, uint64x1x4_t val)
{
  builtin_aarch64_simd_xi o;
  uint64x2x4_t temp;
  temp.val[0] = vcombine_u64 (val.val[0], vcreate_u64 (((uint64_t) 0)));
  temp.val[1] = vcombine_u64 (val.val[1], vcreate_u64 (((uint64_t) 0)));
  temp.val[2] = vcombine_u64 (val.val[2], vcreate_u64 (((uint64_t) 0)));
  temp.val[3] = vcombine_u64 (val.val[3], vcreate_u64 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[2], 2);
  o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) temp.val[3], 3);
  builtin_aarch64_st4di ((builtin_aarch64_simd_di *) a, o);
}

DEF: void
vst4_f64 (float64_t * a, float64x1x4_t val)
{
  builtin_aarch64_simd_xi o;
  float64x2x4_t temp;
  temp.val[0] = vcombine_f64 (val.val[0], vcreate_f64 (((uint64_t) 0)));
  temp.val[1] = vcombine_f64 (val.val[1], vcreate_f64 (((uint64_t) 0)));
  temp.val[2] = vcombine_f64 (val.val[2], vcreate_f64 (((uint64_t) 0)));
  temp.val[3] = vcombine_f64 (val.val[3], vcreate_f64 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregxiv2df (o, (float64x2_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregxiv2df (o, (float64x2_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregxiv2df (o, (float64x2_t) temp.val[2], 2);
  o = builtin_aarch64_set_qregxiv2df (o, (float64x2_t) temp.val[3], 3);
  builtin_aarch64_st4df ((builtin_aarch64_simd_df *) a, o);
}

DEF: void
vst4_s8 (int8_t * a, int8x8x4_t val)
{
  builtin_aarch64_simd_xi o;
  int8x16x4_t temp;
  temp.val[0] = vcombine_s8 (val.val[0], vcreate_s8 (((int64_t) 0)));
  temp.val[1] = vcombine_s8 (val.val[1], vcreate_s8 (((int64_t) 0)));
  temp.val[2] = vcombine_s8 (val.val[2], vcreate_s8 (((int64_t) 0)));
  temp.val[3] = vcombine_s8 (val.val[3], vcreate_s8 (((int64_t) 0)));
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[2], 2);
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[3], 3);
  builtin_aarch64_st4v8qi ((builtin_aarch64_simd_qi *) a, o);
}

DEF: void
vst4_p8 (poly8_t * a, poly8x8x4_t val)
{
  builtin_aarch64_simd_xi o;
  poly8x16x4_t temp;
  temp.val[0] = vcombine_p8 (val.val[0], vcreate_p8 (((uint64_t) 0)));
  temp.val[1] = vcombine_p8 (val.val[1], vcreate_p8 (((uint64_t) 0)));
  temp.val[2] = vcombine_p8 (val.val[2], vcreate_p8 (((uint64_t) 0)));
  temp.val[3] = vcombine_p8 (val.val[3], vcreate_p8 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[2], 2);
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[3], 3);
  builtin_aarch64_st4v8qi ((builtin_aarch64_simd_qi *) a, o);
}

DEF: void
vst4_s16 (int16_t * a, int16x4x4_t val)
{
  builtin_aarch64_simd_xi o;
  int16x8x4_t temp;
  temp.val[0] = vcombine_s16 (val.val[0], vcreate_s16 (((int64_t) 0)));
  temp.val[1] = vcombine_s16 (val.val[1], vcreate_s16 (((int64_t) 0)));
  temp.val[2] = vcombine_s16 (val.val[2], vcreate_s16 (((int64_t) 0)));
  temp.val[3] = vcombine_s16 (val.val[3], vcreate_s16 (((int64_t) 0)));
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[2], 2);
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[3], 3);
  builtin_aarch64_st4v4hi ((builtin_aarch64_simd_hi *) a, o);
}

DEF: void
vst4_p16 (poly16_t * a, poly16x4x4_t val)
{
  builtin_aarch64_simd_xi o;
  poly16x8x4_t temp;
  temp.val[0] = vcombine_p16 (val.val[0], vcreate_p16 (((uint64_t) 0)));
  temp.val[1] = vcombine_p16 (val.val[1], vcreate_p16 (((uint64_t) 0)));
  temp.val[2] = vcombine_p16 (val.val[2], vcreate_p16 (((uint64_t) 0)));
  temp.val[3] = vcombine_p16 (val.val[3], vcreate_p16 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[2], 2);
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[3], 3);
  builtin_aarch64_st4v4hi ((builtin_aarch64_simd_hi *) a, o);
}

DEF: void
vst4_s32 (int32_t * a, int32x2x4_t val)
{
  builtin_aarch64_simd_xi o;
  int32x4x4_t temp;
  temp.val[0] = vcombine_s32 (val.val[0], vcreate_s32 (((int64_t) 0)));
  temp.val[1] = vcombine_s32 (val.val[1], vcreate_s32 (((int64_t) 0)));
  temp.val[2] = vcombine_s32 (val.val[2], vcreate_s32 (((int64_t) 0)));
  temp.val[3] = vcombine_s32 (val.val[3], vcreate_s32 (((int64_t) 0)));
  o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[2], 2);
  o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[3], 3);
  builtin_aarch64_st4v2si ((builtin_aarch64_simd_si *) a, o);
}

DEF: void
vst4_u8 (uint8_t * a, uint8x8x4_t val)
{
  builtin_aarch64_simd_xi o;
  uint8x16x4_t temp;
  temp.val[0] = vcombine_u8 (val.val[0], vcreate_u8 (((uint64_t) 0)));
  temp.val[1] = vcombine_u8 (val.val[1], vcreate_u8 (((uint64_t) 0)));
  temp.val[2] = vcombine_u8 (val.val[2], vcreate_u8 (((uint64_t) 0)));
  temp.val[3] = vcombine_u8 (val.val[3], vcreate_u8 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[2], 2);
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) temp.val[3], 3);
  builtin_aarch64_st4v8qi ((builtin_aarch64_simd_qi *) a, o);
}

DEF: void
vst4_u16 (uint16_t * a, uint16x4x4_t val)
{
  builtin_aarch64_simd_xi o;
  uint16x8x4_t temp;
  temp.val[0] = vcombine_u16 (val.val[0], vcreate_u16 (((uint64_t) 0)));
  temp.val[1] = vcombine_u16 (val.val[1], vcreate_u16 (((uint64_t) 0)));
  temp.val[2] = vcombine_u16 (val.val[2], vcreate_u16 (((uint64_t) 0)));
  temp.val[3] = vcombine_u16 (val.val[3], vcreate_u16 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[2], 2);
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) temp.val[3], 3);
  builtin_aarch64_st4v4hi ((builtin_aarch64_simd_hi *) a, o);
}

DEF: void
vst4_u32 (uint32_t * a, uint32x2x4_t val)
{
  builtin_aarch64_simd_xi o;
  uint32x4x4_t temp;
  temp.val[0] = vcombine_u32 (val.val[0], vcreate_u32 (((uint64_t) 0)));
  temp.val[1] = vcombine_u32 (val.val[1], vcreate_u32 (((uint64_t) 0)));
  temp.val[2] = vcombine_u32 (val.val[2], vcreate_u32 (((uint64_t) 0)));
  temp.val[3] = vcombine_u32 (val.val[3], vcreate_u32 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[2], 2);
  o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) temp.val[3], 3);
  builtin_aarch64_st4v2si ((builtin_aarch64_simd_si *) a, o);
}

DEF: void
vst4_f32 (float32_t * a, float32x2x4_t val)
{
  builtin_aarch64_simd_xi o;
  float32x4x4_t temp;
  temp.val[0] = vcombine_f32 (val.val[0], vcreate_f32 (((uint64_t) 0)));
  temp.val[1] = vcombine_f32 (val.val[1], vcreate_f32 (((uint64_t) 0)));
  temp.val[2] = vcombine_f32 (val.val[2], vcreate_f32 (((uint64_t) 0)));
  temp.val[3] = vcombine_f32 (val.val[3], vcreate_f32 (((uint64_t) 0)));
  o = builtin_aarch64_set_qregxiv4sf (o, (float32x4_t) temp.val[0], 0);
  o = builtin_aarch64_set_qregxiv4sf (o, (float32x4_t) temp.val[1], 1);
  o = builtin_aarch64_set_qregxiv4sf (o, (float32x4_t) temp.val[2], 2);
  o = builtin_aarch64_set_qregxiv4sf (o, (float32x4_t) temp.val[3], 3);
  builtin_aarch64_st4v2sf ((builtin_aarch64_simd_sf *) a, o);
}

DEF: void
vst4q_s8 (int8_t * a, int8x16x4_t val)
{
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) val.val[0], 0);
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) val.val[1], 1);
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) val.val[2], 2);
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) val.val[3], 3);
  builtin_aarch64_st4v16qi ((builtin_aarch64_simd_qi *) a, o);
}

DEF: void
vst4q_p8 (poly8_t * a, poly8x16x4_t val)
{
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) val.val[0], 0);
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) val.val[1], 1);
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) val.val[2], 2);
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) val.val[3], 3);
  builtin_aarch64_st4v16qi ((builtin_aarch64_simd_qi *) a, o);
}

DEF: void
vst4q_s16 (int16_t * a, int16x8x4_t val)
{
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) val.val[0], 0);
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) val.val[1], 1);
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) val.val[2], 2);
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) val.val[3], 3);
  builtin_aarch64_st4v8hi ((builtin_aarch64_simd_hi *) a, o);
}

DEF: void
vst4q_p16 (poly16_t * a, poly16x8x4_t val)
{
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) val.val[0], 0);
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) val.val[1], 1);
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) val.val[2], 2);
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) val.val[3], 3);
  builtin_aarch64_st4v8hi ((builtin_aarch64_simd_hi *) a, o);
}

DEF: void
vst4q_s32 (int32_t * a, int32x4x4_t val)
{
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) val.val[0], 0);
  o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) val.val[1], 1);
  o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) val.val[2], 2);
  o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) val.val[3], 3);
  builtin_aarch64_st4v4si ((builtin_aarch64_simd_si *) a, o);
}

DEF: void
vst4q_s64 (int64_t * a, int64x2x4_t val)
{
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) val.val[0], 0);
  o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) val.val[1], 1);
  o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) val.val[2], 2);
  o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) val.val[3], 3);
  builtin_aarch64_st4v2di ((builtin_aarch64_simd_di *) a, o);
}

DEF: void
vst4q_u8 (uint8_t * a, uint8x16x4_t val)
{
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) val.val[0], 0);
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) val.val[1], 1);
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) val.val[2], 2);
  o = builtin_aarch64_set_qregxiv16qi (o, (int8x16_t) val.val[3], 3);
  builtin_aarch64_st4v16qi ((builtin_aarch64_simd_qi *) a, o);
}

DEF: void
vst4q_u16 (uint16_t * a, uint16x8x4_t val)
{
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) val.val[0], 0);
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) val.val[1], 1);
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) val.val[2], 2);
  o = builtin_aarch64_set_qregxiv8hi (o, (int16x8_t) val.val[3], 3);
  builtin_aarch64_st4v8hi ((builtin_aarch64_simd_hi *) a, o);
}

DEF: void
vst4q_u32 (uint32_t * a, uint32x4x4_t val)
{
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) val.val[0], 0);
  o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) val.val[1], 1);
  o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) val.val[2], 2);
  o = builtin_aarch64_set_qregxiv4si (o, (int32x4_t) val.val[3], 3);
  builtin_aarch64_st4v4si ((builtin_aarch64_simd_si *) a, o);
}

DEF: void
vst4q_u64 (uint64_t * a, uint64x2x4_t val)
{
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) val.val[0], 0);
  o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) val.val[1], 1);
  o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) val.val[2], 2);
  o = builtin_aarch64_set_qregxiv2di (o, (int64x2_t) val.val[3], 3);
  builtin_aarch64_st4v2di ((builtin_aarch64_simd_di *) a, o);
}

DEF: void
vst4q_f32 (float32_t * a, float32x4x4_t val)
{
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_set_qregxiv4sf (o, (float32x4_t) val.val[0], 0);
  o = builtin_aarch64_set_qregxiv4sf (o, (float32x4_t) val.val[1], 1);
  o = builtin_aarch64_set_qregxiv4sf (o, (float32x4_t) val.val[2], 2);
  o = builtin_aarch64_set_qregxiv4sf (o, (float32x4_t) val.val[3], 3);
  builtin_aarch64_st4v4sf ((builtin_aarch64_simd_sf *) a, o);
}

DEF: void
vst4q_f64 (float64_t * a, float64x2x4_t val)
{
  builtin_aarch64_simd_xi o;
  o = builtin_aarch64_set_qregxiv2df (o, (float64x2_t) val.val[0], 0);
  o = builtin_aarch64_set_qregxiv2df (o, (float64x2_t) val.val[1], 1);
  o = builtin_aarch64_set_qregxiv2df (o, (float64x2_t) val.val[2], 2);
  o = builtin_aarch64_set_qregxiv2df (o, (float64x2_t) val.val[3], 3);
  builtin_aarch64_st4v2df ((builtin_aarch64_simd_df *) a, o);
}

DEF: int64
vsubd_s64 (int64_t a, int64_t b)
{
  return a - b;
}

DEF: uint64
vsubd_u64 (uint64_t a, uint64_t b)
{
  return a - b;
}

DEF: int8x8
vtbx1_s8 (int8x8_t r, int8x8_t tab, int8x8_t idx)
{
  uint8x8_t mask = vclt_u8 (vreinterpret_u8_s8 (idx),
         vmov_n_u8 (8));
  int8x8_t tbl = vtbl1_s8 (tab, idx);

  return vbsl_s8 (mask, tbl, r);
}

DEF: uint8x8
vtbx1_u8 (uint8x8_t r, uint8x8_t tab, uint8x8_t idx)
{
  uint8x8_t mask = vclt_u8 (idx, vmov_n_u8 (8));
  uint8x8_t tbl = vtbl1_u8 (tab, idx);

  return vbsl_u8 (mask, tbl, r);
}

DEF: poly8x8
vtbx1_p8 (poly8x8_t r, poly8x8_t tab, uint8x8_t idx)
{
  uint8x8_t mask = vclt_u8 (idx, vmov_n_u8 (8));
  poly8x8_t tbl = vtbl1_p8 (tab, idx);

  return vbsl_p8 (mask, tbl, r);
}

DEF: int8x8
vtbx3_s8 (int8x8_t r, int8x8x3_t tab, int8x8_t idx)
{
  uint8x8_t mask = vclt_u8 (vreinterpret_u8_s8 (idx),
         vmov_n_u8 (24));
  int8x8_t tbl = vtbl3_s8 (tab, idx);

  return vbsl_s8 (mask, tbl, r);
}

DEF: uint8x8
vtbx3_u8 (uint8x8_t r, uint8x8x3_t tab, uint8x8_t idx)
{
  uint8x8_t mask = vclt_u8 (idx, vmov_n_u8 (24));
  uint8x8_t tbl = vtbl3_u8 (tab, idx);

  return vbsl_u8 (mask, tbl, r);
}

DEF: poly8x8
vtbx3_p8 (poly8x8_t r, poly8x8x3_t tab, uint8x8_t idx)
{
  uint8x8_t mask = vclt_u8 (idx, vmov_n_u8 (24));
  poly8x8_t tbl = vtbl3_p8 (tab, idx);

  return vbsl_p8 (mask, tbl, r);
}

DEF: float32x2
vtrn1_f32 (float32x2_t a, float32x2_t b)
{

  return builtin_shuffle (a, b, (uint32x2_t) {0, 2});

}

DEF: poly8x8
vtrn1_p8 (poly8x8_t a, poly8x8_t b)
{

  return builtin_shuffle (a, b, (uint8x8_t) {0, 8, 2, 10, 4, 12, 6, 14});

}

DEF: poly16x4
vtrn1_p16 (poly16x4_t a, poly16x4_t b)
{

  return builtin_shuffle (a, b, (uint16x4_t) {0, 4, 2, 6});

}

DEF: int8x8
vtrn1_s8 (int8x8_t a, int8x8_t b)
{

  return builtin_shuffle (a, b, (uint8x8_t) {0, 8, 2, 10, 4, 12, 6, 14});

}

DEF: int16x4
vtrn1_s16 (int16x4_t a, int16x4_t b)
{

  return builtin_shuffle (a, b, (uint16x4_t) {0, 4, 2, 6});

}

DEF: int32x2
vtrn1_s32 (int32x2_t a, int32x2_t b)
{

  return builtin_shuffle (a, b, (uint32x2_t) {0, 2});

}

DEF: uint8x8
vtrn1_u8 (uint8x8_t a, uint8x8_t b)
{

  return builtin_shuffle (a, b, (uint8x8_t) {0, 8, 2, 10, 4, 12, 6, 14});

}

DEF: uint16x4
vtrn1_u16 (uint16x4_t a, uint16x4_t b)
{

  return builtin_shuffle (a, b, (uint16x4_t) {0, 4, 2, 6});

}

DEF: uint32x2
vtrn1_u32 (uint32x2_t a, uint32x2_t b)
{

  return builtin_shuffle (a, b, (uint32x2_t) {0, 2});

}

DEF: float32x4
vtrn1q_f32 (float32x4_t a, float32x4_t b)
{

  return builtin_shuffle (a, b, (uint32x4_t) {0, 4, 2, 6});

}

DEF: float64x2
vtrn1q_f64 (float64x2_t a, float64x2_t b)
{

  return builtin_shuffle (a, b, (uint64x2_t) {0, 2});

}

DEF: poly8x16
vtrn1q_p8 (poly8x16_t a, poly8x16_t b)
{


  return builtin_shuffle (a, b,
      (uint8x16_t) {0, 16, 2, 18, 4, 20, 6, 22, 8, 24, 10, 26, 12, 28, 14, 30});

}

DEF: poly16x8
vtrn1q_p16 (poly16x8_t a, poly16x8_t b)
{

  return builtin_shuffle (a, b, (uint16x8_t) {0, 8, 2, 10, 4, 12, 6, 14});

}

DEF: int8x16
vtrn1q_s8 (int8x16_t a, int8x16_t b)
{


  return builtin_shuffle (a, b,
      (uint8x16_t) {0, 16, 2, 18, 4, 20, 6, 22, 8, 24, 10, 26, 12, 28, 14, 30});

}

DEF: int16x8
vtrn1q_s16 (int16x8_t a, int16x8_t b)
{

  return builtin_shuffle (a, b, (uint16x8_t) {0, 8, 2, 10, 4, 12, 6, 14});

}

DEF: int32x4
vtrn1q_s32 (int32x4_t a, int32x4_t b)
{

  return builtin_shuffle (a, b, (uint32x4_t) {0, 4, 2, 6});

}

DEF: int64x2
vtrn1q_s64 (int64x2_t a, int64x2_t b)
{

  return builtin_shuffle (a, b, (uint64x2_t) {0, 2});

}

DEF: uint8x16
vtrn1q_u8 (uint8x16_t a, uint8x16_t b)
{


  return builtin_shuffle (a, b,
      (uint8x16_t) {0, 16, 2, 18, 4, 20, 6, 22, 8, 24, 10, 26, 12, 28, 14, 30});

}

DEF: uint16x8
vtrn1q_u16 (uint16x8_t a, uint16x8_t b)
{

  return builtin_shuffle (a, b, (uint16x8_t) {0, 8, 2, 10, 4, 12, 6, 14});

}

DEF: uint32x4
vtrn1q_u32 (uint32x4_t a, uint32x4_t b)
{

  return builtin_shuffle (a, b, (uint32x4_t) {0, 4, 2, 6});

}

DEF: uint64x2
vtrn1q_u64 (uint64x2_t a, uint64x2_t b)
{

  return builtin_shuffle (a, b, (uint64x2_t) {0, 2});

}

DEF: float32x2
vtrn2_f32 (float32x2_t a, float32x2_t b)
{

  return builtin_shuffle (a, b, (uint32x2_t) {1, 3});

}

DEF: poly8x8
vtrn2_p8 (poly8x8_t a, poly8x8_t b)
{

  return builtin_shuffle (a, b, (uint8x8_t) {1, 9, 3, 11, 5, 13, 7, 15});

}

DEF: poly16x4
vtrn2_p16 (poly16x4_t a, poly16x4_t b)
{

  return builtin_shuffle (a, b, (uint16x4_t) {1, 5, 3, 7});

}

DEF: int8x8
vtrn2_s8 (int8x8_t a, int8x8_t b)
{

  return builtin_shuffle (a, b, (uint8x8_t) {1, 9, 3, 11, 5, 13, 7, 15});

}

DEF: int16x4
vtrn2_s16 (int16x4_t a, int16x4_t b)
{

  return builtin_shuffle (a, b, (uint16x4_t) {1, 5, 3, 7});

}

DEF: int32x2
vtrn2_s32 (int32x2_t a, int32x2_t b)
{

  return builtin_shuffle (a, b, (uint32x2_t) {1, 3});

}

DEF: uint8x8
vtrn2_u8 (uint8x8_t a, uint8x8_t b)
{

  return builtin_shuffle (a, b, (uint8x8_t) {1, 9, 3, 11, 5, 13, 7, 15});

}

DEF: uint16x4
vtrn2_u16 (uint16x4_t a, uint16x4_t b)
{

  return builtin_shuffle (a, b, (uint16x4_t) {1, 5, 3, 7});

}

DEF: uint32x2
vtrn2_u32 (uint32x2_t a, uint32x2_t b)
{

  return builtin_shuffle (a, b, (uint32x2_t) {1, 3});

}

DEF: float32x4
vtrn2q_f32 (float32x4_t a, float32x4_t b)
{

  return builtin_shuffle (a, b, (uint32x4_t) {1, 5, 3, 7});

}

DEF: float64x2
vtrn2q_f64 (float64x2_t a, float64x2_t b)
{

  return builtin_shuffle (a, b, (uint64x2_t) {1, 3});

}

DEF: poly8x16
vtrn2q_p8 (poly8x16_t a, poly8x16_t b)
{


  return builtin_shuffle (a, b,
      (uint8x16_t) {1, 17, 3, 19, 5, 21, 7, 23, 9, 25, 11, 27, 13, 29, 15, 31});

}

DEF: poly16x8
vtrn2q_p16 (poly16x8_t a, poly16x8_t b)
{

  return builtin_shuffle (a, b, (uint16x8_t) {1, 9, 3, 11, 5, 13, 7, 15});

}

DEF: int8x16
vtrn2q_s8 (int8x16_t a, int8x16_t b)
{


  return builtin_shuffle (a, b,
      (uint8x16_t) {1, 17, 3, 19, 5, 21, 7, 23, 9, 25, 11, 27, 13, 29, 15, 31});

}

DEF: int16x8
vtrn2q_s16 (int16x8_t a, int16x8_t b)
{

  return builtin_shuffle (a, b, (uint16x8_t) {1, 9, 3, 11, 5, 13, 7, 15});

}

DEF: int32x4
vtrn2q_s32 (int32x4_t a, int32x4_t b)
{

  return builtin_shuffle (a, b, (uint32x4_t) {1, 5, 3, 7});

}

DEF: int64x2
vtrn2q_s64 (int64x2_t a, int64x2_t b)
{

  return builtin_shuffle (a, b, (uint64x2_t) {1, 3});

}

DEF: uint8x16
vtrn2q_u8 (uint8x16_t a, uint8x16_t b)
{


  return builtin_shuffle (a, b,
      (uint8x16_t) {1, 17, 3, 19, 5, 21, 7, 23, 9, 25, 11, 27, 13, 29, 15, 31});

}

DEF: uint16x8
vtrn2q_u16 (uint16x8_t a, uint16x8_t b)
{

  return builtin_shuffle (a, b, (uint16x8_t) {1, 9, 3, 11, 5, 13, 7, 15});

}

DEF: uint32x4
vtrn2q_u32 (uint32x4_t a, uint32x4_t b)
{

  return builtin_shuffle (a, b, (uint32x4_t) {1, 5, 3, 7});

}

DEF: uint64x2
vtrn2q_u64 (uint64x2_t a, uint64x2_t b)
{

  return builtin_shuffle (a, b, (uint64x2_t) {1, 3});

}

DEF: float32x2x2
vtrn_f32 (float32x2_t a, float32x2_t b)
{
  return (float32x2x2_t) {vtrn1_f32 (a, b), vtrn2_f32 (a, b)};
}

DEF: poly8x8x2
vtrn_p8 (poly8x8_t a, poly8x8_t b)
{
  return (poly8x8x2_t) {vtrn1_p8 (a, b), vtrn2_p8 (a, b)};
}

DEF: poly16x4x2
vtrn_p16 (poly16x4_t a, poly16x4_t b)
{
  return (poly16x4x2_t) {vtrn1_p16 (a, b), vtrn2_p16 (a, b)};
}

DEF: int8x8x2
vtrn_s8 (int8x8_t a, int8x8_t b)
{
  return (int8x8x2_t) {vtrn1_s8 (a, b), vtrn2_s8 (a, b)};
}

DEF: int16x4x2
vtrn_s16 (int16x4_t a, int16x4_t b)
{
  return (int16x4x2_t) {vtrn1_s16 (a, b), vtrn2_s16 (a, b)};
}

DEF: int32x2x2
vtrn_s32 (int32x2_t a, int32x2_t b)
{
  return (int32x2x2_t) {vtrn1_s32 (a, b), vtrn2_s32 (a, b)};
}

DEF: uint8x8x2
vtrn_u8 (uint8x8_t a, uint8x8_t b)
{
  return (uint8x8x2_t) {vtrn1_u8 (a, b), vtrn2_u8 (a, b)};
}

DEF: uint16x4x2
vtrn_u16 (uint16x4_t a, uint16x4_t b)
{
  return (uint16x4x2_t) {vtrn1_u16 (a, b), vtrn2_u16 (a, b)};
}

DEF: uint32x2x2
vtrn_u32 (uint32x2_t a, uint32x2_t b)
{
  return (uint32x2x2_t) {vtrn1_u32 (a, b), vtrn2_u32 (a, b)};
}

DEF: float32x4x2
vtrnq_f32 (float32x4_t a, float32x4_t b)
{
  return (float32x4x2_t) {vtrn1q_f32 (a, b), vtrn2q_f32 (a, b)};
}

DEF: poly8x16x2
vtrnq_p8 (poly8x16_t a, poly8x16_t b)
{
  return (poly8x16x2_t) {vtrn1q_p8 (a, b), vtrn2q_p8 (a, b)};
}

DEF: poly16x8x2
vtrnq_p16 (poly16x8_t a, poly16x8_t b)
{
  return (poly16x8x2_t) {vtrn1q_p16 (a, b), vtrn2q_p16 (a, b)};
}

DEF: int8x16x2
vtrnq_s8 (int8x16_t a, int8x16_t b)
{
  return (int8x16x2_t) {vtrn1q_s8 (a, b), vtrn2q_s8 (a, b)};
}

DEF: int16x8x2
vtrnq_s16 (int16x8_t a, int16x8_t b)
{
  return (int16x8x2_t) {vtrn1q_s16 (a, b), vtrn2q_s16 (a, b)};
}

DEF: int32x4x2
vtrnq_s32 (int32x4_t a, int32x4_t b)
{
  return (int32x4x2_t) {vtrn1q_s32 (a, b), vtrn2q_s32 (a, b)};
}

DEF: uint8x16x2
vtrnq_u8 (uint8x16_t a, uint8x16_t b)
{
  return (uint8x16x2_t) {vtrn1q_u8 (a, b), vtrn2q_u8 (a, b)};
}

DEF: uint16x8x2
vtrnq_u16 (uint16x8_t a, uint16x8_t b)
{
  return (uint16x8x2_t) {vtrn1q_u16 (a, b), vtrn2q_u16 (a, b)};
}

DEF: uint32x4x2
vtrnq_u32 (uint32x4_t a, uint32x4_t b)
{
  return (uint32x4x2_t) {vtrn1q_u32 (a, b), vtrn2q_u32 (a, b)};
}

DEF: uint8x8
vtst_s8 (int8x8_t a, int8x8_t b)
{
  return (uint8x8_t) ((a & b) != 0);
}

DEF: uint16x4
vtst_s16 (int16x4_t a, int16x4_t b)
{
  return (uint16x4_t) ((a & b) != 0);
}

DEF: uint32x2
vtst_s32 (int32x2_t a, int32x2_t b)
{
  return (uint32x2_t) ((a & b) != 0);
}

DEF: uint64x1
vtst_s64 (int64x1_t a, int64x1_t b)
{
  return (uint64x1_t) ((a & b) != ((int64_t) 0));
}

DEF: uint8x8
vtst_u8 (uint8x8_t a, uint8x8_t b)
{
  return ((a & b) != 0);
}

DEF: uint16x4
vtst_u16 (uint16x4_t a, uint16x4_t b)
{
  return ((a & b) != 0);
}

DEF: uint32x2
vtst_u32 (uint32x2_t a, uint32x2_t b)
{
  return ((a & b) != 0);
}

DEF: uint64x1
vtst_u64 (uint64x1_t a, uint64x1_t b)
{
  return ((a & b) != ((uint64_t) 0));
}

DEF: uint8x16
vtstq_s8 (int8x16_t a, int8x16_t b)
{
  return (uint8x16_t) ((a & b) != 0);
}

DEF: uint16x8
vtstq_s16 (int16x8_t a, int16x8_t b)
{
  return (uint16x8_t) ((a & b) != 0);
}

DEF: uint32x4
vtstq_s32 (int32x4_t a, int32x4_t b)
{
  return (uint32x4_t) ((a & b) != 0);
}

DEF: uint64x2
vtstq_s64 (int64x2_t a, int64x2_t b)
{
  return (uint64x2_t) ((a & b) != ((int64_t) 0));
}

DEF: uint8x16
vtstq_u8 (uint8x16_t a, uint8x16_t b)
{
  return ((a & b) != 0);
}

DEF: uint16x8
vtstq_u16 (uint16x8_t a, uint16x8_t b)
{
  return ((a & b) != 0);
}

DEF: uint32x4
vtstq_u32 (uint32x4_t a, uint32x4_t b)
{
  return ((a & b) != 0);
}

DEF: uint64x2
vtstq_u64 (uint64x2_t a, uint64x2_t b)
{
  return ((a & b) != ((uint64_t) 0));
}

DEF: uint64
vtstd_s64 (int64_t a, int64_t b)
{
  return (a & b) ? -1ll : 0ll;
}

DEF: uint64
vtstd_u64 (uint64_t a, uint64_t b)
{
  return (a & b) ? -1ll : 0ll;
}

DEF: int8x8
vuqadd_s8 (int8x8_t a, uint8x8_t b)
{
  return builtin_aarch64_suqaddv8qi_ssu (a, b);
}

DEF: int16x4
vuqadd_s16 (int16x4_t a, uint16x4_t b)
{
  return builtin_aarch64_suqaddv4hi_ssu (a, b);
}

DEF: int32x2
vuqadd_s32 (int32x2_t a, uint32x2_t b)
{
  return builtin_aarch64_suqaddv2si_ssu (a, b);
}

DEF: int64x1
vuqadd_s64 (int64x1_t a, uint64x1_t b)
{
  return (int64x1_t) {builtin_aarch64_suqadddi_ssu (a[0], b[0])};
}

DEF: int8x16
vuqaddq_s8 (int8x16_t a, uint8x16_t b)
{
  return builtin_aarch64_suqaddv16qi_ssu (a, b);
}

DEF: int16x8
vuqaddq_s16 (int16x8_t a, uint16x8_t b)
{
  return builtin_aarch64_suqaddv8hi_ssu (a, b);
}

DEF: int32x4
vuqaddq_s32 (int32x4_t a, uint32x4_t b)
{
  return builtin_aarch64_suqaddv4si_ssu (a, b);
}

DEF: int64x2
vuqaddq_s64 (int64x2_t a, uint64x2_t b)
{
  return builtin_aarch64_suqaddv2di_ssu (a, b);
}

DEF: int8
vuqaddb_s8 (int8_t a, uint8_t b)
{
  return builtin_aarch64_suqaddqi_ssu (a, b);
}

DEF: int16
vuqaddh_s16 (int16_t a, uint16_t b)
{
  return builtin_aarch64_suqaddhi_ssu (a, b);
}

DEF: int32
vuqadds_s32 (int32_t a, uint32_t b)
{
  return builtin_aarch64_suqaddsi_ssu (a, b);
}

DEF: int64
vuqaddd_s64 (int64_t a, uint64_t b)
{
  return builtin_aarch64_suqadddi_ssu (a, b);
}
# 23981 "arm_neon.h"
DEF: float32x2
vuzp1_f32 (float32x2_t a, float32x2_t b)
{

  return builtin_shuffle (a, b, (uint32x2_t) {0, 2});

}

DEF: poly8x8
vuzp1_p8 (poly8x8_t a, poly8x8_t b)
{

  return builtin_shuffle (a, b, (uint8x8_t) {0, 2, 4, 6, 8, 10, 12, 14});

}

DEF: poly16x4
vuzp1_p16 (poly16x4_t a, poly16x4_t b)
{

  return builtin_shuffle (a, b, (uint16x4_t) {0, 2, 4, 6});

}

DEF: int8x8
vuzp1_s8 (int8x8_t a, int8x8_t b)
{

  return builtin_shuffle (a, b, (uint8x8_t) {0, 2, 4, 6, 8, 10, 12, 14});

}

DEF: int16x4
vuzp1_s16 (int16x4_t a, int16x4_t b)
{

  return builtin_shuffle (a, b, (uint16x4_t) {0, 2, 4, 6});

}

DEF: int32x2
vuzp1_s32 (int32x2_t a, int32x2_t b)
{

  return builtin_shuffle (a, b, (uint32x2_t) {0, 2});

}

DEF: uint8x8
vuzp1_u8 (uint8x8_t a, uint8x8_t b)
{

  return builtin_shuffle (a, b, (uint8x8_t) {0, 2, 4, 6, 8, 10, 12, 14});

}

DEF: uint16x4
vuzp1_u16 (uint16x4_t a, uint16x4_t b)
{

  return builtin_shuffle (a, b, (uint16x4_t) {0, 2, 4, 6});

}

DEF: uint32x2
vuzp1_u32 (uint32x2_t a, uint32x2_t b)
{

  return builtin_shuffle (a, b, (uint32x2_t) {0, 2});

}

DEF: float32x4
vuzp1q_f32 (float32x4_t a, float32x4_t b)
{

  return builtin_shuffle (a, b, (uint32x4_t) {0, 2, 4, 6});

}

DEF: float64x2
vuzp1q_f64 (float64x2_t a, float64x2_t b)
{

  return builtin_shuffle (a, b, (uint64x2_t) {0, 2});

}

DEF: poly8x16
vuzp1q_p8 (poly8x16_t a, poly8x16_t b)
{


  return builtin_shuffle (a, b, (uint8x16_t)
      {0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30});

}

DEF: poly16x8
vuzp1q_p16 (poly16x8_t a, poly16x8_t b)
{

  return builtin_shuffle (a, b, (uint16x8_t) {0, 2, 4, 6, 8, 10, 12, 14});

}

DEF: int8x16
vuzp1q_s8 (int8x16_t a, int8x16_t b)
{


  return builtin_shuffle (a, b,
      (uint8x16_t) {0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30});

}

DEF: int16x8
vuzp1q_s16 (int16x8_t a, int16x8_t b)
{

  return builtin_shuffle (a, b, (uint16x8_t) {0, 2, 4, 6, 8, 10, 12, 14});

}

DEF: int32x4
vuzp1q_s32 (int32x4_t a, int32x4_t b)
{

  return builtin_shuffle (a, b, (uint32x4_t) {0, 2, 4, 6});

}

DEF: int64x2
vuzp1q_s64 (int64x2_t a, int64x2_t b)
{

  return builtin_shuffle (a, b, (uint64x2_t) {0, 2});

}

DEF: uint8x16
vuzp1q_u8 (uint8x16_t a, uint8x16_t b)
{


  return builtin_shuffle (a, b,
      (uint8x16_t) {0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30});

}

DEF: uint16x8
vuzp1q_u16 (uint16x8_t a, uint16x8_t b)
{

  return builtin_shuffle (a, b, (uint16x8_t) {0, 2, 4, 6, 8, 10, 12, 14});

}

DEF: uint32x4
vuzp1q_u32 (uint32x4_t a, uint32x4_t b)
{

  return builtin_shuffle (a, b, (uint32x4_t) {0, 2, 4, 6});

}

DEF: uint64x2
vuzp1q_u64 (uint64x2_t a, uint64x2_t b)
{

  return builtin_shuffle (a, b, (uint64x2_t) {0, 2});

}

DEF: float32x2
vuzp2_f32 (float32x2_t a, float32x2_t b)
{

  return builtin_shuffle (a, b, (uint32x2_t) {1, 3});

}

DEF: poly8x8
vuzp2_p8 (poly8x8_t a, poly8x8_t b)
{

  return builtin_shuffle (a, b, (uint8x8_t) {1, 3, 5, 7, 9, 11, 13, 15});

}

DEF: poly16x4
vuzp2_p16 (poly16x4_t a, poly16x4_t b)
{

  return builtin_shuffle (a, b, (uint16x4_t) {1, 3, 5, 7});

}

DEF: int8x8
vuzp2_s8 (int8x8_t a, int8x8_t b)
{

  return builtin_shuffle (a, b, (uint8x8_t) {1, 3, 5, 7, 9, 11, 13, 15});

}

DEF: int16x4
vuzp2_s16 (int16x4_t a, int16x4_t b)
{

  return builtin_shuffle (a, b, (uint16x4_t) {1, 3, 5, 7});

}

DEF: int32x2
vuzp2_s32 (int32x2_t a, int32x2_t b)
{

  return builtin_shuffle (a, b, (uint32x2_t) {1, 3});

}

DEF: uint8x8
vuzp2_u8 (uint8x8_t a, uint8x8_t b)
{

  return builtin_shuffle (a, b, (uint8x8_t) {1, 3, 5, 7, 9, 11, 13, 15});

}

DEF: uint16x4
vuzp2_u16 (uint16x4_t a, uint16x4_t b)
{

  return builtin_shuffle (a, b, (uint16x4_t) {1, 3, 5, 7});

}

DEF: uint32x2
vuzp2_u32 (uint32x2_t a, uint32x2_t b)
{

  return builtin_shuffle (a, b, (uint32x2_t) {1, 3});

}

DEF: float32x4
vuzp2q_f32 (float32x4_t a, float32x4_t b)
{

  return builtin_shuffle (a, b, (uint32x4_t) {1, 3, 5, 7});

}

DEF: float64x2
vuzp2q_f64 (float64x2_t a, float64x2_t b)
{

  return builtin_shuffle (a, b, (uint64x2_t) {1, 3});

}

DEF: poly8x16
vuzp2q_p8 (poly8x16_t a, poly8x16_t b)
{


  return builtin_shuffle (a, b,
      (uint8x16_t) {1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31});

}

DEF: poly16x8
vuzp2q_p16 (poly16x8_t a, poly16x8_t b)
{

  return builtin_shuffle (a, b, (uint16x8_t) {1, 3, 5, 7, 9, 11, 13, 15});

}

DEF: int8x16
vuzp2q_s8 (int8x16_t a, int8x16_t b)
{


  return builtin_shuffle (a, b,
      (uint8x16_t) {1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31});

}

DEF: int16x8
vuzp2q_s16 (int16x8_t a, int16x8_t b)
{

  return builtin_shuffle (a, b, (uint16x8_t) {1, 3, 5, 7, 9, 11, 13, 15});

}

DEF: int32x4
vuzp2q_s32 (int32x4_t a, int32x4_t b)
{

  return builtin_shuffle (a, b, (uint32x4_t) {1, 3, 5, 7});

}

DEF: int64x2
vuzp2q_s64 (int64x2_t a, int64x2_t b)
{

  return builtin_shuffle (a, b, (uint64x2_t) {1, 3});

}

DEF: uint8x16
vuzp2q_u8 (uint8x16_t a, uint8x16_t b)
{


  return builtin_shuffle (a, b, (uint8x16_t)
      {1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31});

}

DEF: uint16x8
vuzp2q_u16 (uint16x8_t a, uint16x8_t b)
{

  return builtin_shuffle (a, b, (uint16x8_t) {1, 3, 5, 7, 9, 11, 13, 15});

}

DEF: uint32x4
vuzp2q_u32 (uint32x4_t a, uint32x4_t b)
{

  return builtin_shuffle (a, b, (uint32x4_t) {1, 3, 5, 7});

}

DEF: uint64x2
vuzp2q_u64 (uint64x2_t a, uint64x2_t b)
{

  return builtin_shuffle (a, b, (uint64x2_t) {1, 3});

}

DEF: float32x2x2 vuzp_f32 (float32x2_t a, float32x2_t b) { return (float32x2x2_t) {vuzp1_f32 (a, b), vuzp2_f32 (a, b)}; }
DEF: poly8x8x2 vuzp_p8 (poly8x8_t a, poly8x8_t b) { return (poly8x8x2_t) {vuzp1_p8 (a, b), vuzp2_p8 (a, b)}; }
DEF: poly16x4x2 vuzp_p16 (poly16x4_t a, poly16x4_t b) { return (poly16x4x2_t) {vuzp1_p16 (a, b), vuzp2_p16 (a, b)}; }
DEF: int8x8x2 vuzp_s8 (int8x8_t a, int8x8_t b) { return (int8x8x2_t) {vuzp1_s8 (a, b), vuzp2_s8 (a, b)}; }
DEF: int16x4x2 vuzp_s16 (int16x4_t a, int16x4_t b) { return (int16x4x2_t) {vuzp1_s16 (a, b), vuzp2_s16 (a, b)}; }
DEF: int32x2x2 vuzp_s32 (int32x2_t a, int32x2_t b) { return (int32x2x2_t) {vuzp1_s32 (a, b), vuzp2_s32 (a, b)}; }
DEF: uint8x8x2 vuzp_u8 (uint8x8_t a, uint8x8_t b) { return (uint8x8x2_t) {vuzp1_u8 (a, b), vuzp2_u8 (a, b)}; }
DEF: uint16x4x2 vuzp_u16 (uint16x4_t a, uint16x4_t b) { return (uint16x4x2_t) {vuzp1_u16 (a, b), vuzp2_u16 (a, b)}; }
DEF: uint32x2x2 vuzp_u32 (uint32x2_t a, uint32x2_t b) { return (uint32x2x2_t) {vuzp1_u32 (a, b), vuzp2_u32 (a, b)}; }
DEF: float32x4x2 vuzpq_f32 (float32x4_t a, float32x4_t b) { return (float32x4x2_t) {vuzp1q_f32 (a, b), vuzp2q_f32 (a, b)}; }
DEF: poly8x16x2 vuzpq_p8 (poly8x16_t a, poly8x16_t b) { return (poly8x16x2_t) {vuzp1q_p8 (a, b), vuzp2q_p8 (a, b)}; }
DEF: poly16x8x2 vuzpq_p16 (poly16x8_t a, poly16x8_t b) { return (poly16x8x2_t) {vuzp1q_p16 (a, b), vuzp2q_p16 (a, b)}; }
DEF: int8x16x2 vuzpq_s8 (int8x16_t a, int8x16_t b) { return (int8x16x2_t) {vuzp1q_s8 (a, b), vuzp2q_s8 (a, b)}; }
DEF: int16x8x2 vuzpq_s16 (int16x8_t a, int16x8_t b) { return (int16x8x2_t) {vuzp1q_s16 (a, b), vuzp2q_s16 (a, b)}; }
DEF: int32x4x2 vuzpq_s32 (int32x4_t a, int32x4_t b) { return (int32x4x2_t) {vuzp1q_s32 (a, b), vuzp2q_s32 (a, b)}; }
DEF: uint8x16x2 vuzpq_u8 (uint8x16_t a, uint8x16_t b) { return (uint8x16x2_t) {vuzp1q_u8 (a, b), vuzp2q_u8 (a, b)}; }
DEF: uint16x8x2 vuzpq_u16 (uint16x8_t a, uint16x8_t b) { return (uint16x8x2_t) {vuzp1q_u16 (a, b), vuzp2q_u16 (a, b)}; }
DEF: uint32x4x2 vuzpq_u32 (uint32x4_t a, uint32x4_t b) { return (uint32x4x2_t) {vuzp1q_u32 (a, b), vuzp2q_u32 (a, b)}; }

DEF: float32x2
vzip1_f32 (float32x2_t a, float32x2_t b)
{

  return builtin_shuffle (a, b, (uint32x2_t) {0, 2});

}

DEF: poly8x8
vzip1_p8 (poly8x8_t a, poly8x8_t b)
{

  return builtin_shuffle (a, b, (uint8x8_t) {0, 8, 1, 9, 2, 10, 3, 11});

}

DEF: poly16x4
vzip1_p16 (poly16x4_t a, poly16x4_t b)
{

  return builtin_shuffle (a, b, (uint16x4_t) {0, 4, 1, 5});

}

DEF: int8x8
vzip1_s8 (int8x8_t a, int8x8_t b)
{

  return builtin_shuffle (a, b, (uint8x8_t) {0, 8, 1, 9, 2, 10, 3, 11});

}

DEF: int16x4
vzip1_s16 (int16x4_t a, int16x4_t b)
{

  return builtin_shuffle (a, b, (uint16x4_t) {0, 4, 1, 5});

}

DEF: int32x2
vzip1_s32 (int32x2_t a, int32x2_t b)
{

  return builtin_shuffle (a, b, (uint32x2_t) {0, 2});

}

DEF: uint8x8
vzip1_u8 (uint8x8_t a, uint8x8_t b)
{

  return builtin_shuffle (a, b, (uint8x8_t) {0, 8, 1, 9, 2, 10, 3, 11});

}

DEF: uint16x4
vzip1_u16 (uint16x4_t a, uint16x4_t b)
{

  return builtin_shuffle (a, b, (uint16x4_t) {0, 4, 1, 5});

}

DEF: uint32x2
vzip1_u32 (uint32x2_t a, uint32x2_t b)
{

  return builtin_shuffle (a, b, (uint32x2_t) {0, 2});

}

DEF: float32x4
vzip1q_f32 (float32x4_t a, float32x4_t b)
{

  return builtin_shuffle (a, b, (uint32x4_t) {0, 4, 1, 5});

}

DEF: float64x2
vzip1q_f64 (float64x2_t a, float64x2_t b)
{

  return builtin_shuffle (a, b, (uint64x2_t) {0, 2});

}

DEF: poly8x16
vzip1q_p8 (poly8x16_t a, poly8x16_t b)
{


  return builtin_shuffle (a, b, (uint8x16_t)
      {0, 16, 1, 17, 2, 18, 3, 19, 4, 20, 5, 21, 6, 22, 7, 23});

}

DEF: poly16x8
vzip1q_p16 (poly16x8_t a, poly16x8_t b)
{


  return builtin_shuffle (a, b, (uint16x8_t) {0, 8, 1, 9, 2, 10, 3, 11});

}

DEF: int8x16
vzip1q_s8 (int8x16_t a, int8x16_t b)
{


  return builtin_shuffle (a, b, (uint8x16_t)
      {0, 16, 1, 17, 2, 18, 3, 19, 4, 20, 5, 21, 6, 22, 7, 23});

}

DEF: int16x8
vzip1q_s16 (int16x8_t a, int16x8_t b)
{


  return builtin_shuffle (a, b, (uint16x8_t) {0, 8, 1, 9, 2, 10, 3, 11});

}

DEF: int32x4
vzip1q_s32 (int32x4_t a, int32x4_t b)
{

  return builtin_shuffle (a, b, (uint32x4_t) {0, 4, 1, 5});

}

DEF: int64x2
vzip1q_s64 (int64x2_t a, int64x2_t b)
{

  return builtin_shuffle (a, b, (uint64x2_t) {0, 2});

}

DEF: uint8x16
vzip1q_u8 (uint8x16_t a, uint8x16_t b)
{


  return builtin_shuffle (a, b, (uint8x16_t)
      {0, 16, 1, 17, 2, 18, 3, 19, 4, 20, 5, 21, 6, 22, 7, 23});

}

DEF: uint16x8
vzip1q_u16 (uint16x8_t a, uint16x8_t b)
{


  return builtin_shuffle (a, b, (uint16x8_t) {0, 8, 1, 9, 2, 10, 3, 11});

}

DEF: uint32x4
vzip1q_u32 (uint32x4_t a, uint32x4_t b)
{

  return builtin_shuffle (a, b, (uint32x4_t) {0, 4, 1, 5});

}

DEF: uint64x2
vzip1q_u64 (uint64x2_t a, uint64x2_t b)
{

  return builtin_shuffle (a, b, (uint64x2_t) {0, 2});

}

DEF: float32x2
vzip2_f32 (float32x2_t a, float32x2_t b)
{

  return builtin_shuffle (a, b, (uint32x2_t) {1, 3});

}

DEF: poly8x8
vzip2_p8 (poly8x8_t a, poly8x8_t b)
{

  return builtin_shuffle (a, b, (uint8x8_t) {4, 12, 5, 13, 6, 14, 7, 15});

}

DEF: poly16x4
vzip2_p16 (poly16x4_t a, poly16x4_t b)
{

  return builtin_shuffle (a, b, (uint16x4_t) {2, 6, 3, 7});

}

DEF: int8x8
vzip2_s8 (int8x8_t a, int8x8_t b)
{

  return builtin_shuffle (a, b, (uint8x8_t) {4, 12, 5, 13, 6, 14, 7, 15});

}

DEF: int16x4
vzip2_s16 (int16x4_t a, int16x4_t b)
{

  return builtin_shuffle (a, b, (uint16x4_t) {2, 6, 3, 7});

}

DEF: int32x2
vzip2_s32 (int32x2_t a, int32x2_t b)
{

  return builtin_shuffle (a, b, (uint32x2_t) {1, 3});

}

DEF: uint8x8
vzip2_u8 (uint8x8_t a, uint8x8_t b)
{

  return builtin_shuffle (a, b, (uint8x8_t) {4, 12, 5, 13, 6, 14, 7, 15});

}

DEF: uint16x4
vzip2_u16 (uint16x4_t a, uint16x4_t b)
{

  return builtin_shuffle (a, b, (uint16x4_t) {2, 6, 3, 7});

}

DEF: uint32x2
vzip2_u32 (uint32x2_t a, uint32x2_t b)
{

  return builtin_shuffle (a, b, (uint32x2_t) {1, 3});

}

DEF: float32x4
vzip2q_f32 (float32x4_t a, float32x4_t b)
{

  return builtin_shuffle (a, b, (uint32x4_t) {2, 6, 3, 7});

}

DEF: float64x2
vzip2q_f64 (float64x2_t a, float64x2_t b)
{

  return builtin_shuffle (a, b, (uint64x2_t) {1, 3});

}

DEF: poly8x16
vzip2q_p8 (poly8x16_t a, poly8x16_t b)
{


  return builtin_shuffle (a, b, (uint8x16_t)
      {8, 24, 9, 25, 10, 26, 11, 27, 12, 28, 13, 29, 14, 30, 15, 31});

}

DEF: poly16x8
vzip2q_p16 (poly16x8_t a, poly16x8_t b)
{

  return builtin_shuffle (a, b, (uint16x8_t)
      {4, 12, 5, 13, 6, 14, 7, 15});

}

DEF: int8x16
vzip2q_s8 (int8x16_t a, int8x16_t b)
{


  return builtin_shuffle (a, b, (uint8x16_t)
      {8, 24, 9, 25, 10, 26, 11, 27, 12, 28, 13, 29, 14, 30, 15, 31});

}

DEF: int16x8
vzip2q_s16 (int16x8_t a, int16x8_t b)
{

  return builtin_shuffle (a, b, (uint16x8_t)
      {4, 12, 5, 13, 6, 14, 7, 15});

}

DEF: int32x4
vzip2q_s32 (int32x4_t a, int32x4_t b)
{

  return builtin_shuffle (a, b, (uint32x4_t) {2, 6, 3, 7});

}

DEF: int64x2
vzip2q_s64 (int64x2_t a, int64x2_t b)
{

  return builtin_shuffle (a, b, (uint64x2_t) {1, 3});

}

DEF: uint8x16
vzip2q_u8 (uint8x16_t a, uint8x16_t b)
{


  return builtin_shuffle (a, b, (uint8x16_t)
      {8, 24, 9, 25, 10, 26, 11, 27, 12, 28, 13, 29, 14, 30, 15, 31});

}

DEF: uint16x8
vzip2q_u16 (uint16x8_t a, uint16x8_t b)
{

  return builtin_shuffle (a, b, (uint16x8_t)
      {4, 12, 5, 13, 6, 14, 7, 15});

}

DEF: uint32x4
vzip2q_u32 (uint32x4_t a, uint32x4_t b)
{

  return builtin_shuffle (a, b, (uint32x4_t) {2, 6, 3, 7});

}

DEF: uint64x2
vzip2q_u64 (uint64x2_t a, uint64x2_t b)
{

  return builtin_shuffle (a, b, (uint64x2_t) {1, 3});

}

DEF: float32x2x2 vzip_f32 (float32x2_t a, float32x2_t b) { return (float32x2x2_t) {vzip1_f32 (a, b), vzip2_f32 (a, b)}; }
DEF: poly8x8x2 vzip_p8 (poly8x8_t a, poly8x8_t b) { return (poly8x8x2_t) {vzip1_p8 (a, b), vzip2_p8 (a, b)}; }
DEF: poly16x4x2 vzip_p16 (poly16x4_t a, poly16x4_t b) { return (poly16x4x2_t) {vzip1_p16 (a, b), vzip2_p16 (a, b)}; }
DEF: int8x8x2 vzip_s8 (int8x8_t a, int8x8_t b) { return (int8x8x2_t) {vzip1_s8 (a, b), vzip2_s8 (a, b)}; }
DEF: int16x4x2 vzip_s16 (int16x4_t a, int16x4_t b) { return (int16x4x2_t) {vzip1_s16 (a, b), vzip2_s16 (a, b)}; }
DEF: int32x2x2 vzip_s32 (int32x2_t a, int32x2_t b) { return (int32x2x2_t) {vzip1_s32 (a, b), vzip2_s32 (a, b)}; }
DEF: uint8x8x2 vzip_u8 (uint8x8_t a, uint8x8_t b) { return (uint8x8x2_t) {vzip1_u8 (a, b), vzip2_u8 (a, b)}; }
DEF: uint16x4x2 vzip_u16 (uint16x4_t a, uint16x4_t b) { return (uint16x4x2_t) {vzip1_u16 (a, b), vzip2_u16 (a, b)}; }
DEF: uint32x2x2 vzip_u32 (uint32x2_t a, uint32x2_t b) { return (uint32x2x2_t) {vzip1_u32 (a, b), vzip2_u32 (a, b)}; }
DEF: float32x4x2 vzipq_f32 (float32x4_t a, float32x4_t b) { return (float32x4x2_t) {vzip1q_f32 (a, b), vzip2q_f32 (a, b)}; }
DEF: poly8x16x2 vzipq_p8 (poly8x16_t a, poly8x16_t b) { return (poly8x16x2_t) {vzip1q_p8 (a, b), vzip2q_p8 (a, b)}; }
DEF: poly16x8x2 vzipq_p16 (poly16x8_t a, poly16x8_t b) { return (poly16x8x2_t) {vzip1q_p16 (a, b), vzip2q_p16 (a, b)}; }
DEF: int8x16x2 vzipq_s8 (int8x16_t a, int8x16_t b) { return (int8x16x2_t) {vzip1q_s8 (a, b), vzip2q_s8 (a, b)}; }
DEF: int16x8x2 vzipq_s16 (int16x8_t a, int16x8_t b) { return (int16x8x2_t) {vzip1q_s16 (a, b), vzip2q_s16 (a, b)}; }
DEF: int32x4x2 vzipq_s32 (int32x4_t a, int32x4_t b) { return (int32x4x2_t) {vzip1q_s32 (a, b), vzip2q_s32 (a, b)}; }
DEF: uint8x16x2 vzipq_u8 (uint8x16_t a, uint8x16_t b) { return (uint8x16x2_t) {vzip1q_u8 (a, b), vzip2q_u8 (a, b)}; }
DEF: uint16x8x2 vzipq_u16 (uint16x8_t a, uint16x8_t b) { return (uint16x8x2_t) {vzip1q_u16 (a, b), vzip2q_u16 (a, b)}; }
DEF: uint32x4x2 vzipq_u32 (uint32x4_t a, uint32x4_t b) { return (uint32x4x2_t) {vzip1q_u32 (a, b), vzip2q_u32 (a, b)}; }
# 24914 "arm_neon.h"
#pragma GCC pop_options
# 2 "test.c" 2
