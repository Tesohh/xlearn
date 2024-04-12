<script lang="ts">
	import type { Org } from '$lib/types';
	import { selectedOrg } from '$lib/writables';
	export let data: Org;
</script>

<a href="/{data.tag}" data-sveltekit-preload-data>
	<div
		class="w-[440px] h-[108px] border-{$selectedOrg == data.tag
			? `primary`
			: `inactive`} border-2 {$selectedOrg == data.tag
			? `border-l-0 rounded-l-none`
			: ``} rounded-md flex flex-row justify-center text-center"
	>
		<div class="w-2/3 h-full flex flex-col justify-center">
			<h1 class="text-xl text-primary {$selectedOrg == data.tag ? `font-bold` : ``}">
				{data.name}
			</h1>
			<p class="text-sm text-primary {$selectedOrg == data.tag ? `font-bold` : ``}">
				{data.adventures?.length} Adventures | {data.is_unprotected
					? 'Is Protected'
					: 'Unprotected'}
			</p>
		</div>
        <div class="w-2/3">
            <img src="https://www.imagelighteditor.com/img/bg-after.jpg" alt="" class="h-full w-fit">

        </div>
	</div>
</a>
