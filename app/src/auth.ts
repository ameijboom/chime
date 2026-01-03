import { SvelteKitAuth } from '@auth/sveltekit';
import Discord from '@auth/sveltekit/providers/discord';

export const REQUIRED_GUILD_ID = '1436078066991632396';
export const MANAGE_EVENTS_PERMISSION = 0x0000000200000000n; // MANAGE_EVENTS
export const CREATE_EVENTS_PERMISSION = 0x0000100000000000n; // CREATE_EVENTS

export async function checkGuildMembershipAndPermissions(accessToken: string) {
	try {
		const guildsResponse = await fetch('https://discord.com/api/v10/users/@me/guilds', {
			headers: {
				Authorization: `Bearer ${accessToken}`
			}
		});

		if (!guildsResponse.ok) {
			return false;
		}

		const guilds = await guildsResponse.json();
		const guild = guilds.find((g: any) => g.id === REQUIRED_GUILD_ID);

		if (!guild) {
			return false;
		}
		const userPermissions = BigInt(guild.permissions);

		const hasManageEvents =
			(userPermissions & MANAGE_EVENTS_PERMISSION) === MANAGE_EVENTS_PERMISSION;
		const hasCreateEvents =
			(userPermissions & CREATE_EVENTS_PERMISSION) === CREATE_EVENTS_PERMISSION;

		if (!hasManageEvents || !hasCreateEvents) {
			return false;
		}

		return true;
	} catch (error) {
		console.error('encoutered issue while retrieving user permissions:\n', error);
		return false;
	}
}

export const { handle, signIn } = SvelteKitAuth({
	providers: [
		Discord({
			authorization: {
				params: {
					scope: 'identify guilds'
				}
			}
		})
	],
	trustHost: true,
	callbacks: {
		async jwt({ token, account, user }) {
			if (account?.access_token) {
				token.accessToken = account.access_token;
			}
			if (user?.id) {
				token.userId = user.id;
			}

			return token;
		},
		async session({ session, token }) {
			if (token.accessToken) {
				session.accessToken = token.accessToken as string;
			}
			if (token.userId) {
				session.userId = token.userId as string;
			}

			return session;
		}
	}
});
