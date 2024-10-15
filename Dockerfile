FROM python:3.12-slim AS build
WORKDIR /opt/notigram/

RUN pip install poetry
COPY pyproject.toml poetry.lock ./
RUN poetry install --without dev --no-root

COPY notigram notigram


FROM build AS production

COPY --from=build /opt/notigram /opt/notigram

EXPOSE 8000/tcp
CMD ["poetry", "run", "gunicorn", "--worker-class=uvicorn.workers.UvicornWorker", "--bind=0.0.0.0:8000", "notigram:app"]
