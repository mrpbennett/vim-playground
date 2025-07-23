use serde::Deserialize;

#[derive(Debug, Deserialize)]
struct ApiResponse {
    // e.g. field1: String,
    //      field2: i32,
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let url = "https://api.example.com/data";
    let resp = reqwest::get(url).await?;

    if resp.status().is_success() {
        let data: ApiResponse = resp.json().await?;
        println!("Fetched data: {:#?}", data);
    } else {
        eprintln!("HTTP error: {}", resp.status());
    }

    Ok(())
}
