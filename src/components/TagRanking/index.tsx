import axios from 'axios'
import React, { useCallback, useEffect, useMemo, useState } from 'react'

type Tags = {
  id: string
  name: string
  numberOfReference: number
}

type TagsResponse = {
  tags: Tags[]
}

const TagRanking = () => {
  const [top10Tags, setTop10Tags] = useState<Tags[]>([])
  const [queryInput, setQueryInput] = useState<string>('')

  useEffect(() => {
    ;(async () => {
      const res = await axios.get<TagsResponse>('/api/tags', {
        params: {
          sort: 'reference',
          offset: 0,
          size: 10,
        },
      })
      setTop10Tags(res.data.tags)
    })()
  }, [])

  const divFrameClick = useCallback((tagId: string) => {
    return () => (document.location.href = `/tags/${tagId}`)
  }, [])

  const tableRows = useMemo(() => {
    return top10Tags.map((tag) => {
      return (
        <tr key={tag.id} onClick={divFrameClick(tag.id)}>
          <td className="center"></td>
          <td>
            <a href={`/tags/${tag.id}`} className="tag">
              {tag.name}
            </a>
          </td>
          <td className="center">記事数:{tag.numberOfReference}</td>
        </tr>
      )
    })
  }, [top10Tags])

  const changeQuery: React.ChangeEventHandler<HTMLInputElement> = useCallback((event) => {
    setQueryInput(event.target.value)
  }, [])

  const submitSearchForm: React.MouseEventHandler<HTMLInputElement> = (event) => {
    document.location.href = '/search?q=' + queryInput
    event.preventDefault()
  }

  return (
    <div className="tag_container">
      <h6 className="tag_title">タグランキング</h6>
      <table className="tag_table striped">
        <tbody>{tableRows}</tbody>
      </table>
      <form className="search_container">
        <h6>キーワード検索</h6>
        <input id="search_input" onChange={changeQuery} placeholder="キーワードを入力" type="text" />
        <input id="search_submit" onClick={submitSearchForm} type="submit" value="検索" />
      </form>
      <ins
        className="adsbygoogle"
        style={{ display: 'block' }}
        data-ad-client="ca-pub-3445201371824373"
        data-ad-slot="2399876013"
        data-ad-format="auto"
        data-full-width-responsive="true"
      ></ins>
    </div>
  )
}

export default TagRanking
