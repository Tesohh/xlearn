<script lang='ts'>
	import { page } from "$app/stores";
    
</script>

<div class="flex justify-center items-center w-screen h-screen">
    <div class="grid grid-cols-1 text-center bg-gray-400 p-10 rounded-md">
        <h1 class="p-4 text-2xl text-white">Error {$page.status}</h1>
        <img src={`https://http.cat/${$page.status}`} alt="Cat HTTP Status Image" class="max-h-[350px]"/>
        <a href="/" class="p-4"><button class="bg-gray-600 h-12 w-36 rounded-md">Back to safety</button></a>
    </div>
</div>