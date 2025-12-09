# Frontend UI Reference - Dynamic IVR Campaign Form

## Campaign Form Layout

The campaign form has been redesigned to accommodate the new dynamic IVR features. Below is a visual reference of the new UI structure.

### Form Structure

```
┌────────────────────────────────────────────────────────────────┐
│  Create Campaign                                          [X]   │
├────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ╔═══════════════════════════════════════════════════════╗    │
│  ║ Basic Information                                      ║    │
│  ╚═══════════════════════════════════════════════════════╝    │
│                                                                 │
│  Campaign Name *                                               │
│  ┌─────────────────────────────────────────────────────┐      │
│  │ Summer Sale 2025                                     │      │
│  └─────────────────────────────────────────────────────┘      │
│                                                                 │
│  Description *                                                 │
│  ┌─────────────────────────────────────────────────────┐      │
│  │ Promotional campaign for summer products            │      │
│  │                                                       │      │
│  └─────────────────────────────────────────────────────┘      │
│                                                                 │
│  Default Language *                                            │
│  ┌─────────────────────────────────────────────────────┐      │
│  │ EN - English                                    ▼   │      │
│  └─────────────────────────────────────────────────────┘      │
│                                                                 │
│  Intro Text *                                                  │
│  ┌─────────────────────────────────────────────────────┐      │
│  │ Welcome to our service. We have exclusive offers    │      │
│  │ for you today.                                       │      │
│  │                                                       │      │
│  └─────────────────────────────────────────────────────┘      │
│  This message will be played at the start of the call          │
│                                                                 │
│  ☑ Campaign is active                                         │
│                                                                 │
├────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ╔═══════════════════════════════════════════════════════╗    │
│  ║ IVR Actions                      [+] Add Action      ║    │
│  ╚═══════════════════════════════════════════════════════╝    │
│  Define what happens when users press keys                     │
│                                                                 │
│  ┌─────────────────────────────────────────────────────┐      │
│  │ Action 1                                       [🗑]  │      │
│  │                                                       │      │
│  │ Action Type           Key Press                     │      │
│  │ ┌──────────────┐     ┌──────────┐                 │      │
│  │ │ Information▼ │     │    1     │                 │      │
│  │ └──────────────┘     └──────────┘                 │      │
│  │                                                       │      │
│  │ Message (Text or Audio URL)                          │      │
│  │ ┌────────────────────────────────────────────────┐ │      │
│  │ │ Get 50% off on all summer collections.         │ │      │
│  │ │ Visit our website for more details.            │ │      │
│  │ └────────────────────────────────────────────────┘ │      │
│  │ Use text for speech or provide an audio file URL    │      │
│  └─────────────────────────────────────────────────────┘      │
│                                                                 │
│  ┌─────────────────────────────────────────────────────┐      │
│  │ Action 2                                       [🗑]  │      │
│  │                                                       │      │
│  │ Action Type           Key Press                     │      │
│  │ ┌──────────────┐     ┌──────────┐                 │      │
│  │ │ Forward    ▼ │     │    2     │                 │      │
│  │ └──────────────┘     └──────────┘                 │      │
│  │                                                       │      │
│  │ Forward to Phone Number                              │      │
│  │ ┌────────────────────────────────────────────────┐ │      │
│  │ │ +11234567890                                    │ │      │
│  │ └────────────────────────────────────────────────┘ │      │
│  │ Call will be forwarded to this number                │      │
│  └─────────────────────────────────────────────────────┘      │
│                                                                 │
│  ┌─────────────────────────────────────────────────────┐      │
│  │ Action 3                                       [🗑]  │      │
│  │                                                       │      │
│  │ Action Type           Key Press                     │      │
│  │ ┌──────────────┐     ┌──────────┐                 │      │
│  │ │ Information▼ │     │    3     │                 │      │
│  │ └──────────────┘     └──────────┘                 │      │
│  │                                                       │      │
│  │ Message (Text or Audio URL)                          │      │
│  │ ┌────────────────────────────────────────────────┐ │      │
│  │ │ https://example.com/audio/promo.mp3            │ │      │
│  │ │                                                  │ │      │
│  │ └────────────────────────────────────────────────┘ │      │
│  │ Use text for speech or provide an audio file URL    │      │
│  └─────────────────────────────────────────────────────┘      │
│                                                                 │
├────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌──────────┐                         ┌──────────┐            │
│  │  Cancel  │                         │  Create  │            │
│  └──────────┘                         └──────────┘            │
│                                                                 │
└────────────────────────────────────────────────────────────────┘
```

## Field Descriptions

### Basic Information Section

| Field              | Type            | Required | Description                                   |
| ------------------ | --------------- | -------- | --------------------------------------------- |
| Campaign Name      | Text Input      | Yes      | Unique identifier for the campaign            |
| Description        | Textarea        | Yes      | Detailed description of campaign purpose      |
| Default Language   | Select Dropdown | Yes      | Language for IVR prompts (EN, ES, FR, DE, HI) |
| Intro Text         | Textarea        | Yes      | Message played at the start of every call     |
| Campaign is active | Checkbox        | No       | Determines if campaign can receive calls      |

