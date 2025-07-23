export const postPosts = async () => {
  await fetch("/api/v1/posts", {
    method: "POST",
    headers: {
      authorization: "secret",
    },
  });
};