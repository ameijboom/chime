<script lang="ts">
	import { Modal as FlowbiteModal } from 'flowbite-svelte';

	interface Props {
		open: boolean;
		title: string;
		size?: 'xs' | 'sm' | 'md' | 'lg' | 'xl';
		onClose?: () => void;
		children?: any;
	}

	let { open = $bindable(), title, size = 'md', onClose, children }: Props = $props();

	function handleClose() {
		open = false;
		onClose?.();
	}
</script>

<FlowbiteModal bind:open {size} dismissable={true}>
	<div class="flex flex-col space-y-6">
		<h3 class="text-xl font-medium text-gray-900 dark:text-white">{title}</h3>
		{@render children?.()}
	</div>
</FlowbiteModal>
