<script lang="ts">
	import type { Org } from '$lib/types';
	import { selectedOrg } from '$lib/writables';
	export let data: Org;
</script>

<a href="/org/{data.tag}" data-sveltekit-reload>
	<div
		class="w-[440px] h-[108px] border-{$selectedOrg == data.tag
			? `primary`
			: `inactive`} border-2 {$selectedOrg == data.tag
			? `border-l-0 rounded-l-none`
			: ``} rounded-md flex flex-col justify-center"
	>
		<div class="pl-[30px]">
			<h1 class="text-xl text-primary {$selectedOrg == data.tag ? `font-bold` : ``}">
				{data.name}
			</h1>
			<p class="text-sm text-primary {$selectedOrg == data.tag ? `font-bold` : ``}">
				{data.adventures?.length} Adventures | {data.is_unprotected
					? 'Is Protected'
					: 'Unprotected'}
			</p>
		</div>
	</div>
</a>
