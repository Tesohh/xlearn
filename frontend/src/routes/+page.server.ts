import { authCookieName } from '$lib/const.js';
import errorMessages from '$lib/errorMessages.js';
import { joinOrgByJoinCode, leaveOrgById } from '$lib/org.js';
import { redirect } from '@sveltejs/kit';

export const load = async ({ locals, cookies }) => {
	console.log(locals);
	if (!locals.user) throw redirect(303, '/login');
};

// Server actions
export const actions = {
	joinorg: async ({ request, cookies, locals }) => {
		if (!locals.user) return { error: errorMessages.notLogged };

		const userCookie = cookies.get(authCookieName);

		if (!userCookie) return { error: errorMessages.notLogged };

		const data = await request.formData();

		const joinCode = data.get('code')?.toString();

		if (joinCode == null) return { error: errorMessages.noValidCodeFound };

		const resp = await joinOrgByJoinCode(joinCode, userCookie);

		if (resp.error) return { error: errorMessages.errorWhileJoiningOrg };

		throw redirect(303, '/');
	},

	leaveorg: async ({ request, cookies, locals }) => {
		if (!locals.user) return { error: errorMessages.notLogged };

		const userCookie = cookies.get(authCookieName);

		if (!userCookie) return { error: errorMessages.notLogged };

		const data = await request.formData();

		const leftCode = data.get('code');

		if (leftCode == null) return { error: errorMessages.noValidCodeFound };

		const resp = await leaveOrgById(leftCode.toString(), userCookie);

		if (resp.error) return { error: errorMessages.errorWhileLeavingOrg };

		throw redirect(303, '/');
	}
};
