#include "main.h"

int main(void)
{
	uint32_t timer = getSYSTIM();
	while(1)
	{
		if((chk4TimeoutSYSTIM(timer, 100) == (SYSTIM_TIMEOUT)) && (g_USB2BCOM_INFO.usb_init_state == 1))
		{
			g_USB2BCOM_INFO.usb_init_state = 0;
			initUSB2BCOM();
		}
		chkADC();
		chkGUI();
		chkUSB2BCOM();
	}
	return 1;
}
