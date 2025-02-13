from fastapi import FastAPI

app = FastAPI()

@app.get("/")
def read_root():
    return {"Hello": "World"}

@app.get("/items/{item_id}")
def read_item(item_id: int):
    return {"itemd": item_id, "name": "Fake Item"}

@app.post("/items")
def create_item(item_name: str):
    return {"item_name": item_name}
