; 前方参照（シンボル、ラベル mix）
fw_sym	equ def + xyz + 1
fw_label equ end_addr  + 2
fw_mix	equ end_addr + def + 3

	ld hl, fw_sym 
	ld hl, fw_label 
	ld hl, fw_mix 
def 	equ 2
xyz 	equ 3
end_addr:

