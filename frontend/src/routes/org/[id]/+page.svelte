<script lang='ts'>
	import { toastStyle } from '$lib/const.js';
	import { onMount } from 'svelte';
	import toast from 'svelte-french-toast';
    import Icon from '@iconify/svelte';

    export let data;
    
    onMount(() => {
        if (data.error) toast.error(data.error, toastStyle)
    })

</script>

<div class="flex justify-center w-full text-center">

    <div class="grid grid-col-1 text-center items-center">
        <h1 class="font-semibold text-2xl p-10">{data.org?.name}</h1>
        
        <div class="grid grid-cols-1">
            <div class="flex flex-col items-center justify-center gap-20">
                {#if data.adventures}
                        {#each data.adventures as adv}
                            <div class="flex flex-col items-center gap-5">
                                <div class="p-4 bg-gray-400">
                                    <h1 class="text-lg font-semibold">{adv.name}</h1>
                                </div>
                                {#each adv.steps as step}
                                    <div class="h-24 w-24 rounded-full bg-gray-600">
                                        <div class="flex w-full h-full justify-center items-center">
                                            <Icon icon="mdi:achievement-award" height="64" class="h-48" color="white"/>
                                        </div>
                                    </div>
                                {/each}
                            </div>
                        {/each}
                {/if}
            </div> 

        </div>

    </div>
</div>