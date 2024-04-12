<script lang="ts">
	import type { Adventure } from '$lib/types';

	export let data: Adventure;
</script>

<div
	class="w-[400px] h-[100px] border-primary border-2 rounded-md flex items-center justify-center"
>
	<h1 class="text-lg font-medium text-primary">{data.name}</h1>
</div>
