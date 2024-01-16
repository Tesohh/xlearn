<script lang="ts">
	import { page } from '$app/stores';
	import type { Org } from '$lib/types';
	import { selectedOrg } from '$lib/writables';
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';
	export let data: Org;
</script>

<a href="/org/{data.tag}" data-sveltekit-preload-data>
	<div
		class="w-[440px] h-[108px] border-{$selectedOrg == data.tag
			? `primary`
			: `notSelected`} border-2 {$selectedOrg == data.tag
			? `border-l-0 rounded-l-none`
			: ``} rounded-md flex flex-col justify-center"
	>
		<div class="pl-[30px]">
			<h1 class="text-xl text-primary {$selectedOrg == data.tag ? `font-bold` : ``}">
				Matematica per la 4 classe
			</h1>
			<p class="text-sm text-primary {$selectedOrg == data.tag ? `font-bold` : ``}">
				50% | 100 Points
			</p>
		</div>
	</div>
</a>
