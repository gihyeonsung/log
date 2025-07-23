import { useQuery } from "@tanstack/react-query";
import { useDebouncedCallback, useDebouncedValue } from "@tanstack/react-pacer";

import { getPosts } from "./api/get-posts";
import type { Post } from "./type/post";
import { useState } from "react";
import { getPostsSearch } from "./api/get-posts-search";

const PostItem = ({ post }: { post: Post }) => {
  const updatedAt = new Date(post.updatedAt);
  const year = updatedAt.getFullYear();
  const month = String(updatedAt.getMonth() + 1).padStart(2, "0");
  const day = String(updatedAt.getDate()).padStart(2, "0");

  return (
    <a href={`/posts/${post.slug}`}>
      <div className="text-neutral-500 font-bold text-sm hover:bg-neutral-100 cursor-pointer">
        {year}.{month}.{day} {post.title}
      </div>
    </a>
  );
};

const PostItemList = ({ posts }: { posts: Post[] }) => {
  return (
    <div className="flex flex-col">
      <div className="text-neutral-500 font-bold pb-2 text-sm">포스트</div>

      {posts.map((post) => (
        <PostItem key={post.id} post={post} />
      ))}
    </div>
  );
};

const Header = ({ site }: { site: string }) => {
  return (
    <div className="text-neutral-500 font-bold pb-2 text-sm flex flex-row justify-between items-center">
      <div>{site}</div>

      <SearchBar />
    </div>
  );
};

const SearchBar = () => {
  const [query, setQuery] = useState("");
  const [posts, setPosts] = useState<Array<Post>>([]);

  const handleSearch = async (query: string) => {
    if (!query) {
      setPosts([]);
      return;
    }

    const postsSearched = await getPostsSearch(query);
    setPosts(postsSearched);
    return postsSearched;
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setQuery(e.target.value);
  };

  return (
    <div className="text-neutral-500 pb-2 text-sm">
      <input
        type="text"
        placeholder="Search"
        value={query}
        onChange={handleChange}
      />
    </div>
  );
};

const App = () => {
  const { data } = useQuery<Post[]>({
    queryKey: ["posts"],
    queryFn: getPosts,
  });

  return (
    <div className="p-4">
      <Header site="log" />

      {data && <PostItemList posts={data} />}
    </div>
  );
};

export default App;
