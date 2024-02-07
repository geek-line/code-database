export const textToQueryValue = (text: string) => {
  const queries = text.split(/\s+/g)
  for (let j = 0; j < queries.length; j++) {
    queries[j] = encodeURIComponent(queries[j])
  }
  return queries.join('+')
}
