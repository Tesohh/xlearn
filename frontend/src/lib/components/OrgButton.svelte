<script lang='ts'>
	import type { Org } from "$lib/types";
    import Icon from '@iconify/svelte';
    export let data: Org;

</script>

<a href="/org/{data.tag}" data-sveltekit-preload-data>
    <div class="w-12 h-12 bg-blue-800 rounded-md hover:bg-blue-400 hover:rounded-sm">

        <div class="flex justify-center items-center w-full h-full group">
            <Icon icon="octicon:organization-24" color="white"/>
            <span class="absolute w-auto p-2 m-2 min-w-max left-16 rounded-md shadow-md text-white bg-blue-700 font-bold scale-0 group-hover:scale-100">{data.name}</span>
            
        </div>
        
    </div>
</a>
