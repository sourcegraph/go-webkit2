#include <stdint.h>
//#include <stdio.h>

#include "cairo_endianes.h"

const uint32_t gowk2_endian_test = 0x00010203;

// Endianes constants
#define GOWK2_BIG_ENDIAN   	0x00010203
#define GOWK2_LITTLE_ENDIAN_8  	0x03020100
#define GOWK2_LITTLE_ENDIAN_16  0x02030001
#define GOWK2_MIDLE_ENDIAN  	0x01000302

uint8_t* gowk2_detect_endianes_byte() {
	return (uint8_t*)&gowk2_endian_test;
}

uint32_t gowk2_detect_endianes_word() {
	uint8_t *b_ptr;
	b_ptr = (uint8_t*)&gowk2_endian_test;
	return ((uint32_t)b_ptr[0] << 24) | ((uint32_t)b_ptr[1] << 16) | ((uint32_t)b_ptr[2] << 8) | ((uint32_t)b_ptr[3]);
}

void gowk2_cairo_endian_depended_ARGB32_to_RGBA(unsigned char *data, unsigned char *out, unsigned int len) {
	// s1 and s2 are swap variables for the case, when data==out
	register char sw;
	register unsigned int i;
	uint32_t order;
	//printf("webkit endian 0\n");
	

	order = gowk2_detect_endianes_word();

	//printf("webkit endian 1\n");
	switch(order) {
	case GOWK2_BIG_ENDIAN:
		// A R G B ->   R G B A
		// 0 1 2 3 ->   1 2 3 0
		for(i=0; i<len; i+=4) {
			sw = data[i+0];
			out[i+0] = data[i+1];
			out[i+1] = data[i+2];
			out[i+2] = data[i+3];
			out[i+3] = sw;
		}
		break;
	case GOWK2_LITTLE_ENDIAN_8:
		// A R G B ->   R G B A
		// 3 2 1 0 ->   2 1 0 3
		if(data == out) {
			// just swap
			// if we are swapping bytes inside same array, then we do not need to copy all data.
			// this saves time
			for(i=0; i<len; i+=4) {
				sw =data[i+0];
				out[i+0] = data[i+2];
				out[i+2] = sw;
			}
		} else {
			// also copy
			for(i=0; i<len; i+=4) {
				sw =data[i+0];
				out[i+0] = data[i+2];
				out[i+1] = data[i+1];
				out[i+2] = sw;
				out[i+3] = data[i+3];
			}
		}
		break;
	case GOWK2_LITTLE_ENDIAN_16:
		// A R G B ->   R G B A
		// 2 3 0 1 ->   3 0 1 2
		for(i=0; i<len; i+=4) {
			sw = data[i+0];
			out[i+0] = data[i+3];
			out[i+3] = data[i+2];
			out[i+2] = data[i+1];
			out[i+1] = sw;
		}
		break;
	case GOWK2_MIDLE_ENDIAN:
		// A R G B ->   R G B A
		// 1 0 3 2 ->   0 3 2 1
		if(data == out) {
			// just swap
			// if we are swapping bytes inside same array, then we do not need to copy all data.
			// this saves time
			for(i=0; i<len; i+=4) {
				sw = data[i+1];
				out[i+1] = data[i+3];
				out[i+3] = sw;
			}
		} else {
			// also copy
			for(i=0; i<len; i+=4) {
				sw = data[i+1];
				out[i+0] = data[i+0];
				out[i+1] = data[i+3];
				out[i+2] = data[i+2];
				out[i+3] = sw;
			}
		}
		break;
	default:
		gowk2_cairo_endian_depended_ARGB32_to_RGBA_smart(data, out, len);
		break;
	}
}

void gowk2_cairo_endian_depended_ARGB32_to_RGBA_smart(unsigned char *data, unsigned char *out, unsigned int len) {
	register char r, g, b, a;
	register unsigned int i,j;
	char* order;
	//printf("webkit endian 0\n");
	

	order = gowk2_detect_endianes_byte();

	//printf("webkit endian 1\n");
	for(i=0; i<len; i+=4) {
		for(j=0;j<4;j++) {
			switch(order[j]) {
			case 0x00:
				a = data[i+j];
				break;
			case 0x01:
				r = data[i+j];
				break;
			case 0x02:
				g = data[i+j];
				break;
			case 0x03:
				b = data[i+j];
				break;
			}
		}
		out[i+0] = r;
		out[i+1] = g;
		out[i+2] = b;
		out[i+3] = a;
	}
	//printf("webkit endian 2\n");
      return;
}
