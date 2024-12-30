.text
.globl hoho
hoho:
	pushq %rbp
	movq %rsp, %rbp
	movl $5, %eax
	leave
	ret
.type hoho, @function
.size hoho, .-hoho
/* end function hoho */

.text
.globl main
main:
	pushq %rbp
	movq %rsp, %rbp
	callq hoho
	leave
	ret
.type main, @function
.size main, .-main
/* end function main */

.section .note.GNU-stack,"",@progbits
