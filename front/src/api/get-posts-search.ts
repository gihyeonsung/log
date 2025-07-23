import type { Post } from "../type/post";

export const getPostsSearch = async (query: string): Promise<Post[]> => {
  const response = await fetch(`/api/v1/posts/search?query=${query}`);
  const responseJson = await response.json();
  return responseJson;
};