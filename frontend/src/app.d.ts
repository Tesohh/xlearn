// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces

interface User {
	username: string;
	level: number;
}

declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			user: User | null;
		}
		// interface PageData {}
		// interface Platform {}
	}
}

export {};
