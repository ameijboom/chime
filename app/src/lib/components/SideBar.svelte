<script lang="ts">
	import { Sidebar, SidebarGroup, SidebarItem, DropdownDivider } from 'flowbite-svelte';
	import { CirclePlusOutline } from 'flowbite-svelte-icons';

	interface Props {
		onCreateNew?: () => void;
	}

	let { onCreateNew }: Props = $props();
</script>

<Sidebar alwaysOpen backdrop={false} isOpen={true} closeSidebar={undefined} class="z-50 h-full">
	<SidebarGroup>
		<SidebarItem label="Create New" onclick={onCreateNew} class="cursor-pointer">
			{#snippet icon()}
				<CirclePlusOutline
					class="h-5 w-5 text-gray-500 transition duration-75 group-hover:text-gray-900 dark:text-gray-400 dark:group-hover:text-white"
				/>
			{/snippet}
		</SidebarItem>
		<DropdownDivider />
	</SidebarGroup>
</Sidebar>
