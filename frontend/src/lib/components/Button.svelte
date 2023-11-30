<script lang='ts'>
	import type { OrgType } from "$lib/types";

    import { IconBrandCodesandbox, IconPlus } from '@tabler/icons-svelte';

    export let data: OrgType | null;

</script>

<a href="{data?.name == undefined ? "/" : `/org/${data.name}`}">

    <div class="w-12 h-12 bg-blue-800 rounded-md hover:bg-blue-400 hover:rounded-sm">

        <div class="flex justify-center items-center w-full h-full group">

            {#if data}
                <IconBrandCodesandbox color="white" />
                <span class="absolute w-auto p-2 m-2 min-w-max left-16 rounded-md shadow-md text-white bg-blue-700 font-bold scale-0 group-hover:scale-100">{data.name}</span>
            {:else}
                <IconPlus color="white"/>
                <span class="absolute w-auto p-2 m-2 min-w-max left-16 rounded-md shadow-md text-white bg-blue-700 font-bold scale-0 group-hover:scale-100">Join new org</span>
            {/if}
        
        </div>
        
    </div>
</a>