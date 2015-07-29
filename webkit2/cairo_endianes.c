#include <stdint.h>
//#include <stdio.h>

#include "cairo_endianes.h"

const uint32_t gowk2_endian_test = 0x00010203;

uint8_t* gowk2_detect_endianes() {
	return (uint8_t*)&gowk2_endian_test;
}

void gowk2_cairo_endian_depended_ARGB32_to_RGBA(unsigned char *data, unsigned int len) {
	char* order;
	unsigned int i,j;
	char r, g, b, a;
	//printf("webkit endian 0\n");
	

	order = gowk2_detect_endianes();

	//save processing time and check for special case where data is already in desired order. Desired order is Go's natural order RGBA.
        //cairo ARGB32
        //go    RGBA
	if(order[0]==0x01 && order[1]==0x02 && order[2]==0x03 && order[3]==0x00)
		return;
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
		data[i+0] = r;
		data[i+1] = g;
		data[i+2] = b;
		data[i+3] = a;
	}
	//printf("webkit endian 2\n");
      return;
}

void gowk2_cairo_endian_depended_ARGB32_to_RGBA_copy(unsigned char *data, unsigned char *out, unsigned int len) {
	char* order;
	unsigned int i,j;
	char r, g, b, a;
	//printf("webkit endian 0\n");
	

	order = gowk2_detect_endianes();

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
