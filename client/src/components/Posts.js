import React,{useEffect, useState} from 'react';
import axios from 'axios';

function Posts() {

    const [posts, setPosts] = useState('[]');

    const getPosts = () => {
    axios.get("http://localhost:9090/posts")
      .then(response => {
        setPosts(JSON.parse(JSON.stringify(response.data)))
      })
    };

    useEffect(() => {
      getPosts()
    })

  if (posts !== '[]') {
    return (
      <p>
        {posts.forEach((post) => {
          <p>{post.body}</p>
        })}
      </p>
    )
  } else {
    return (
      <p>
        No posts      
      </p>
    )
  }

}
 
export default Posts
