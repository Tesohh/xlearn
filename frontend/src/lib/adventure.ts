import { backendUrl } from './const';
import { parseAdventure, type Adventure } from './types';

export const getAdventuresByOrg = async (orgTag: string, cookie: string) => {
	let resp: Response;

	try {
		resp = await fetch(`${backendUrl}/api/org/@${orgTag}/adventure/all`, {
			method: 'GET',
			headers: {
				Cookie: `token=${cookie}`
			}
		});
	} catch (err) {
		return { error: true, adventures: null };
	}
	if (!resp.ok) return { error: true, adventures: null };

	const adventureObj = await resp.json();

	const adventures: Adventure[] = [];

	adventureObj.forEach((adv: Object) => {
		const parsed = parseAdventure(adv);
		if (!parsed.error) {
			if (parsed.adventure) adventures.push(parsed.adventure);
		}
	});

	return { error: false, adventures: adventures };
};
