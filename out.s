.text
abc:
	pushq %rbp
	movq %rsp, %rbp
	movl %edi, %eax
	addl $95, %eax
	leave
	ret
.type abc, @function
.size abc, .-abc
/* end function abc */

.text
.globl main
main:
	pushq %rbp
	movq %rsp, %rbp
	movl $5, %edi
	callq abc
	leave
	ret
.type main, @function
.size main, .-main
/* end function main */

.section .note.GNU-stack,"",@progbits
