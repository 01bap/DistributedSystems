// Should hopfully not catch anything

// Handles form submissions or other POST requests to "/"
export const actions = {
  default: async ({ request }) => {
    const formData = await request.formData();
    console.log("FAIL CATCH");
    console.log('Received POST data:', Object.fromEntries(formData));
    
    return { success: true };
  }
};