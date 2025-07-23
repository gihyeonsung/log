export const getPosts = async () => {
  // const response = await fetch("/api/v1/posts");
  // const responseJson = await response.json();
  // return responseJson;

  return [
    {
      id: "1",
      title: "Hello",
      slug: "hello",
      createdAt: "2021-01-01T00:00:00Z",
      updatedAt: "2021-01-01T00:00:00Z",
      revision: 1,
      content: "Hello",
    },
    {
      id: "2",
      title: "World",
      slug: "world",
      createdAt: "2021-01-01T00:00:00Z",
      updatedAt: "2021-01-01T00:00:00Z",
      revision: 1,
      content: "World",
    },
  ];
};