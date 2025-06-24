import { env } from "$env/dynamic/private";
// import { env } from "$env/dynamic/public";
import { VITE_PUBLIC_BACKEND_URL } from "$env/static/private";
import { type RequestHandler } from "@sveltejs/kit";

export const GET: RequestHandler = async () => {

    console.log("API ENDPOINT: DYNAMIC_Backend=" + env.PUBLIC_BACKEND_URL);
    console.log("API ENDPOINT: STATIC_Backend=" + VITE_PUBLIC_BACKEND_URL);

    return new Response(
        JSON.stringify({ "VITE_PUBLIC_BACKEND_URL": VITE_PUBLIC_BACKEND_URL }),
        { status: 200 },
    );
};