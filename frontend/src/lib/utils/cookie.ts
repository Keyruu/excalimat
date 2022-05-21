export function getCookie(cookieName: string): string {
	const value = `; ${document.cookie}`;
	const parts = value.split(`; ${cookieName}=`);
	if (parts.length === 2 && parts != null) return parts.pop()!.split(';').shift()!;
	return '';
}

export function getUserinfo(): any {
	const cookie = getCookie('userinfo');
	console.log(cookie);
	const buffer = atob(cookie);
	console.log(buffer);

	return JSON.parse(buffer);
}
