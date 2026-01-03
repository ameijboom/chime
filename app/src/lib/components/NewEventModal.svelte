<script lang="ts">
	import Modal from './Modal.svelte';
	import { Button, Label, Input, Textarea, Select } from 'flowbite-svelte';

	interface Props {
		open: boolean;
	}

	let { open = $bindable() }: Props = $props();

	// Form state
	let formData = $state({
		title: '',
		description: '',
		thumbnail: '',
		status: 'SCHEDULED_EVENT_STATUS_ONGOING',
		type: '',
		recurring: '',
		startDate: '',
		startTime: '',
		endDate: '',
		endTime: ''
	});

	let error = $state('');

	const statusOptions = [
		{ value: 'SCHEDULED_EVENT_STATUS_ONGOING', name: 'Ongoing' },
		{ value: 'SCHEDULED_EVENT_STATUS_ENDED', name: 'Ended' },
		{ value: 'SCHEDULED_EVENT_STATUS_ONHOLD', name: 'On Hold' }
	];

	const partyTypeOptions = [
		{ value: '', name: 'None (Optional)' },
		{ value: 'SCHEDULED_EVENT_PARTY_TYPE_LIGHT', name: 'Light Party' },
		{ value: 'SCHEDULED_EVENT_PARTY_TYPE_FULL', name: 'Full Party' }
	];

	const frequencyOptions = [
		{ value: '', name: 'None (Optional)' },
		{ value: 'SCHEDULED_EVENT_FREQUENCY_DAILY', name: 'Daily' },
		{ value: 'SCHEDULED_EVENT_FREQUENCY_WEEKLY', name: 'Weekly' },
		{ value: 'SCHEDULED_EVENT_FREQUENCY_BIWEEKLY', name: 'Biweekly' },
		{ value: 'SCHEDULED_EVENT_FREQUENCY_MONTHLY', name: 'Monthly' },
		{ value: 'SCHEDULED_EVENT_FREQUENCY_YEARLY', name: 'Yearly' }
	];

	function combineDateTime(date: string, time: string): string {
		if (!date || !time) return '';
		return `${date}T${time}`;
	}

	function handleSubmit(event: Event) {
		event.preventDefault();
		error = '';

		// Basic validation
		if (!formData.title.trim()) {
			error = 'Title is required';
			return;
		}

		if (!formData.startDate) {
			error = 'Start date is required';
			return;
		}

		if (!formData.startTime) {
			error = 'Start time is required';
			return;
		}

		if (!formData.endDate) {
			error = 'End date is required';
			return;
		}

		if (!formData.endTime) {
			error = 'End time is required';
			return;
		}

		const startDateTime = combineDateTime(formData.startDate, formData.startTime);
		const endDateTime = combineDateTime(formData.endDate, formData.endTime);

		// Validate end time is after start time
		if (new Date(endDateTime) <= new Date(startDateTime)) {
			error = 'End time must be after start time';
			return;
		}

		// Here you would submit the form data to your API
		console.log('Form submitted:', {
			...formData,
			startDateTime,
			endDateTime
		});

		// Close modal and reset form
		open = false;
		resetForm();
	}

	function resetForm() {
		formData = {
			title: '',
			description: '',
			thumbnail: '',
			status: 'SCHEDULED_EVENT_STATUS_ONGOING',
			type: '',
			recurring: '',
			startDate: '',
			startTime: '',
			endDate: '',
			endTime: ''
		};
		error = '';
	}

	function handleClose() {
		resetForm();
	}
</script>

<Modal bind:open title="Create New Event" size="lg" onClose={handleClose}>
	<form onsubmit={handleSubmit} class="space-y-5">
		{#if error}
			<div
				class="rounded-lg bg-red-50 p-3 text-sm text-red-700 dark:bg-red-900/20 dark:text-red-400"
			>
				{error}
			</div>
		{/if}

		<!-- Title -->
		<div>
			<Label for="title" class="mb-2">
				Title <span class="text-red-600 dark:text-red-500">*</span>
			</Label>
			<Input
				id="title"
				type="text"
				bind:value={formData.title}
				placeholder="e.g., Weekly Raid Night"
				required
			/>
		</div>

		<!-- Description -->
		<div>
			<Label for="description" class="mb-2">Description</Label>
			<Textarea
				id="description"
				bind:value={formData.description}
				placeholder="Add details about your event..."
				rows={3}
			/>
		</div>

		<!-- Date & Time -->
		<div>
			<Label class="mb-3 block">
				Schedule <span class="text-red-600 dark:text-red-500">*</span>
			</Label>
			<div class="space-y-3">
				<div class="grid gap-3 sm:grid-cols-2">
					<div>
						<Label for="startDate" class="mb-1.5 text-xs text-gray-600 dark:text-gray-400">
							Start Date
						</Label>
						<Input id="startDate" type="date" bind:value={formData.startDate} required />
					</div>
					<div>
						<Label for="startTime" class="mb-1.5 text-xs text-gray-600 dark:text-gray-400">
							Start Time
						</Label>
						<Input id="startTime" type="time" bind:value={formData.startTime} required />
					</div>
				</div>
				<div class="grid gap-3 sm:grid-cols-2">
					<div>
						<Label for="endDate" class="mb-1.5 text-xs text-gray-600 dark:text-gray-400">
							End Date
						</Label>
						<Input id="endDate" type="date" bind:value={formData.endDate} required />
					</div>
					<div>
						<Label for="endTime" class="mb-1.5 text-xs text-gray-600 dark:text-gray-400">
							End Time
						</Label>
						<Input id="endTime" type="time" bind:value={formData.endTime} required />
					</div>
				</div>
			</div>
		</div>

		<!-- Status -->
		<div>
			<Label for="status" class="mb-2">
				Status <span class="text-red-600 dark:text-red-500">*</span>
			</Label>
			<Select id="status" bind:value={formData.status} items={statusOptions} required />
		</div>

		<!-- Optional Settings -->
		<details class="group">
			<summary
				class="dark:hover:bg-gray-750 flex cursor-pointer items-center justify-between rounded-lg border border-gray-200 bg-gray-50 px-4 py-3 hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-800"
			>
				<span class="text-sm font-medium text-gray-700 dark:text-gray-300">Optional Settings</span>
				<span class="text-gray-500 transition-transform group-open:rotate-180 dark:text-gray-400">
					▼
				</span>
			</summary>
			<div class="mt-3 space-y-3 rounded-lg border border-gray-200 p-4 dark:border-gray-700">
				<div>
					<Label for="partyType" class="mb-1.5 text-xs text-gray-600 dark:text-gray-400">
						Party Type
					</Label>
					<Select id="partyType" bind:value={formData.type} items={partyTypeOptions} />
				</div>
				<div>
					<Label for="recurring" class="mb-1.5 text-xs text-gray-600 dark:text-gray-400">
						Recurring Frequency
					</Label>
					<Select id="recurring" bind:value={formData.recurring} items={frequencyOptions} />
				</div>
				<div>
					<Label for="thumbnail" class="mb-1.5 text-xs text-gray-600 dark:text-gray-400">
						Thumbnail URL
					</Label>
					<Input
						id="thumbnail"
						type="url"
						bind:value={formData.thumbnail}
						placeholder="https://example.com/image.jpg"
					/>
				</div>
			</div>
		</details>

		<!-- Action Buttons -->
		<div class="flex justify-end gap-2 border-t border-gray-200 pt-4 dark:border-gray-700">
			<Button color="alternative" onclick={() => (open = false)}>Cancel</Button>
			<Button type="submit" color="blue">Create Event</Button>
		</div>
	</form>
</Modal>
