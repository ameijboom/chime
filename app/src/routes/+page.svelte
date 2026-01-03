<script lang="ts">
	import NavBar from '$lib/components/NavBar.svelte';
	import NewEventModal from '$lib/components/NewEventModal.svelte';
	import { Button, Card } from 'flowbite-svelte';
	import { CalendarMonthOutline, ClockOutline, CirclePlusOutline } from 'flowbite-svelte-icons';
	import type { PageData } from './$types';

	const { data }: { data: PageData } = $props();

	let showNewEventModal = $state(false);

	function handleCreateNew() {
		showNewEventModal = true;
	}

	// Mock data for demonstration
	const mockEvents = [
		{
			id: 1,
			title: 'Weekly Raid Night',
			description: 'Join us for our weekly progression raid!',
			status: 'Ongoing',
			startTime: '2024-01-15T19:00',
			endTime: '2024-01-15T22:00',
			type: 'Full Party',
			recurring: 'Weekly'
		}
	];
</script>

<div class="flex h-screen flex-col bg-gray-50 dark:bg-gray-900">
	<!-- Navigation Bar -->
	<NavBar />

	<!-- Main Content -->
	<main class="flex-1 overflow-y-auto">
		<div class="mx-auto max-w-7xl px-4 py-8 sm:px-6 lg:px-8">
			<!-- Header Section -->
			<div class="mb-8 flex items-center justify-between">
				<div>
					<h1 class="text-4xl font-bold text-gray-900 dark:text-white">Events</h1>
					<p class="mt-2 text-lg text-gray-600 dark:text-gray-400">
						Manage and track your scheduled events
					</p>
				</div>
				<Button size="lg" color="blue" onclick={handleCreateNew}>
					<CirclePlusOutline class="me-2 h-5 w-5" />
					Create Event
				</Button>
			</div>

			<!-- Events Grid -->
			{#if mockEvents.length === 0}
				<div
					class="rounded-lg border-2 border-dashed border-gray-300 bg-white p-12 text-center dark:border-gray-700 dark:bg-gray-800"
				>
					<CalendarMonthOutline class="mx-auto mb-4 h-12 w-12 text-gray-400" />
					<h3 class="mb-2 text-lg font-semibold text-gray-900 dark:text-white">No events yet</h3>
					<p class="mb-4 text-sm text-gray-500 dark:text-gray-400">
						Get started by creating your first event
					</p>
					<Button color="blue" onclick={handleCreateNew}>
						<CirclePlusOutline class="me-2 h-4 w-4" />
						Create Event
					</Button>
				</div>
			{:else}
				<div class="grid gap-4 md:grid-cols-2 xl:grid-cols-3">
					{#each mockEvents as event (event.id)}
						<div
							class="group cursor-pointer rounded-lg border border-gray-200 bg-white p-6 transition-all hover:border-blue-500 hover:shadow-md dark:border-gray-700 dark:bg-gray-800 dark:hover:border-blue-600"
						>
							<!-- Header -->
							<div class="mb-4 flex items-start justify-between gap-3">
								<h3 class="text-lg font-semibold text-gray-900 dark:text-white">
									{event.title}
								</h3>
								<span
									class="shrink-0 rounded-md bg-green-100 px-2 py-1 text-xs font-medium text-green-700 dark:bg-green-900/30 dark:text-green-400"
								>
									{event.status}
								</span>
							</div>

							<!-- Description -->
							{#if event.description}
								<p class="mb-4 text-sm text-gray-600 dark:text-gray-400">
									{event.description}
								</p>
							{/if}

							<!-- Metadata -->
							<div class="space-y-2 text-sm">
								<div class="flex items-center gap-2 text-gray-700 dark:text-gray-300">
									<CalendarMonthOutline class="h-4 w-4 text-gray-400" />
									<span>{new Date(event.startTime).toLocaleDateString()}</span>
								</div>
								<div class="flex items-center gap-2 text-gray-700 dark:text-gray-300">
									<ClockOutline class="h-4 w-4 text-gray-400" />
									<span>
										{new Date(event.startTime).toLocaleTimeString([], {
											hour: '2-digit',
											minute: '2-digit'
										})}
										-
										{new Date(event.endTime).toLocaleTimeString([], {
											hour: '2-digit',
											minute: '2-digit'
										})}
									</span>
								</div>
							</div>

							<!-- Tags -->
							{#if event.type || event.recurring}
								<div class="mt-4 flex flex-wrap gap-2">
									{#if event.type}
										<span
											class="rounded bg-blue-100 px-2 py-1 text-xs text-blue-700 dark:bg-blue-900/30 dark:text-blue-400"
										>
											{event.type}
										</span>
									{/if}
									{#if event.recurring}
										<span
											class="rounded bg-purple-100 px-2 py-1 text-xs text-purple-700 dark:bg-purple-900/30 dark:text-purple-400"
										>
											{event.recurring}
										</span>
									{/if}
								</div>
							{/if}
						</div>
					{/each}
				</div>
			{/if}
		</div>
	</main>
</div>

<!-- New Event Modal -->
<NewEventModal bind:open={showNewEventModal} />
