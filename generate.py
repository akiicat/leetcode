import time
from pathlib import Path
import frontmatter
from python_graphql_client import GraphqlClient

def get_data(total = -1):
    # Instantiate the client with an endpoint.
    client = GraphqlClient(endpoint="https://leetcode.com/graphql/")

    # Create the query string and variables required for the request.
    query = """
        query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {
          problemsetQuestionList: questionList(
            categorySlug: $categorySlug
            limit: $limit
            skip: $skip
            filters: $filters
          ) {
            total: totalNum
            questions: data {
              acRate
              difficulty
              questionFrontendId
              isPaidOnly
              title
              titleSlug
              topicTags {
                name
                slug
              }
            }
          }
        }
    """
    variables = {
            "categorySlug": "",
            "skip": 0,
            "limit": 1000,
            "filters": {}
            }

    res = []

    # do while loop
    while (True):
        # user request the number of question
        # limit the minimum number of question
        # otherwise user does not request the number of question
        # we will set `total` later after query
        if total >= 0:
            variables['limit'] = min(total - variables['skip'], 1000)

        # Synchronous request
        data = client.execute(query=query, variables=variables)

        # set valid total question number
        if total < 0 or total > data['data']['problemsetQuestionList']['total']:
            total = data['data']['problemsetQuestionList']['total']

        # successful get the `variables['limit']` number of data
        variables['skip'] += variables['limit']


        res.extend(data['data']['problemsetQuestionList']['questions'])
        # print(len(data['data']['problemsetQuestionList']['questions']))

        # time.sleep(10)

        if variables['skip'] >= total:
            break

    return res
    # return data['data']['problemsetQuestionList']['questions']

# data = get_data
# print(len(data))
# print(data[0])
# {
# 'acRate': 49.528488184612684,
# 'difficulty': 'Easy',
# 'questionFrontendId': '1',
# 'isPaidO()nly': False,
# 'title': 'Two Sum',
# 'titleSlug': 'two-sum',
# 'topicTags': [
#   {'name': 'Array', 'slug': 'array'},
#   {'name': 'Hash Table', 'slug': 'hash-table'}
# ]}

default_content = """
# {}. {}

=== "C++"

    $Time: O()$

    $Space: O()$

    ```c++
    ```
"""

for data in get_data():

    path = "docs/leetcode/{}-{}.md".format(data['questionFrontendId'], data['titleSlug'])

    # create a Path object with the path to the file
    if Path(path).is_file():
        pass
    else:
        post = frontmatter.loads("")
        
        # post['title'] = "{}. {}".format(data['questionFrontendId'], data['title'])
        post['tags'] = [tag['name'] for tag in data['topicTags']]
        post['tags'].append('Unsolved')

        post.content = default_content.format(data['questionFrontendId'], data['title'])

        # write back
        with open(path, 'wb') as f:
            frontmatter.dump(post,f)
