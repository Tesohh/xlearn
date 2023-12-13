<script lang='ts'>
	import { toastStyle } from "$lib/const.js";
	import { onMount } from "svelte";
	import toast, { Toaster } from "svelte-french-toast";
	import { writable } from "svelte/store";

    export let form;

    let password: string = "", confirmPassword: string = ""

    let passwordError = writable(false)
    
    onMount(() => {
        if (form?.error) toast.error(form.error, toastStyle)
    })

    function checkPasswordMatch() {
        if (password.length == 0) return passwordError.set(false)
        if (confirmPassword.length == 0) return passwordError.set(true)
        
        if (password != confirmPassword) return passwordError.set(true)
        
        passwordError.set(false)
    }
    
</script>

<Toaster/>

<div class="flex w-screen h-screen items-center justify-center">
    
    <div class="bg-gray-300 p-10 rounded-md">
        
        <h1 class="text-center py-5 text-2xl">Sign up</h1>
    
        <form action="?/signup" method="post" class="w-48" on:submit={() => {toast.loading("Signing up", toastStyle)}}>
        
            <div class="flex flex-col items-center text-center justify-center">
                <label for="username">Username</label>
                <input name="username" id="username" type="text" class="border-gray-400 border-[0.5px] border-solid" required>
    
                <label for="password">Password</label>
                <input name="password" id="password" type="password" class="border-gray-400 border-[0.5px] border-solid" minlength="12" bind:value={password} on:keyup={checkPasswordMatch}>
                
                <label for="confirmPassword">Confirm Password</label>
                <input name="confirmPassword" id="confirmPassword" type="password" class="border-gray-400 border-[0.5px] border-solid" minlength="12" bind:value={confirmPassword} on:keyup={checkPasswordMatch}>

                {#if $passwordError}
                    <p>Passwords dont match</p>
                {/if}

                <div class="p-5">

                    <button type="submit" class="bg-{$passwordError ? 'gray-500' : 'black'} text-white flex justify-center py-2 px-4 rounded-md" disabled="{$passwordError}">Sign up</button>

                </div>

                <div class="text-[15px]">

                    <a href="/login">Login</a>
                </div>

            </div>
    
        </form>
    </div>
</div>