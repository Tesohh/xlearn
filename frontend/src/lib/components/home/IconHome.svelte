<script lang="ts">
	import { selectedTab } from '$lib/writables';
	export let iconStr: string;
	export let id: string;
</script>

<button
	on:click={() => {
		selectedTab.set(id);
	}}
>
	<div
		class="h-full w-[110px] flex items-center justify-center border-{$selectedTab == id
			? `primary border-2`
			: `none`} "
	>
		<svg class="{iconStr} h-[50px] w-[50px] bg-primary text-sm" />
	</div>
</button>
