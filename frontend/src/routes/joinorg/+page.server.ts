import { authCookieName } from '$lib/const.js';
import { joinOrgByJoinCode } from '$lib/org.js';
import { fail, redirect } from '@sveltejs/kit';

export const load = async ({ locals, url, cookies }) => {
	if (locals.user == null) throw redirect(303, '/login');

	const inviteCode = url.searchParams.get('code');

	if (inviteCode == null) return { error: 'Please provide a valid join code' };

	const authCookie = cookies.get(authCookieName);

	if (authCookie == undefined) return { error: 'Auth error. Maybe page reload?' };

	let resp = await joinOrgByJoinCode(inviteCode, authCookie);
	if (resp.error) return { error: 'Error while joining the org. Invalid join code?' };

	return { error: null, org: 'Test' };
};
