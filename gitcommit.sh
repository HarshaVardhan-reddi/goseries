COMMIT_MSG="$1"
echo "Starting commit process with commit message ${COMMIT_MSG}..."
git add .
git commit -m "${COMMIT_MSG}"
git push origin main