Here is a detailed structured description of the **Document Approval Workflow** process:

---

### **Workflow Name**: Document Approval

#### **Description**:
The workflow facilitates the review and approval of documents. It involves three main steps: submission, review, and finalization, supported by specific actors and conditions to ensure proper flow and accountability.

---

### **Steps**:

#### **1. Submit Document**
- **Objective**: To initiate the workflow by allowing a user to submit a document for review.
- **Actor**: 
  - **John Doe** (Role: User) is responsible for submitting the document.
- **Inputs**:
  - The document to be reviewed (`Document.pdf`).
- **Outputs**:
  - A confirmation that the document has been successfully submitted.
- **Process**:
  - The user uploads the document via the system.
  - A confirmation output is generated to acknowledge receipt.
- **Next Step**:
  - The workflow moves to **Step 2: Review Document**.

---

#### **2. Review Document**
- **Objective**: To evaluate the submitted document and decide whether to approve or reject it.
- **Actor**: 
  - **Jane Smith** (Role: Reviewer) is tasked with evaluating the document.
- **Inputs**:
  - The document submitted in Step 1.
- **Outputs**:
  - A decision: either `Approved` or `Rejected`.
- **Conditions**:
  - The document must pass a validity check (e.g., format, completeness).
  - Condition expression: `valid(document)`.
- **Process**:
  - The reviewer assesses the document's content.
  - Based on the review, the reviewer marks the document as either `Approved` or `Rejected`.
- **Next Step**:
  - The workflow transitions to **Step 3: Finalize Document**.

---

#### **3. Finalize Document**
- **Objective**: To perform final actions based on the review decision and notify relevant parties.
- **Actor**: 
  - **System Automation** (Role: System) is responsible for handling finalization tasks automatically.
- **Inputs**:
  - The decision from Step 2 (`Approved` or `Rejected`).
- **Outputs**:
  - A notification confirming that the document has been finalized.
  - Notifications sent to appropriate stakeholders.
- **Conditions**:
  - A decision must exist before finalization.
  - Condition expression: `decision != null`.
- **Process**:
  - If the document is **Approved**, the system:
    - Sends a notification to the user confirming approval.
    - Archives the approved document.
  - If the document is **Rejected**, the system:
    - Notifies the user about the rejection.
    - Provides feedback if necessary.
- **Next Step**:
  - None. This marks the end of the workflow.

---

### **Actors**:
1. **John Doe** (Role: User): Submits the document in Step 1.
2. **Jane Smith** (Role: Reviewer): Reviews and decides the document’s status in Step 2.
3. **System Automation** (Role: System): Finalizes the workflow based on the decision in Step 3.

---

### **Triggers**:
- The workflow starts automatically when a document is uploaded.
  - **Trigger Type**: Event-based.
  - **Trigger Criteria**: `document_uploaded`.

---

### **End State**:
- **Objective**: The workflow reaches its final state when the document is either approved or rejected, and all relevant parties have been notified.
- **Outputs**:
  - Finalized notifications.
  - Document archived if approved.

---

### **Metrics**:
1. **Total Time**: The workflow is expected to take approximately 2 hours.
2. **Completed Steps**: Tracks the number of completed steps (e.g., 2 out of 3).
3. **Pending Steps**: Tracks the remaining steps (e.g., 1 at a given time).
4. **Bottlenecks**: Identifies any steps causing delays (e.g., if Step 2 takes too long).

---

### **Summary of Flow**:
1. **Start Trigger**: Workflow begins when the user uploads a document.
2. **Step 1**: User submits the document.
3. **Step 2**: Reviewer evaluates the document and makes a decision.
4. **Step 3**: System finalizes the document's status, notifies stakeholders, and completes the workflow.
5. **End State**: Document approval or rejection is finalized, and notifications are sent.

This structured explanation captures all elements of the workflow with clarity and detail, ensuring all aspects of the process are covered. Let me know if you'd like any additional information!