import React, { useState, useEffect } from 'react';
import { Pane, majorScale, Heading, Paragraph, Link, Spinner } from 'evergreen-ui'
import CollectionItemElement from './components/CollectionItem'
import Search from './components/Search'
import ReactGA from 'react-ga'
import InfiniteScroll from 'react-infinite-scroll-component';
import CollectionSummaryFilter from './components/CollectionSummaryFilter'
import ArendtImage from './static/arendt.jpg'
import { fetchSummary, CollectionSummary, doSearch, CollectionItem } from './api'

function App() {
  const [query, setQuery] = useState<string>("")
  const [searchWords, setSearchWords] = useState<string[]>([])
  const [collectionsFilter, setCollectionsFilter] = useState<CollectionSummary[]>([])
  const [results, setResults] = useState<CollectionItem[]>([])
  const [summary, setSummary] = useState<CollectionSummary[]>([])
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [summaryIsLoading, setSummaryIsLoading] = useState<boolean>(true)
  const [page, setPage] = useState(0)
  const [hasMore, setHasMore] = useState(true)

  useEffect(() => {
    if (query) {
      setSearchWords(query.trim().split(" "))
    }
  }, [query])

  useEffect(() => {
    const updateSummary = async () => {
      setSummaryIsLoading(true)
      try {
        const res = await fetchSummary(query)
        setSummary(res.summary)
      } finally {
        setSummaryIsLoading(false)
      }
    }
    updateSummary()

  }, [query])

  useEffect(() => {

    setPage(0)
    setHasMore(true)
    const updateResults = async () => {
      setIsLoading(true)
      const res = await doSearch(query, collectionsFilter.map(c => c.collectionId), 0)
      setResults(res.results)
      setIsLoading(false)
    }

    updateResults()

  }, [query, collectionsFilter])

  const fetchMoreData = async () => {
    const res = await doSearch(query, collectionsFilter.map(c => c.collectionId), page + 1)
    if (res.results.length === 0) {
      setHasMore(false)
      return
    }

    setResults(results.concat(res.results))
    setPage(page + 1)
  }

  const doNewSearch = (query: string) => {
    ReactGA.event({
      category: 'User',
      action: 'Searched'
    })

    setQuery(query)
  }

  return (
    <Pane display="flex" minHeight={"100%"} minWidth={majorScale(150)}>
      <Pane width={majorScale(70)}>
        <Pane position="sticky" top={0} borderRight width={majorScale(70)} marginX='auto' padding={majorScale(2)} display="flex" alignItems="center" flexDirection="column">
          <Pane display="flex" alignItems="center" justifyContent="center" marginBottom={majorScale(2)}>
            <Pane width={majorScale(18)} marginRight={majorScale(2)}>
              <img alt="Hannah Arendt" src={ArendtImage} width={"100%"} />
            </Pane>
            <Heading fontFamily={"Typewriter"} fontSize={42} lineHeight={"1.2em"}>Arendt Archives</Heading>
          </Pane>
          <Search onSubmit={doNewSearch} />
          <CollectionSummaryFilter selected={collectionsFilter} onChange={setCollectionsFilter} summary={summary} isLoading={summaryIsLoading} />

          <Pane marginTop={majorScale(2)} borderTop padding={majorScale(2)} textAlign="center">
            <Paragraph>This project uses documents sourced from the <Link target="_blank" href="https://www.loc.gov/collections/hannah-arendt-papers/">Library of Congress Hannah Arendt Papers Collection</Link>. See <Link target="_blank" href="https://www.loc.gov/collections/hannah-arendt-papers/about-this-collection/rights-and-access/">Rights and Access for usage</Link>.  It was built by <Link href="https://robertmcqueen.com/" target="_blank"> Robert McQueen</Link>. Please use the Google Chrome browser for best results.</Paragraph>
          </Pane>
        </Pane>
      </Pane>

      <Pane flex={1} backgroundColor="#FAFAFA" padding={majorScale(2)}>
        <Pane display="flex" height={"100%"}>
          {isLoading && (
            <Pane margin="auto" display="flex" flexDirection="column" alignItems="center" width={"300px"} border backgroundColor={"white"} padding={majorScale(2)}>
              <Heading size={700} marginBottom={majorScale(2)}>Searching...</Heading>
              <Spinner marginBottom={majorScale(2)} />
              <Paragraph textAlign="center">Searching...</Paragraph>
            </Pane>
          )}
          {!isLoading &&
            <Pane height={"100%"} width="100%" display="flex">
              {results.length === 0 && (
                <Pane margin="auto" width={majorScale(50)} padding={majorScale(2)} border backgroundColor="white" textAlign="center">
                  <Pane><Heading size={700}>Darn, no results :(</Heading></Pane>
                  <Paragraph>Try a different search.</Paragraph>
                </Pane>
              )}
              {results.length > 0 &&
                <InfiniteScroll dataLength={results.length} next={fetchMoreData} hasMore={hasMore} loader={<Spinner />} endMessage={<Paragraph textAlign="center" size={500}>No More Items Found</Paragraph>}>
                  {results.map(r => (
                    <CollectionItemElement key={r.locId} item={r} searchWords={searchWords} />
                  ))}
                </InfiniteScroll>
              }
            </Pane>
          }
        </Pane>
      </Pane>
    </Pane>
  );
}

export default App;
