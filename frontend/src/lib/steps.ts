import { backendUrl } from './const';

export const getAdventureSteps = async (
	adventureID: string,
	cookie: string
): Promise<{ error: null | boolean; steps: null | Step[] }> => {
	let resp;

	try {
		resp = await fetch(`${backendUrl}/#TODO`, {
			method: 'GET',
			headers: {
				Cookie: `token=${cookie}`
			}
		});
	} catch (err) {
		return { error: true, adventures: null };
	}
	if (!resp.ok) return { error: true, steps: null };

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
