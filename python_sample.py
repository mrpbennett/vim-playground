import requests


def fetch_data(url):
    try:
        response = requests.get(url)
        response.raise_for_status()  # throws on 4xx/5xx
        return response.json()  # parse JSON into dict/list
    except requests.exceptions.RequestException as e:
        print(f"Error fetching data: {e}")
        return None


if __name__ == "__main__":
    api_url = "https://api.example.com/data"
    data = fetch_data(api_url)
    if data is not None:
        print("Fetched hello:", data)
