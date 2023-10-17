<script lang='ts'>
	import { onMount } from "svelte";
	import toast, { Toaster } from "svelte-french-toast";

    export let form;

    onMount(() => {
        if (form?.invalid) toast.error("Error while signing up", {
            style: "border-radius: 200px; background: #333; color: #fff;"
        })
        else if (form?.loginError) toast.error("Error while trying to login", {
            style: "border-radius: 200px; background: #333; color: #fff;"
        })
    })
    
</script>

<Toaster/>

<div class="flex w-screen h-screen items-center justify-center">
    
    <div class="bg-gray-300 p-10 rounded-md">
        
        <h1 class="text-center py-5 text-2xl">Sign up</h1>
    
        <form action="?/signup" method="post" class="w-48" on:submit={() => {toast.loading("Signing up", {
            style: "border-radius: 200px; background: #333; color: #fff;"
        })}}>
        
            <div class="flex flex-col items-center text-center justify-center">
                <label for="username">Username</label>
                <input name="username" id="username" type="text" class="border-gray-400 border-[0.5px] border-solid">
    
                <label for="password">Password</label>
                <input name="password" id="password" type="password" class="border-gray-400 border-[0.5px] border-solid" minlength="12">
                
                <label for="confirmPassword">Confirm Password</label>
                <input name="confirmPassword" id="confirmPassword" type="password" class="border-gray-400 border-[0.5px] border-solid" minlength="12">

                <div class="p-5">
                    <button type="submit" class="bg-black text-white flex justify-center py-2 px-4 rounded-md">Sign up</button>
                </div>

                <div class="text-[15px]">
                    <a href="/login">Login</a>
                </div>

            </div>
    
        </form>
    </div>
</div>