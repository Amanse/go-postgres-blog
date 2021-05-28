import React, {useState, useEffect} from 'react'
import axios from 'axios'


const Posts = () => {
  const [posts, setPosts] = useState([])

  useEffect(() => {
    axios.get("http://localhost:9090/posts")
      .then(resp => resp.data)
      .then(data => setPosts(data))
  }, [])

  return (
    <>
      {posts.map((post, index) => (
        <div className="self-center p-10 mb-3 overflow-hidden text-center border border-indigo-200 rounded shadow shadow-lg" key={index}>
          <h2 className="pb-2 md:text-5xl sm:text-3xl text-bold">{post.body}</h2>
          <span className="text-sm text-white-200"> - {post.email}</span>
        </div>
      ))} 
    </>
  )
}

export default Posts
