<script lang='ts'>
	import type { OrgType } from "$lib/types";
	import { onMount } from "svelte";

    export let data: OrgType;
</script>

<a href="/org/@{data.name}" class="decoration-transparent">
    <div class="flex justify-center items-center text-center bg-gray-400 h-28">
        <div class="grid grid-cols-1 gap-2">
            <h1 class="w-48">{data.name}</h1>
            <p>{data.adventures?.length} adventures</p>
        </div>
    </div>
</a>