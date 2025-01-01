.text
ahoj:
	pushq %rbp
	movq %rsp, %rbp
	movq %rdi, %rax
	addq $9, %rax
	leave
	ret
.type ahoj, @function
.size ahoj, .-ahoj
/* end function ahoj */

.text
.globl main
main:
	pushq %rbp
	movq %rsp, %rbp
	movl $7, %edi
	callq ahoj
	leave
	ret
.type main, @function
.size main, .-main
/* end function main */

.section .note.GNU-stack,"",@progbits
