import { writable } from 'svelte/store';

export const joinPopupTrigger = writable(false);

export const tabState = writable<'Login' | 'Signup'>('Login');

export const selectedOrg = writable<string>('');
