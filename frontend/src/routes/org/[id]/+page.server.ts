import { getAdventuresByOrg } from '$lib/adventure.js';
import { authCookieName } from '$lib/const.js';
import errorMessages from '$lib/errorMessages.js';
import { getOrgByID } from '$lib/org.js';
import { getAdventureSteps } from '$lib/steps';
import type { Step } from '$lib/types';
import { redirect, error } from '@sveltejs/kit';

export const load = async ({ locals, params, cookies }) => {
	if (!locals.user) throw redirect(303, '/login');

	const orgTag = params.id;

	const cookie = cookies.get(authCookieName);

	if (!cookie) return { error: errorMessages.errorLoadingAuth };

	const resp = await getOrgByID(orgTag, cookie);

	if (resp.error) error(404, 'Org not found');

	const data = await getAdventuresByOrg(resp.org!.tag, cookie);

	if (data.error) return { error: errorMessages.cannotLoadAdventures };

	let advSteps: Record<string, Step[]> = {};

	data.adventures?.forEach(async (adv) => {
		advSteps[adv.tag] = new Array<Step>();

		const resp = await getAdventureSteps(adv, cookie);
		console.log(resp);
	});

	return { error: null, org: resp.org, adventures: data.adventures };
};
