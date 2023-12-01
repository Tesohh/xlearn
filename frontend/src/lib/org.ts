import { backendUrl } from './const';
import { Org, type OrgType, parseOrg } from './types';

export const getOrgByID = async (
	orgID: string,
	cookie: string
): Promise<{ error: boolean; org: OrgType | null }> => {
	let resp;

	try {
		resp = await fetch(`${backendUrl}/api/org/@${orgID}`, {
			method: 'GET',
			headers: {
				Cookie: `token=${cookie}`
			}
		});
	} catch (err) {
		return { error: true, org: null };
	}

	if (!resp.ok) return { error: true, org: null };

	const parsed = parseOrg(await resp.json());

	if (parsed.error) return { error: true, org: null };

	return { error: false, org: parsed.org };
};

export const joinOrgByJoinCode = async (joinCode: string, cookie: string) => {
	let resp;
	try {
		resp = await fetch(`${backendUrl}/api/user/org/join/${joinCode}`, {
			method: 'POST',
			headers: {
				Cookie: `token=${cookie}`
			}
		});
	} catch (err) {
		return { error: true };
	}

	console.log(await resp.text());
	if (!resp.ok) return { error: true };

	return { error: false };
};

export const leaveOrgById = async (orgCode: string, cookie: string) => {
	let resp;
	try {
		resp = await fetch(`${backendUrl}/api/user/org/leave/${orgCode}`, {
			method: 'POST',
			headers: {
				Cookie: `token=${cookie}`
			}
		});
	} catch (err) {
		return { error: true };
	}

	if (!resp.ok) return { error: true };

	return { error: false };
};
