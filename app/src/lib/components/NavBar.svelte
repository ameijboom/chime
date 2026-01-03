<script lang="ts">
	import { Navbar, NavBrand, Button, Dropdown, DropdownItem, Avatar } from 'flowbite-svelte';
	import { UserCircleOutline, ArrowRightToBracketOutline } from 'flowbite-svelte-icons';
	import chime from '$lib/assets/Chime.png';
	import { signOut } from '@auth/sveltekit/client';
</script>

<Navbar class="border-b border-gray-200 bg-white px-4 py-2.5 dark:border-gray-700 dark:bg-gray-800">
	<NavBrand href="/">
		<img src={chime} class="me-3 h-8 sm:h-10" alt="Chime Logo" />
		<span class="self-center text-2xl font-bold whitespace-nowrap text-gray-900 dark:text-white"
			>Chime</span
		>
	</NavBrand>
	<div class="flex items-center gap-3 md:order-2">
		<Button size="lg" color="light" onclick={() => signOut()} class="gap-2">
			<ArrowRightToBracketOutline class="h-5 w-5" />
			Sign out
		</Button>
	</div>
</Navbar>