### IVR Actions Section

Each action consists of:

#### For Information Actions:

| Field       | Type                         | Required | Description                        |
| ----------- | ---------------------------- | -------- | ---------------------------------- |
| Action Type | Select (Information/Forward) | Yes      | Type of action to perform          |
| Key Press   | Number Input (1-9)           | Yes      | Which key user will press          |
| Message     | Textarea                     | Yes      | Text to speak OR URL to audio file |

#### For Forward Actions:

| Field            | Type                         | Required | Description                             |
| ---------------- | ---------------------------- | -------- | --------------------------------------- |
| Action Type      | Select (Information/Forward) | Yes      | Type of action to perform               |
| Key Press        | Number Input (1-9)           | Yes      | Which key user will press               |
| Forward to Phone | Phone Input                  | Yes      | Destination phone number (E.164 format) |

### UI Components

#### Add Action Button

```
┌─────────────────────┐
│ [+] Add Action      │
└─────────────────────┘
```

- Located in the IVR Actions section header
- Adds a new blank action to the form
- Can add unlimited actions (recommended 3-5 for UX)

#### Delete Action Button

```
┌────┐
│ 🗑 │
└────┘
```

- Located in the top-right of each action card
- Removes the specific action from the form
- Confirms deletion before removing

#### Action Type Toggle

```
┌──────────────────┐
│ Information  ▼  │  or  │ Forward      ▼  │
└──────────────────┘      └──────────────────┘
```

- Dropdown selector
- Changes fields displayed below
- Two options: Information, Forward

## Form States

### Empty State (No Actions)

```
┌────────────────────────────────────────────────┐
│  IVR Actions              [+] Add Action       │
│                                                 │
│  ┌───────────────────────────────────────┐    │
│  │                                        │    │
│  │   No actions defined. Click "Add      │    │
│  │   Action" to create IVR menu options. │    │
│  │                                        │    │
│  └───────────────────────────────────────┘    │
└────────────────────────────────────────────────┘
```

### Loading State

```
┌──────────────────┐
│   Saving...      │  (button disabled)
└──────────────────┘
```

### Error State

```
┌────────────────────────────────────────────────┐
│  ⚠ Campaign name is required                  │
└────────────────────────────────────────────────┘
```

## Validation Rules

### Frontend Validation

- ✅ Campaign name cannot be empty
- ✅ Description cannot be empty
- ✅ Intro text cannot be empty
- ✅ Language must be selected
- ✅ Each action must have action type
- ✅ Each action must have key press (1-9)
- ✅ Information actions must have message
- ✅ Forward actions must have phone number

### Visual Indicators

- Required fields marked with red asterisk (\*)
- Error messages shown in red banner at top of form
- Helper text shown below relevant fields in gray
- Invalid inputs highlighted with red border

## Responsive Behavior

### Desktop (> 768px)

- Modal width: 768px (max-w-3xl)
- Two-column layout for Action Type and Key Press
- Full form visible without scrolling (if < 6 actions)

### Tablet (768px - 1024px)

- Modal width: 90% of screen
- Same layout as desktop
- Scrollable if many actions

### Mobile (< 768px)

- Modal width: 95% of screen with margins
- Single column layout
- Sticky header and footer
- Scrollable content area
- Touch-friendly buttons and inputs

## Keyboard Interactions

| Key           | Action                            |
| ------------- | --------------------------------- |
| `Esc`         | Close modal (same as Cancel)      |
| `Enter`       | Submit form (when in input field) |
| `Tab`         | Navigate between fields           |
| `Shift + Tab` | Navigate backwards                |

## Color Scheme

```css
Primary Blue: #2563eb (buttons, focus rings)
Gray Background: #f9fafb (action cards)
Border: #d1d5db
Text Primary: #111827
Text Secondary: #6b7280
Error Red: #dc2626
Success Green: #059669
```

## Icons Used

- `Plus` - Add action button
- `Trash2` - Delete action button
- `X` - Close modal button

## Action Card Styling

```
Background: Light gray (#f9fafb)
Border: 1px solid #e5e7eb
Border Radius: 8px
Padding: 16px
Gap between elements: 12px
```

## Best Practices for UX

1. **Progressive Disclosure**: Show fields based on action type
2. **Clear Labels**: Every field has descriptive label
3. **Helper Text**: Guidance text below complex fields
4. **Visual Hierarchy**: Sections clearly separated
5. **Inline Validation**: Errors shown immediately
6. **Loading States**: Disabled buttons during save
7. **Responsive Design**: Works on all screen sizes
8. **Accessibility**: Proper labels, focus management
9. **Error Recovery**: Clear error messages with solutions
10. **Empty States**: Helpful messages when no data

## Accessibility Features

- ✅ All inputs have associated labels
- ✅ Required fields indicated with asterisk and aria-required
- ✅ Error messages linked to inputs with aria-describedby
- ✅ Keyboard navigation supported throughout
- ✅ Focus visible on all interactive elements
- ✅ Color not sole indicator of state
- ✅ Sufficient color contrast ratios
- ✅ Screen reader friendly
