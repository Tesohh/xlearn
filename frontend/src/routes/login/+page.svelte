<div class="flex w-screen h-screen items-center justify-center">
    
    <div class="bg-gray-300 p-10 rounded-md">
        
        <h1 class="text-center py-5 text-2xl">Login</h1>
    
        <form action="?/login" method="post" class="w-48">
        
            <div class="flex flex-col items-center text-center justify-center">
                <label for="username">Username</label>
                <input name="username" id="username" type="text" class="border-gray-400 border-[0.5px] border-solid">
    
                <label for="password">Password</label>
                <input name="password" id="password" type="password" class="border-gray-400 border-[0.5px] border-solid">

                <div class="p-5">
                    <button type="submit" class="bg-black text-white flex justify-center py-2 px-4 rounded-md">Login</button>
                </div>

            </div>
    
        </form>
    </div>
</div>