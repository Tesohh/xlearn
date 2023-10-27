import { z } from 'zod';

export const User = z.object({
	display: z.string(),
	username: z.string(),
	coins: z.number(),
	role: z.number(),
	joined_orgs: z.array(z.string().optional()).or(z.null())
});

export type UserType = z.infer<typeof User>;

export const Org = z.object({
	name: z.string(),
	tag: z.string(),
	secret: z.string().optional(),
	adventures: z.array(z.string()).optional()
});

export type OrgType = z.infer<typeof Org>;
